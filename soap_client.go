package equifax

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type Logger interface {
	Log(keyvals ...interface{}) error
}

type NullLogger struct {
}

func (*NullLogger) Log(keyvals ...interface{}) error {
    return nil
}

type SOAPEnvelopeResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SOAPBodyResponse
}

type SOAPBodyResponse struct {
	XMLName xml.Name    `xml:"Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

func (b *SOAPBodyResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}
	return nil
}

type SOAPFault struct {
	XMLName xml.Name `xml:"Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}

func (f *SOAPFault) Error() string {
	return f.String
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Soapenv string   `xml:"xmlns:soapenv,attr"`
	Fps     string   `xml:"xmlns:fps,attr"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"soapenv:Header"`
	Header  interface{}
}

type SOAPBody struct {
	XMLName xml.Name    `xml:"soapenv:Body"`
	Content interface{} `xml:",omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url       string
	enableTLS bool
	auth      *BasicAuth
	header    interface{}
	client    *http.Client
	logger    Logger
}

func NewSOAPClient(url string, enableTLS bool, timeout time.Duration, auth *BasicAuth, logger Logger) *SOAPClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: enableTLS,
		},
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
	}

	if logger == nil {
		logger = new(NullLogger)
	}

	client := &http.Client{Transport: tr}
	return &SOAPClient{
		url:       url,
		enableTLS: enableTLS,
		auth:      auth,
		client:    client,
		logger:    logger,
	}
}

func (s *SOAPClient) SetHeader(header interface{}) {
	s.header = header
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	buffer, err := s.buildRequest(request)
	if err != nil {
		return err
	}

    s.logger.Log("equfax_request", buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "equifaxFraud-client/0.1")
	req.Close = true

	res, err := s.client.Do(req)
	if err != nil {
		return err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	rawBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

    s.logger.Log("equfax_response", string(rawBody))

	respEnvelope, err := s.makeResponse(rawBody, response)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}

func (s *SOAPClient) buildRequest(request interface{}) (*bytes.Buffer, error) {
	envelope := SOAPEnvelope{
		Xsi:     "http://schemas.xmlsoap.org/soap/envelope/",
		Xsd:     "http://www.w3.org/2001/XMLSchema",
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		Fps:     "http://example.org/FPSPartner",
	}

	if s.header != nil {
		envelope.Header = &SOAPHeader{Header: s.header}
	}

	envelope.Body.Content = request

	buffer := new(bytes.Buffer)
	buffer.Write([]byte(xml.Header))

	encoder := xml.NewEncoder(buffer)

	if err := encoder.Encode(envelope); err != nil {
		return nil, err
	}

	if err := encoder.Flush(); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (s *SOAPClient) makeResponse(rawBody []byte, response interface{}) (*SOAPEnvelopeResponse, error) {
	if len(rawBody) == 0 {
		log.Println("empty response")
		return nil, nil
	}
	respEnvelope := new(SOAPEnvelopeResponse)
	respEnvelope.Body = SOAPBodyResponse{Content: response}
	err := xml.Unmarshal(rawBody, respEnvelope)
	if err != nil {
		return nil, err
	}
	return respEnvelope, nil
}
