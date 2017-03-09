package equifax

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/l-vitaly/acharset"
	"github.com/l-vitaly/cryptopro"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/charmap"
)

const EquifaxCreditVersion = "3.4"

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrInvalidCertificate = errors.New("invalid certificate")
)

type bkiRequest struct {
	XMLName   xml.Name `xml:"bki_request"`
	Version   string   `xml:"version,attr"`
	PartnerID string   `xml:"partnerid,attr"` // код партнера
	Request   *CreditRequest
}

type AddressReg struct {
	XMLName   xml.Name     `xml:"addr_reg"`
	Owner     AddressOwner `xml:"owner"`                    // статус регистрации по данному адресу
	Index     string       `xml:"index"`                    // индекс
	AddrTotal string       `xml:"addr_reg_total,omitempty"` // адрес регистрации одной строкой
	Country   Country      `xml:"country"`                  // страна
	Region    Region       `xml:"region"`                   // код региона
	City      string       `xml:"city"`                     // населенный пункт
	District  string       `xml:"district,omitempty"`       // район
	Street    string       `xml:"street"`                   // улица
	House     string       `xml:"house"`                    // дом/блок/строение
	Flat      string       `xml:"flat"`                     // квартира/офис/комната
}

type AddressFact struct {
	XMLName   xml.Name     `xml:"addr_fact"`
	Owner     AddressOwner `xml:"owner"`                     // статус регистрации по данному адресу
	Index     string       `xml:"index"`                     // индекс
	AddrTotal string       `xml:"addr_fact_total,omitempty"` // адрес фактического местонахождения одной строкой
	Country   Country      `xml:"country"`                   // страна
	Region    Region       `xml:"region"`                    // код региона
	City      string       `xml:"city"`                      // населенный пункт
	District  string       `xml:"district,omitempty"`        // район
	Street    string       `xml:"street"`                    // улица
	House     string       `xml:"house"`                     // дом/блок/строение
	Flat      string       `xml:"flat"`                      // квартира/офис/комната
}

type EmploymentCompany struct {
	XMLName  xml.Name     `xml:"company"`
	Name     string       `xml:"name"`                // Наименование компании
	State    CompanyState `xml:"state"`               // Признак государственности компании
	Size     CompanySize  `xml:"size"`                // Размер компании
	Area     CompanyArea  `xml:"area"`                // Вид деятельности компании
	AreaText string       `xml:"area_text,omitempty"` // Вид деятельности компании (другое)
}

type Employment struct {
	XMLName        xml.Name          `xml:"employment"`
	Current        EmploymentCurrent `xml:"current"`                   // текущее / предыдущее место работы
	Duration       int               `xml:"duration"`                  // время работы на данном месте, месяцев
	Type           EmploymentType    `xml:"type"`                      // тип занятости
	Profession     Profession        `xml:"profession"`                // профессия
	ProfessionText string            `xml:"profession_text,omitempty"` // профессия (другое)
	Company        *EmploymentCompany
}

type ApplicationIndividual struct {
	XMLName         xml.Name    `xml:"private"`
	Citizenship     Country     `xml:"citizenship,omitempty"`      // гражданство
	Marriage        Marital     `xml:"marriage,omitempty"`         // семейное положение
	DependantsBel18 int         `xml:"dependants_bel18,omitempty"` // количество иждивенцев до 18 лет включительно
	DependantsUnd18 int         `xml:"dependants_und18,omitempty"` // количество иждивенцев старше 18 лет
	Education       Education   `xml:"education,omitempty"`        // образование
	PhoneMobile     string      `xml:"phone_mobile,omitempty"`     // мобильный телефон
	PhoneHome       string      `xml:"phone_home,omitempty"`       // домашний телефон
	PhoneWork       string      `xml:"phone_work,omitempty"`       // рабочий телефон
	Email           string      `xml:"email,omitempty"`            // адрес эл. почты
	Employment      *Employment // инф. о работе
}

type ApplicationLegalEntity struct {
	XMLName       xml.Name     `xml:"commercial"`
	State         CompanyState `xml:"company_state"`               // признак государственности компании
	Size          CompanySize  `xml:"company_size"`                // размер компании
	Area          CompanyArea  `xml:"company_area"`                // вид деятельности компании
	AreaText      string       `xml:"company_area_text,omitempty"` // вид деятельности компании (другое)
	BeginningDate Date         `xml:"company_beginning_date"`      // дата регистрации юридического лица
}

