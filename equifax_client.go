package equifax

import (
	"encoding/xml"
	"time"
)

const timeEquifaxFormat = "02.01.2006 15:04:05"
const dateEquifaxFormat = "02.01.2006"
const strEmpty = "EMPTY"
const Null = "NULL"

type EmptyString string

func (t EmptyString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == "" {
		e.EncodeElement(strEmpty, start)
		return nil
	}
	e.EncodeElement((string)(t), start)
	return nil
}

type Time struct {
	time.Time
}

func (et *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(et.Format(timeEquifaxFormat), start)
	return nil
}

type Date struct {
	time.Time
}

func (et *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if et.IsZero() {
		emptyDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
		e.EncodeElement(emptyDate.Format(dateEquifaxFormat), start)
		return nil
	}
	e.EncodeElement(et.Format(dateEquifaxFormat), start)
	return nil
}

type Credential struct {
	Login     string `xml:"login"`
	Password  string `xml:"password"`
	PartnerID string `xml:"partnerid"`
}

type NewApplication struct {
	XMLName xml.Name `xml:"fps:newApplication"`
	Credential
	ApplicationID              string                 `xml:"applicationid"`
	ApplicationDate            Time                   `xml:"applicationdate"`
	PhotoID                    string                 `xml:"photoid"`
	LastName                   string                 `xml:"lastname"`
	FirstName                  string                 `xml:"firstname"`
	MiddleName                 EmptyString            `xml:"middlename"`
	PastLastName               EmptyString             `xml:"pastlastname"`
	Birthday                   Date                   `xml:"birthday"`
	Birthplace                 EmptyString             `xml:"birthplace"`
	DocType                    DocType                `xml:"doctype"`
	DocNo                      string                 `xml:"docno"`
	DocPlace                   EmptyString            `xml:"docplace"`
	DocDate                    Date                   `xml:"docdate"`
	DocCode                    EmptyString             `xml:"doccode"`
	PastDocType                DocType                `xml:"pastdoctype"`
	PastDocNo                  EmptyString             `xml:"pastdocno"`
	PastDocPlace               EmptyString             `xml:"pastdocplace"`
	PastDocDate                Date                   `xml:"pastdocdate"`
	Sex                        Sex                    `xml:"sex"`
	Citizenship                Citizenship            `xml:"citizenship"`
	INN                        EmptyString             `xml:"inn"`
	PFR                        EmptyString             `xml:"pfr"`
	DriverNo                   EmptyString             `xml:"driverno"`
	Education                  Education              `xml:"education"`
	Marital                    Marital                `xml:"marital"`
	NumChildren                EmptyString             `xml:"numchildren"`
	Email                      EmptyString             `xml:"email"`
	HomePhone                  EmptyString             `xml:"homephone"`
	MobilePhone                EmptyString             `xml:"mobilephone"`
	LaCountry                  Citizenship            `xml:"la_country"`
	LaIndex                    EmptyString            `xml:"la_index"`
	LaRegion                   EmptyString            `xml:"la_region"`
	LaDistrict                 EmptyString            `xml:"la_district"`
	LaCity                     EmptyString            `xml:"la_city"`
	LaSettlement               EmptyString            `xml:"la_settlement"`
	LaStreet                   EmptyString            `xml:"la_street"`
	LaHouse                    EmptyString            `xml:"la_house"`
	LaBuilding                 EmptyString            `xml:"la_building"`
	LaStructure                EmptyString            `xml:"la_structure"`
	LaApartment                EmptyString            `xml:"la_apartment"`
	LaYears                    EmptyString             `xml:"la_years"`
	LaMonth                    EmptyString             `xml:"la_month"`
	LaDate                     Date                   `xml:"la_date"`
	RaPhone                    EmptyString             `xml:"ra_phone"`
	RaCountry                  Citizenship            `xml:"ra_country"`
	RaIndex                    EmptyString            `xml:"ra_index"`
	RaRegion                   EmptyString            `xml:"ra_region"`
	RaDistrict                 EmptyString            `xml:"ra_district"`
	RaCity                     EmptyString            `xml:"ra_city"`
	RaSettlement               EmptyString            `xml:"ra_settlement"`
	RaStreet                   EmptyString            `xml:"ra_street"`
	RaHouse                    EmptyString            `xml:"ra_house"`
	RaBuilding                 EmptyString            `xml:"ra_building"`
	RaStructure                EmptyString            `xml:"ra_structure"`
	RaApartment                EmptyString            `xml:"ra_apartment"`
	EmployerName               EmptyString             `xml:"employername"`
	EmployerSize               EmployerSize           `xml:"employersize"`
	BusinessIndustry           BusinessIndustry       `xml:"businessindustry"`
	Position                   EmptyString             `xml:"position"`
	EmploymentYear             EmptyString             `xml:"employment_year"`
	EmploymentMonth            EmptyString             `xml:"employment_month"`
	EmploymentDate             Date                   `xml:"employment_date"`
	EmploymentINN              EmptyString             `xml:"employment_inn"`
	IncomeProof                IncomeProof            `xml:"incomeproof"`
	MonthlyIncome              EmptyString             `xml:"monthlyincome"`
	BaPhone                    EmptyString             `xml:"ba_phone"`
	BaCountry                  Citizenship            `xml:"ba_country"`
	BaIndex                    EmptyString            `xml:"ba_index"`
	BaRegion                   EmptyString            `xml:"ba_region"`
	BaDistrict                 EmptyString            `xml:"ba_district"`
	BaCity                     EmptyString            `xml:"ba_city"`
	BaSettlement               EmptyString            `xml:"ba_settlement"`
	BaStreet                   EmptyString            `xml:"ba_street"`
	BaHouse                    EmptyString            `xml:"ba_house"`
	BaBuilding                 EmptyString            `xml:"ba_building"`
	BaStructure                EmptyString            `xml:"ba_structure"`
	BaApartment                EmptyString            `xml:"ba_apartment"`
	ProductType                ProductType            `xml:"producttype"`
	ProductName                EmptyString             `xml:"productname"`
	OriginalChannel            OriginalChannel        `xml:"originalchannel"`
	ProductSumLimit            EmptyString             `xml:"productsumlimit"`
	ProductSumCurrency         SumCurrency            `xml:"productsumcurrency"`
	DownPaymentAmount          EmptyString             `xml:"downpaymentamount"`
	CollateralExistence        CollateralExistence    `xml:"collateralexistence"`
	CollateralValue            EmptyString             `xml:"collateralvalue"`
	PurchaseExistence          PurchaseExistence      `xml:"purchaseexistence"`
	PurchaseValue              EmptyString             `xml:"purchasevalue"`
	PurchaseModel              EmptyString             `xml:"purchasemodel"`
	OperatorCode               EmptyString             `xml:"operator_code"`
	OperatorName               EmptyString             `xml:"operator_name"`
	PosCode                    EmptyString            `xml:"pos_code"`
	PosPhone                   EmptyString             `xml:"pos_phone"`
	PosCountry                 Citizenship            `xml:"pos_country"`
	PosIndex                   EmptyString            `xml:"pos_index"`
	PosRegion                  EmptyString            `xml:"pos_region"`
	PosDistrict                EmptyString            `xml:"pos_district"`
	PosCity                    EmptyString            `xml:"pos_city"`
	PosSettlement              EmptyString            `xml:"pos_settlement"`
	PosStreet                  EmptyString            `xml:"pos_street"`
	PosHouse                   EmptyString            `xml:"pos_house"`
	PosBuilding                EmptyString            `xml:"pos_building"`
	PosStructure               EmptyString            `xml:"pos_structure"`
	PosApartment               EmptyString            `xml:"pos_apartment"`
	NewApplicant               NewApplicant           `xml:"newapplicant"`
	ApplicantType              ApplicantType          `xml:"applicanttype"`
	ApplicantTypeNum           int64                  `xml:"applicanttypenum"`
	ResponseIsNeeded           ResponseIsNeeded       `xml:"responseisneeded"`
	ApplicationStatus          ApplicationStatus      `xml:"applicationstatus"`
	ApplicantID                EmptyString            `xml:"applicantid"`
	TradeDate                  Date                   `xml:"tradedate"`
	InitialSumLimit            EmptyString             `xml:"initialsumlimit"`
	InitialSumCurrency         SumCurrency            `xml:"initialsumcurrency"`
	ApplicationFraudStatus     ApplicationFraudStatus `xml:"applicationfraudstatus"`
	ApplicationFraudStatusDesc EmptyString             `xml:"applicationfraudstatusdescr"`
	DefaultStatus              DefaultStatus          `xml:"defaultstatus"`
}

