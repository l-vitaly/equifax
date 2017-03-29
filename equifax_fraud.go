package equifax

import (
	"encoding/xml"
	"time"
)

type Credential struct {
	Login     string `xml:"login" csv:"-"`
	Password  string `xml:"password" csv:"-"`
	PartnerID string `xml:"partnerid" csv:"-"`
}

type NewApplication struct {
	XMLName xml.Name `xml:"fps:newApplication" csv:"-"`
	Credential
	ApplicationID              string                 `xml:"applicationid" csv:"applicationid"`
	ApplicationDate            Time                   `xml:"applicationdate" csv:"applicationdate"`
	PhotoID                    string                 `xml:"photoid" csv:"-"`
	LastName                   string                 `xml:"lastname" csv:"lastname"`
	FirstName                  string                 `xml:"firstname" csv:"firstname"`
	MiddleName                 EmptyString            `xml:"middlename" csv:"middlename"`
	PastLastName               EmptyString            `xml:"pastlastname" csv:"pastlastname"`
	Birthday                   Date                   `xml:"birthday" csv:"birthday"`
	Birthplace                 EmptyString            `xml:"birthplace" csv:"birthplace"`
	DocType                    DocType                `xml:"doctype" csv:"doctype"`
	DocNo                      string                 `xml:"docno" csv:"docno"`
	DocPlace                   EmptyString            `xml:"docplace" csv:"docplace"`
	DocDate                    Date                   `xml:"docdate" csv:"docdate"`
	DocCode                    EmptyString            `xml:"doccode" csv:"doccode"`
	PastDocType                DocType                `xml:"pastdoctype" csv:"pastdoctype"`
	PastDocNo                  EmptyString            `xml:"pastdocno" csv:"pastdocno"`
	PastDocPlace               EmptyString            `xml:"pastdocplace" csv:"pastdocplace"`
	PastDocDate                Date                   `xml:"pastdocdate" csv:"pastdocdate"`
	Sex                        Sex                    `xml:"sex" csv:"sex"`
	Citizenship                Country                `xml:"citizenship" csv:"citizenship"`
	INN                        EmptyString            `xml:"inn" csv:"inn"`
	PFR                        EmptyString            `xml:"inn" csv:"pfr"`
	DriverNo                   EmptyString            `xml:"driverno" csv:"driverno"`
	Education                  Education              `xml:"education" csv:"education"`
	Marital                    Marital                `xml:"marital" csv:"marital"`
	NumChildren                EmptyString            `xml:"numchildren" csv:"numchildren"`
	Email                      EmptyString            `xml:"email" csv:"email"`
	HomePhone                  EmptyString            `xml:"homephone" csv:"homephone"`
	MobilePhone                EmptyString            `xml:"mobilephone" csv:"mobilephone"`
	LaCountry                  Country                `xml:"la_country" csv:"la_country"`
	LaIndex                    EmptyString            `xml:"la_index" csv:"la_index"`
	LaRegion                   EmptyString            `xml:"la_region" csv:"la_region"`
	LaDistrict                 EmptyString            `xml:"la_district" csv:"la_district"`
	LaCity                     EmptyString            `xml:"la_city" csv:"la_city"`
	LaSettlement               EmptyString            `xml:"la_settlement" csv:"la_settlement"`
	LaStreet                   EmptyString            `xml:"la_street" csv:"la_street"`
	LaHouse                    EmptyString            `xml:"la_house" csv:"la_house"`
	LaBuilding                 EmptyString            `xml:"la_building" csv:"la_building"`
	LaStructure                EmptyString            `xml:"la_structure" csv:"la_structure"`
	LaApartment                EmptyString            `xml:"la_apartment" csv:"la_apartment"`
	LaYears                    EmptyString            `xml:"la_years" csv:"la_years"`
	LaMonth                    EmptyString            `xml:"la_month" csv:"la_month"`
	LaDate                     Date                   `xml:"la_date" csv:"la_date"`
	RaPhone                    EmptyString            `xml:"ra_phone" csv:"ra_phone"`
	RaCountry                  Country                `xml:"ra_country" csv:"ra_country"`
	RaIndex                    EmptyString            `xml:"ra_index" csv:"ra_index"`
	RaRegion                   EmptyString            `xml:"ra_region" csv:"ra_region"`
	RaDistrict                 EmptyString            `xml:"ra_district" csv:"ra_district"`
	RaCity                     EmptyString            `xml:"ra_city" csv:"ra_city"`
	RaSettlement               EmptyString            `xml:"ra_settlement" csv:"ra_settlement"`
	RaStreet                   EmptyString            `xml:"ra_street" csv:"ra_street"`
	RaHouse                    EmptyString            `xml:"ra_house" csv:"ra_house"`
	RaBuilding                 EmptyString            `xml:"ra_building" csv:"ra_building"`
	RaStructure                EmptyString            `xml:"ra_structure" csv:"ra_structure"`
	RaApartment                EmptyString            `xml:"ra_apartment" csv:"ra_apartment"`
	EmployerName               EmptyString            `xml:"employername" csv:"employername"`
	EmployerSize               EmployerSize           `xml:"employersize" csv:"employersize"`
	BusinessIndustry           BusinessIndustry       `xml:"businessindustry" csv:"businessindustry"`
	Position                   EmptyString            `xml:"position" csv:"position"`
	EmploymentYear             EmptyString            `xml:"employment_year" csv:"employment_year"`
	EmploymentMonth            EmptyString            `xml:"employment_month" csv:"employment_month"`
	EmploymentDate             Date                   `xml:"employment_date" csv:"employment_date"`
	EmploymentINN              EmptyString            `xml:"employment_inn" csv:"employment_inn"`
	IncomeProof                IncomeProof            `xml:"incomeproof" csv:"incomeproof"`
	MonthlyIncome              EmptyString            `xml:"monthlyincome" csv:"monthlyincome"`
	BaPhone                    EmptyString            `xml:"ba_phone" csv:"ba_phone"`
	BaCountry                  Country                `xml:"ba_country" csv:"ba_country"`
	BaIndex                    EmptyString            `xml:"ba_index" csv:"ba_index"`
	BaRegion                   EmptyString            `xml:"ba_region" csv:"ba_region"`
	BaDistrict                 EmptyString            `xml:"ba_district" csv:"ba_district"`
	BaCity                     EmptyString            `xml:"ba_city" csv:"ba_city"`
	BaSettlement               EmptyString            `xml:"ba_settlement" csv:"ba_settlement"`
	BaStreet                   EmptyString            `xml:"ba_street" csv:"ba_street"`
	BaHouse                    EmptyString            `xml:"ba_house" csv:"ba_house"`
	BaBuilding                 EmptyString            `xml:"ba_building" csv:"ba_building"`
	BaStructure                EmptyString            `xml:"ba_structure" csv:"ba_structure"`
	BaApartment                EmptyString            `xml:"ba_apartment" csv:"ba_apartment"`
	ProductType                ProductType            `xml:"producttype" csv:"producttype"`
	ProductName                EmptyString            `xml:"productname" csv:"productname"`
	OriginalChannel            OriginalChannel        `xml:"originalchannel" csv:"originalchannel"`
	ProductSumLimit            EmptyString            `xml:"productsumlimit" csv:"productsumlimit"`
	ProductSumCurrency         SumCurrency            `xml:"productsumcurrency" csv:"productsumcurrency"`
	DownPaymentAmount          EmptyString            `xml:"downpaymentamount" csv:"downpaymentamount"`
	CollateralExistence        CollateralExistence    `xml:"collateralexistence" csv:"collateralexistence"`
	CollateralValue            EmptyString            `xml:"collateralvalue" csv:"collateralvalue"`
	PurchaseExistence          PurchaseExistence      `xml:"purchaseexistence" csv:"purchaseexistence"`
	PurchaseValue              EmptyString            `xml:"purchasevalue" csv:"purchasevalue"`
	PurchaseModel              EmptyString            `xml:"purchasemodel" csv:"purchasemodel"`
	OperatorCode               EmptyString            `xml:"operator_code" csv:"operator_code"`
	OperatorName               EmptyString            `xml:"operator_name" csv:"operator_name"`
	PosCode                    EmptyString            `xml:"pos_code" csv:"pos_code"`
	PosPhone                   EmptyString            `xml:"pos_phone" csv:"pos_phone"`
	PosCountry                 Country                `xml:"pos_country" csv:"pos_country"`
	PosIndex                   EmptyString            `xml:"pos_index" csv:"pos_index"`
	PosRegion                  EmptyString            `xml:"pos_region" csv:"pos_region"`
	PosDistrict                EmptyString            `xml:"pos_district" csv:"pos_district"`
	PosCity                    EmptyString            `xml:"pos_city" csv:"pos_city"`
	PosSettlement              EmptyString            `xml:"pos_settlement" csv:"pos_settlement"`
	PosStreet                  EmptyString            `xml:"pos_street" csv:"pos_street"`
	PosHouse                   EmptyString            `xml:"pos_house" csv:"pos_house"`
	PosBuilding                EmptyString            `xml:"pos_building" csv:"pos_building"`
	PosStructure               EmptyString            `xml:"pos_structure" csv:"pos_structure"`
	PosApartment               EmptyString            `xml:"pos_apartment" csv:"pos_apartment"`
	NewApplicant               NewApplicant           `xml:"newapplicant" csv:"newapplicant"`
	ApplicantType              ApplicantType          `xml:"applicanttype" csv:"applicanttype"`
	ApplicantTypeNum           int64                  `xml:"applicanttypenum" csv:"applicanttypenum"`
	ResponseIsNeeded           ResponseIsNeeded       `xml:"responseisneeded" csv:"responseisneeded"`
	ApplicationStatus          ApplicationStatus      `xml:"applicationstatus" csv:"applicationstatus"`
	ApplicantID                EmptyString            `xml:"applicantid" csv:"applicantid"`
	TradeDate                  Date                   `xml:"tradedate" csv:"tradedate"`
	InitialSumLimit            EmptyString            `xml:"initialsumlimit" csv:"initialsumlimit"`
	InitialSumCurrency         SumCurrency            `xml:"initialsumcurrency" csv:"initialsumcurrency"`
	ApplicationFraudStatus     ApplicationFraudStatus `xml:"applicationfraudstatus" csv:"applicationfraudstatus"`
	ApplicationFraudStatusDesc EmptyString            `xml:"applicationfraudstatusdescr" csv:"applicationfraudstatusdescr"`
	DefaultStatus              DefaultStatus          `xml:"defaultstatus" csv:"defaultstatus"`
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
	MainScoreValue    int32    `xml:"mainscorevalue,-"`
	SpecificRules     string   `xml:"specificrules"`
	ApplicationsFound int32    `xml:"applicationsfound,-"`
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
	InitialSumLimit    string        `xml:"initialsumlimit"`
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

type EquifaxFraud interface {
	SetHeader(header interface{})
	NewApplication(req *NewApplication) (*NewApplicationResponse, error)
	OutputVector(req *OutputVector) (*OutputVectorResponse, error)
	UpdateCreditStatus(req *UpdateCreditStatus) (*UpdateCreditStatusResponse, error)
	UpdateFraudStatus(req *UpdateFraudStatus) (*UpdateFraudStatusResponse, error)
	ProcessingApplication(req *ProcessingApplication) (*ProcessingApplicationResponse, error)
	DeleteApplication(req *DeleteApplication) (*DeleteApplicationResponse, error)
}

type equifaxFraud struct {
	client    *SOAPClient
	login     string
	password  string
	partnerID string
}

func NewEquifaxFraud(
	url string, login string, password string, partnerID string, enabledTLS bool,
	timeout time.Duration, auth *BasicAuth, logger Logger,
) EquifaxFraud {
	client := NewSOAPClient(url, enabledTLS, timeout, auth, logger)
	return &equifaxFraud{
		client:    client,
		login:     login,
		password:  password,
		partnerID: partnerID,
	}
}

func (s *equifaxFraud) SetHeader(header interface{}) {
	s.client.SetHeader(header)
}

func (s *equifaxFraud) NewApplication(req *NewApplication) (*NewApplicationResponse, error) {
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

func (s *equifaxFraud) OutputVector(req *OutputVector) (*OutputVectorResponse, error) {
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

func (s *equifaxFraud) UpdateCreditStatus(req *UpdateCreditStatus) (*UpdateCreditStatusResponse, error) {
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

func (s *equifaxFraud) UpdateFraudStatus(req *UpdateFraudStatus) (*UpdateFraudStatusResponse, error) {
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

func (s *equifaxFraud) UpdateDefaultStatus(req *UpdateDefaultStatus) (*UpdateDefaultStatusResponse, error) {
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

func (s *equifaxFraud) ProcessingApplication(req *ProcessingApplication) (*ProcessingApplicationResponse, error) {
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

func (s *equifaxFraud) DeleteApplication(req *DeleteApplication) (*DeleteApplicationResponse, error) {
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