type Application struct {
	XMLName              xml.Name                `xml:"application"`
	Consent              Consent                 `xml:"consent"`                          // флаг согласия субъекта КИ на получение его кредитного отчета
	ConsentDate          Date                    `xml:"consentdate"`                      // дата выдачи согласия субъекта КИ
	ConsentEndDate       Date                    `xml:"consentenddate"`                   // дата окончания действия согласия субъекта КИ
	AdmCodeInForm        AdmCodeInForm           `xml:"admcode_inform"`                   // флаг информирования пользователя КИ об административной ответственности
	ConsentOwner         string                  `xml:"consent_owner,omitempty"`          // наименование пользователя КИ, который получил согласие субъекта КИ
	Income               float64                 `xml:"income,omitempty"`                 // доход
	IncomeFrequency      IncomeFrequency         `xml:"income_frequency,omitempty"`       // периодичность получения дохода
	Purpose              Purpose                 `xml:"purpose,omitempty"`                // цель финансирования
	PurposeText          string                  `xml:"purpose_text,omitempty"`           // цель финансирования (другая)
	Num                  string                  `xml:"application_num,omitempty"`        // номер заявки
	Date                 Date                    `xml:"application_date"`                 // дата подачи заявки
	CredType             Cred                    `xml:"cred_type,omitempty"`              // тип займа (кредита)
	CredCurrency         SumCurrency             `xml:"cred_currency,omitempty"`          // валюта кредита
	CredSum              float64                 `xml:"cred_sum,omitempty"`               // сумма (лимит) займа (кредита)
	CredDeposit          float64                 `xml:"cred_deposit,omitempty"`           // изначально внесенная сумма
	CredLastPayment      float64                 `xml:"cred_last_payment,omitempty"`      // сумма финального платежа (для лизинга)
	CredSumPayment       float64                 `xml:"cred_sum_payment,omitempty"`       // сумма платежа в период
	CredFrequencyPayment float64                 `xml:"cred_frequency_payment,omitempty"` // периодичность платежей
	CredDuration         float64                 `xml:"cred_duration,omitempty"`          // срок действия займа (кредита)
	CredSecurity         CredSecurity            `xml:"cred_security,omitempty"`          // тип обеспечения займа (кредита)
	Comment              string                  `xml:"comment,omitempty,omitempty"`      // комментарии
	Individual           *ApplicationIndividual  `xml:",omitempty"`                       // физ. лицо
	LegalEntity          *ApplicationLegalEntity `xml:",omitempty"`                       // юр. лицо
}

type LegalEntity struct {
	XMLName     xml.Name `xml:"commercial"`
	FullName    string   `xml:"fullname,omitempty"`    // полное наименование юридического лица
	ShortName   string   `xml:"shortname,omitempty"`   // сокращенное наименование юридического лица
	FirmName    string   `xml:"firmname,omitempty"`    // фирменное наименование юридического лица
	ForeignName string   `xml:"foreignname,omitempty"` // наименование юридического лица на языке народов РФ и (или) иностранном языке
	Resident    Resident `xml:"resident"`              // признак резидентства
	RegCountry  Country  `xml:"regcountry"`            // наименование государства регистрации
	Phone       string   `xml:"phone,omitempty"`       // контактные телефоны
	INN         string   `xml:"inn,omitempty"`         // ИНН
	EGRN        string   `xml:"egrn,omitempty"`        // ОГРН
}

type IdentityDocument struct {
	XMLName    xml.Name `xml:"doc"`
	DocType    DocType  `xml:"doctype"`    // тип документа
	DocNO      string   `xml:"docno"`      // серия и номер документа
	DocDate    Date     `xml:"docdate"`    // дата выдачи документа
	DocEndDate Date     `xml:"docenddate"` // дата окончания действия документа
	DocPlace   string   `xml:"docplace"`   // наименование органа, выдавшего документ, место выдачи документа, код органа, выдавшего документ, код органа вводиться через точку с зяпятой
}