type NewApplicationResponse struct {
	XMLName       xml.Name `xml:"newApplicationResponse"`
	ApplicationID string   `xml:"applicationid"`
	Status        Status   `xml:"status"`
}

type OutputVector struct {
	XMLName xml.Name `xml:"fps:outputVector"`
	Credential
	ApplicationID    string        `xml:"applicationid"`
	ApplicationDate  Time          `xml:"applicationdate"`
	ApplicantType    ApplicantType `xml:"applicanttype"`
	ApplicantTypeNum int64         `xml:"applicanttypenum"`
}

type OutputVectorResponse struct {
	XMLName           xml.Name `xml:"outputVectorResponse"`
	ApplicationID     string   `xml:"applicationid"`
	Status            Status   `xml:"status"`
	MainRules         string   `xml:"mainrules"`
	MainScoreValue    string   `xml:"mainscorevalue"`
	SpecificRules     string   `xml:"specificrules"`
	ApplicationsFound string   `xml:"applicationsfound"`
}

type UpdateCreditStatus struct {
	XMLName xml.Name `xml:"fps:updateCreditStatus"`
	Credential
	ApplicationID      string            `xml:"applicationid"`
	ApplicationDate    Time              `xml:"applicationdate"`
	ApplicationStatus  ApplicationStatus `xml:"applicationstatus"`
	ApplicantID        string            `xml:"applicantid"`
	TradeDate          Date              `xml:"tradedate"`
	InitialSumLimit    string            `xml:"initialsumlimit"`
	InitialSumCurrency SumCurrency       `xml:"initialsumcurrency"`
}