type Individual struct {
	XMLName          xml.Name          `xml:"private"`
	LastName         string            `xml:"lastname"`         // имя
	FirstName        string            `xml:"firstname"`        // фамилия
	MiddleName       string            `xml:"middlename"`       // отчество
	Gender           Gender            `xml:"gender,omitempty"` // пол
	Birthday         Date              `xml:"birthday"`         // дата рождения
	Birthplace       string            `xml:"birthplace"`       // место рождения
	IdentityDocument *IdentityDocument // документ удостоверяющий личность
	INN              string            `xml:"inn,omitempty"`  // ИНН
	PfrNO            string            `xml:"pfno,omitempty"` // СНИЛС
}

type CreditRequest struct {
	XMLName      xml.Name     `xml:"request"`
	Num          int          `xml:"num,attr"` // ID запроса
	DateOfReport Date         `xml:"dateofreport,attr"`
	Individual   *Individual  `xml:",omitempty"`            // информация о физ. лице
	LegalEntity  *LegalEntity `xml:",omitempty"`            // информация о юр. лице
	Reason       Reason       `xml:"reason"`                // цель согласия
	ReasonText   string       `xml:"reason_text,omitempty"` // иная цель согласия
	Application  *Application `xml:""`                      // заявка
	AddressReg   *AddressReg  `xml:""`                      // адрес регистрации
	AddressFact  *AddressFact `xml:""`                      // адрес фактического места проживания
	Type         string       `xml:"type"`                  // идентификатор отчета
}

type CreditResponse struct {
	XMLName   xml.Name     `xml:"bki_response"`
	Version   string       `xml:"version,attr"`
	PartnerID string       `xml:"partnerid,attr"` // код партнера
	DateTime  string       `xml:"datetime,attr"`
	Code      ResponseCode `xml:"response>responsecode"`
	Text      string       `xml:"response>responsestring"`
	Data      []byte       `xml:",innerxml"`
}

type EquifaxCredit interface {
	Get(r *CreditRequest) (*CreditResponse, error)
}

type equifaxCredit struct {
	url       string
	partnerID string
	crt       cryptopro.Cert
}

func NewEquifaxCredit(url string, partnerID string, crt cryptopro.Cert) EquifaxCredit {
	return &equifaxCredit{
		url:       url,
		partnerID: partnerID,
		crt:       crt,
	}
}

func (e *equifaxCredit) Get(r *CreditRequest) (*CreditResponse, error) {
	r.Type = "30033"

	req := bkiRequest{
		Version:   EquifaxCreditVersion,
		PartnerID: e.partnerID,
		Request:   r,
	}

	reqBuf := bytes.NewBuffer([]byte{})
	reqBuf.WriteString(`<?xml version="1.0" encoding="windows-1251"?>` + "\n")

	err := xml.NewEncoder(reqBuf).Encode(req)
	if err != nil {
		return nil, err
	}

	reqEncBytes, err := charmap.Windows1251.NewEncoder().Bytes(reqBuf.Bytes())
	if err != nil {
		return nil, err
	}
	reqBuf = bytes.NewBuffer(reqEncBytes)

	dest := new(bytes.Buffer)

	msg, err := cryptopro.OpenToEncode(dest, cryptopro.EncodeOptions{
		Signers: []cryptopro.Cert{e.crt},
	})
	if err != nil {
		return nil, err
	}

	_, err = reqBuf.WriteTo(msg)
	if err != nil {
		return nil, err
	}

	msg.Close()

	resp, err := http.Post(e.url, "application/octet-stream", bytes.NewReader(dest.Bytes()))
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrInvalidRequest
	}

	respMsg, err := cryptopro.OpenToDecode(resp.Body)

	cBuf := bytes.NewBuffer([]byte{})
	_, err = io.Copy(cBuf, respMsg)
	if err != nil {
		return nil, err
	}

	err = respMsg.Verify(e.crt)
	if err != nil && err != cryptopro.ErrVerifyingSignature {
		return nil, ErrInvalidCertificate
	}

	var result *CreditResponse
	dec := xml.NewDecoder(cBuf)
	dec.CharsetReader = acharset.CharsetReader
	err = dec.Decode(&result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