type UpdateCreditStatusResponse struct {
	XMLName       xml.Name `xml:"updateCreditStatusResponse"`
	ApplicationID string   `xml:"applicationid"`
	Status        Status   `xml:"status"`
}

type UpdateFraudStatus struct {
	XMLName xml.Name `xml:"fps:updateFraudStatus"`
	Credential
	ApplicationID              string                 `xml:"applicationid"`
	ApplicationDate            Time                   `xml:"applicationdate"`
	ApplicationFraudStatus     ApplicationFraudStatus `xml:"applicationfraudstatus"`
	ApplicationFraudStatusDesc string                 `xml:"applicationfraudstatusdescr"`
}

type UpdateFraudStatusResponse struct {
	XMLName       xml.Name `updateFraudStatusResponse"`
	ApplicationID string   `xml:"applicationid"`
	Status        Status   `xml:"status"`
}

type UpdateDefaultStatus struct {
	XMLName xml.Name `xml:"fps:updateDefaultStatus"`
	Credential
	ApplicationID      string        `xml:"applicationid"`
	ApplicationDate    Time          `xml:"applicationdate"`
	ApplicantID        string        `xml:"applicantid"`
	TradeDate          Date          `xml:"tradedate"`
	InitialSumLimit    float32       `xml:"initialsumlimit"`
	InitialSumCurrency SumCurrency   `xml:"initialsumcurrency"`
	DefaultStatus      DefaultStatus `xml:"defaultstatus"`
}

type UpdateDefaultStatusResponse struct {
	XMLName       xml.Name `xml:"updateDefaultStatusResponse"`
	ApplicationID string   `xml:"applicationid"`
	Status        Status   `xml:"status"`
}

type ProcessingApplication struct {
	XMLName xml.Name `xml:"fps:processingApplication"`
	Credential
	ApplicationID    string           `xml:"applicationid"`
	ApplicationDate  Time             `xml:"applicationdate"`
	ApplicantType    ApplicantType    `xml:"applicanttype"`
	ApplicantTypeNum string           `xml:"applicanttypenum"`
	ResponseIsNeeded ResponseIsNeeded `xml:"responseisneeded"`
}

type ProcessingApplicationResponse struct {
	XMLName       xml.Name `xml:"processingApplicationResponse"`
	ApplicationID string   `xml:"applicationid"`
	Status        Status   `xml:"status"`
}

type DeleteApplication struct {
	XMLName xml.Name `xml:"fps:deleteApplication"`
	Credential
	ApplicationID    string        `xml:"applicationid"`
	ApplicationDate  Time          `xml:"applicationdate"`
	ApplicantType    ApplicantType `xml:"applicanttype"`
	ApplicantTypeNum string        `xml:"applicanttypenum"`
}

type DeleteApplicationResponse struct {
	XMLName       xml.Name `xml:"deleteApplicationResponse"`
	ApplicationID string   `xml:"applicationid"`
	Status        Status   `xml:"status"`
}

type FPSPartnerClient struct {
	client    *SOAPClient
	login     string
	password  string
	partnerID string
}

func NewFPSPartnerClient(
	url string, login string, password string, partnerID string, enabledTLS bool,
	timeout time.Duration, auth *BasicAuth, logger Logger,
) *FPSPartnerClient {
	client := NewSOAPClient(url, enabledTLS, timeout, auth, logger)
	return &FPSPartnerClient{
		client:    client,
		login:     login,
		password:  password,
		partnerID: partnerID,
	}
}

func (s *FPSPartnerClient) SetHeader(header interface{}) {
	s.client.SetHeader(header)
}

func (s *FPSPartnerClient) NewApplication(req *NewApplication) (*NewApplicationResponse, error) {
	response := new(NewApplicationResponse)
	req.Credential = Credential{
		Login:     s.login,
		Password:  s.password,
		PartnerID: s.partnerID,
	}

	err := s.client.Call("#newApplication", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *FPSPartnerClient) OutputVector(req *OutputVector) (*OutputVectorResponse, error) {
	response := new(OutputVectorResponse)
	req.Credential = Credential{
		Login:     s.login,
		Password:  s.password,
		PartnerID: s.partnerID,
	}
	err := s.client.Call("#outputVector", req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *FPSPartnerClient) UpdateCreditStatus(req *UpdateCreditStatus) (*UpdateCreditStatusResponse, error) {
	response := new(UpdateCreditStatusResponse)
	req.Credential = Credential{
		Login:     s.login,
		Password:  s.password,
		PartnerID: s.partnerID,
	}
	err := s.client.Call("#updateCreditStatus", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *FPSPartnerClient) UpdateFraudStatus(req *UpdateFraudStatus) (*UpdateFraudStatusResponse, error) {
	response := new(UpdateFraudStatusResponse)
	req.Credential = Credential{
		Login:     s.login,
		Password:  s.password,
		PartnerID: s.partnerID,
	}
	err := s.client.Call("#updateFraudStatus", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *FPSPartnerClient) UpdateDefaultStatus(req *UpdateDefaultStatus) (*UpdateDefaultStatusResponse, error) {
	response := new(UpdateDefaultStatusResponse)
	req.Credential = Credential{
		Login:     s.login,
		Password:  s.password,
		PartnerID: s.partnerID,
	}
	err := s.client.Call("#updateDefaultStatus", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *FPSPartnerClient) ProcessingApplication(req *ProcessingApplication) (*ProcessingApplicationResponse, error) {
	response := new(ProcessingApplicationResponse)
	req.Credential = Credential{
		Login:     s.login,
		Password:  s.password,
		PartnerID: s.partnerID,
	}
	err := s.client.Call("#processingApplication", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *FPSPartnerClient) DeleteApplication(req *DeleteApplication) (*DeleteApplicationResponse, error) {
	response := new(DeleteApplicationResponse)
	req.Credential = Credential{
		Login:     s.login,
		Password:  s.password,
		PartnerID: s.partnerID,
	}
	err := s.client.Call("#deleteApplication", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
