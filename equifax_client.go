package equifax

import (
	"encoding/xml"
	"time"
)

type NewApplication struct {
	XMLName                    xml.Name               `xml:"http://equifax.ru/FPSPartner newApplication"`
	Login                      string                 `xml:"login,omitempty"`
	Password                   string                 `xml:"password,omitempty"`
	PartnerID                  string                 `xml:"partnerid,omitempty"`
	ApplicationID              string                 `xml:"applicationid,omitempty"`
	ApplicationDate            time.Time              `xml:"applicationdate,omitempty"`
	PhotoID                    string                 `xml:"photoid,omitempty"`
	LastName                   string                 `xml:"lastname,omitempty"`
	FirstName                  string                 `xml:"firstname,omitempty"`
	MiddleName                 string                 `xml:"middlename,omitempty"`
	PastLastName               string                 `xml:"pastlastname,omitempty"`
	Birthday                   time.Time              `xml:"birthday,omitempty"`
	Birthplace                 string                 `xml:"birthplace,omitempty"`
	DocType                    DocType                `xml:"doctype,omitempty"`
	DocNo                      string                 `xml:"docno,omitempty"`
	DocPlace                   string                 `xml:"docplace,omitempty"`
	DocDate                    time.Time              `xml:"docdate,omitempty"`
	DocCode                    string                 `xml:"doccode,omitempty"`
	PastDocType                DocType                `xml:"pastdoctype,omitempty"`
	PastDocNo                  string                 `xml:"pastdocno,omitempty"`
	PastDocPlace               string                 `xml:"pastdocplace,omitempty"`
	PastDocDate                time.Time              `xml:"pastdocdate,omitempty"`
	Sex                        Sex                    `xml:"sex,omitempty"`
	Citizenship                Citizenship            `xml:"citizenship,omitempty"`
	INN                        string                 `xml:"inn,omitempty"`
	PFR                        string                 `xml:"pfr,omitempty"`
	DriverNo                   string                 `xml:"driverno,omitempty"`
	Education                  Education              `xml:"education,omitempty"`
	Marital                    Marital                `xml:"marital,omitempty"`
	NumChildren                string                 `xml:"numchildren,omitempty"`
	Email                      string                 `xml:"email,omitempty"`
	HomePhone                  string                 `xml:"homephone,omitempty"`
	MobilePhone                string                 `xml:"mobilephone,omitempty"`
	LaCountry                  Citizenship            `xml:"la_country,omitempty"`
	LaIndex                    string                 `xml:"la_index,omitempty"`
	LaRegion                   string                 `xml:"la_region,omitempty"`
	LaDistrict                 string                 `xml:"la_district,omitempty"`
	LaCity                     string                 `xml:"la_city,omitempty"`
	LaSettlement               string                 `xml:"la_settlement,omitempty"`
	LaStreet                   string                 `xml:"la_street,omitempty"`
	LaHouse                    string                 `xml:"la_house,omitempty"`
	LaBuilding                 string                 `xml:"la_building,omitempty"`
	LaStructure                string                 `xml:"la_structure,omitempty"`
	LaApartment                string                 `xml:"la_apartment,omitempty"`
	LaYears                    string                 `xml:"la_years,omitempty"`
	LaMonth                    string                 `xml:"la_month,omitempty"`
	LaDate                     time.Time              `xml:"la_date,omitempty"`
	RaPhone                    string                 `xml:"ra_phone,omitempty"`
	RaCountry                  Citizenship            `xml:"ra_country,omitempty"`
	RaIndex                    string                 `xml:"ra_index,omitempty"`
	RaRegion                   string                 `xml:"ra_region,omitempty"`
	RaDistrict                 string                 `xml:"ra_district,omitempty"`
	RaCity                     string                 `xml:"ra_city,omitempty"`
	RaSettlement               string                 `xml:"ra_settlement,omitempty"`
	RaStreet                   string                 `xml:"ra_street,omitempty"`
	RaHouse                    string                 `xml:"ra_house,omitempty"`
	RaBuilding                 string                 `xml:"ra_building,omitempty"`
	RaStructure                string                 `xml:"ra_structure,omitempty"`
	RaApartment                string                 `xml:"ra_apartment,omitempty"`
	EmployerName               string                 `xml:"employername,omitempty"`
	EmployerSize               EmployerSize           `xml:"employersize,omitempty"`
	BusinessIndustry           BusinessIndustry       `xml:"businessindustry,omitempty"`
	Position                   string                 `xml:"position,omitempty"`
	EmploymentYear             string                 `xml:"employment_year,omitempty"`
	EmploymentMonth            string                 `xml:"employment_month,omitempty"`
	EmploymentDate             time.Time              `xml:"employment_date,omitempty"`
	EmploymentINN              string                 `xml:"employment_inn,omitempty"`
	IncomeProof                IncomeProof            `xml:"incomeproof,omitempty"`
	MonthlyIncome              string                 `xml:"monthlyincome,omitempty"`
	BaPhone                    string                 `xml:"ba_phone,omitempty"`
	BaCountry                  Citizenship            `xml:"ba_country,omitempty"`
	BaIndex                    string                 `xml:"ba_index,omitempty"`
	BaRegion                   string                 `xml:"ba_region,omitempty"`
	BaDistrict                 string                 `xml:"ba_district,omitempty"`
	BaCity                     string                 `xml:"ba_city,omitempty"`
	BaSettlement               string                 `xml:"ba_settlement,omitempty"`
	BaStreet                   string                 `xml:"ba_street,omitempty"`
	BaHouse                    string                 `xml:"ba_house,omitempty"`
	BaBuilding                 string                 `xml:"ba_building,omitempty"`
	BaStructure                string                 `xml:"ba_structure,omitempty"`
	BaApartment                string                 `xml:"ba_apartment,omitempty"`
	ProductType                ProductType            `xml:"producttype,omitempty"`
	ProductName                string                 `xml:"productname,omitempty"`
	OriginalChannel            OriginalChannel        `xml:"originalchannel,omitempty"`
	ProductSumLimit            string                 `xml:"productsumlimit,omitempty"`
	ProductSumCurrency         ProductSumCurrency     `xml:"productsumcurrency,omitempty"`
	DownPaymentAmount          string                 `xml:"downpaymentamount,omitempty"`
	CollateralExistence        CollateralExistence    `xml:"collateralexistence,omitempty"`
	CollateralValue            string                 `xml:"collateralvalue,omitempty"`
	PurchaseExistence          PurchaseExistence      `xml:"purchaseexistence,omitempty"`
	PurchaseValue              string                 `xml:"purchasevalue,omitempty"`
	PurchaseModel              string                 `xml:"purchasemodel,omitempty"`
	OperatorCode               string                 `xml:"operator_code,omitempty"`
	OperatorName               string                 `xml:"operator_name,omitempty"`
	PosCode                    string                 `xml:"pos_code,omitempty"`
	PosPhone                   string                 `xml:"pos_phone,omitempty"`
	PosCountry                 Citizenship            `xml:"pos_country,omitempty"`
	PosIndex                   string                 `xml:"pos_index,omitempty"`
	PosRegion                  string                 `xml:"pos_region,omitempty"`
	PosDistrict                string                 `xml:"pos_district,omitempty"`
	PosCity                    string                 `xml:"pos_city,omitempty"`
	PosSettlement              string                 `xml:"pos_settlement,omitempty"`
	PosStreet                  string                 `xml:"pos_street,omitempty"`
	PosHouse                   string                 `xml:"pos_house,omitempty"`
	PosBuilding                string                 `xml:"pos_building,omitempty"`
	PosStructure               string                 `xml:"pos_structure,omitempty"`
	PosApartment               string                 `xml:"pos_apartment,omitempty"`
	NewApplicant               NewApplicant           `xml:"newapplicant,omitempty"`
	ApplicantType              ApplicantType          `xml:"applicanttype,omitempty"`
	ApplicantTypeNum           string                 `xml:"applicanttypenum,omitempty"`
	ResponseIsNeeded           ResponseIsNeeded       `xml:"responseisneeded,omitempty"`
	ApplicationStatus          ApplicationStatus      `xml:"applicationstatus,omitempty"`
	ApplicantID                string                 `xml:"applicantid,omitempty"`
	TradeDate                  time.Time              `xml:"tradedate,omitempty"`
	InitialSumLimit            string                 `xml:"initialsumlimit,omitempty"`
	InitialSumCurrency         ProductSumCurrency     `xml:"initialsumcurrency,omitempty"`
	ApplicationFraudStatus     ApplicationFraudStatus `xml:"applicationfraudstatus,omitempty"`
	ApplicationFraudStatusDesc string                 `xml:"applicationfraudstatusdescr,omitempty"`
	DefaultStatus              DefaultStatus          `xml:"defaultstatus,omitempty"`
}

type NewApplicationResponse struct {
	XMLName       xml.Name `xml:"http://equifax.ru/FPSPartner newApplicationResponse"`
	ApplicationID string   `xml:"applicationid,omitempty"`
	Status        Status   `xml:"status,omitempty"`
}

type OutputVector struct {
	XMLName          xml.Name      `xml:"http://equifax.ru/FPSPartner outputVector"`
	Login            string        `xml:"login,omitempty"`
	Password         string        `xml:"password,omitempty"`
	PartnerID        string        `xml:"partnerid,omitempty"`
	ApplicationID    string        `xml:"applicationid,omitempty"`
	ApplicationDate  time.Time     `xml:"applicationdate,omitempty"`
	ApplicantType    ApplicantType `xml:"applicanttype,omitempty"`
	ApplicantTypeNum string        `xml:"applicanttypenum,omitempty"`
}

type OutputVectorResponse struct {
	XMLName           xml.Name `xml:"http://equifax.ru/FPSPartner outputVectorResponse"`
	ApplicationID     string   `xml:"applicationid,omitempty"`
	Status            Status   `xml:"status,omitempty"`
	MainRules         string   `xml:"mainrules,omitempty"`
	MainScoreValue    string   `xml:"mainscorevalue,omitempty"`
	SpecificRules     string   `xml:"specificrules,omitempty"`
	ApplicationsFound string   `xml:"applicationsfound,omitempty"`
}

type UpdateCreditStatus struct {
	XMLName            xml.Name           `xml:"http://equifax.ru/FPSPartner updateCreditStatus"`
	Login              string             `xml:"login,omitempty"`
	Password           string             `xml:"password,omitempty"`
	PartnerID          string             `xml:"partnerid,omitempty"`
	ApplicationID      string             `xml:"applicationid,omitempty"`
	ApplicationDate    time.Time          `xml:"applicationdate,omitempty"`
	ApplicationStatus  ApplicationStatus  `xml:"applicationstatus,omitempty"`
	ApplicantID        string             `xml:"applicantid,omitempty"`
	TradeDate          time.Time          `xml:"tradedate,omitempty"`
	InitialSumLimit    string             `xml:"initialsumlimit,omitempty"`
	InitialSumCurrency ProductSumCurrency `xml:"initialsumcurrency,omitempty"`
}

type UpdateCreditStatusResponse struct {
	XMLName       xml.Name `xml:"http://equifax.ru/FPSPartner updateCreditStatusResponse"`
	ApplicationID string   `xml:"applicationid,omitempty"`
	Status        Status   `xml:"status,omitempty"`
}

type UpdateFraudStatus struct {
	XMLName                    xml.Name               `xml:"http://equifax.ru/FPSPartner updateFraudStatus"`
	Login                      string                 `xml:"login,omitempty"`
	Password                   string                 `xml:"password,omitempty"`
	PartnerID                  string                 `xml:"partnerid,omitempty"`
	ApplicationID              string                 `xml:"applicationid,omitempty"`
	ApplicationDate            time.Time              `xml:"applicationdate,omitempty"`
	ApplicationFraudStatus     ApplicationFraudStatus `xml:"applicationfraudstatus,omitempty"`
	ApplicationFraudStatusDesc string                 `xml:"applicationfraudstatusdescr,omitempty"`
}

type UpdateFraudStatusResponse struct {
	XMLName       xml.Name `xml:"http://equifax.ru/FPSPartner updateFraudStatusResponse"`
	ApplicationID string   `xml:"applicationid,omitempty"`
	Status        Status   `xml:"status,omitempty"`
}

type UpdateDefaultStatus struct {
	XMLName            xml.Name           `xml:"http://equifax.ru/FPSPartner updateDefaultStatus"`
	Login              string             `xml:"login,omitempty"`
	Password           string             `xml:"password,omitempty"`
	PartnerID          string             `xml:"partnerid,omitempty"`
	ApplicationID      string             `xml:"applicationid,omitempty"`
	ApplicationDate    time.Time          `xml:"applicationdate,omitempty"`
	ApplicantID        string             `xml:"applicantid,omitempty"`
	TradeDate          time.Time          `xml:"tradedate,omitempty"`
	InitialSumLimit    float32            `xml:"initialsumlimit,omitempty"`
	InitialSumCurrency ProductSumCurrency `xml:"initialsumcurrency,omitempty"`
	DefaultStatus      DefaultStatus      `xml:"defaultstatus,omitempty"`
}

type UpdateDefaultStatusResponse struct {
	XMLName       xml.Name `xml:"http://equifax.ru/FPSPartner updateDefaultStatusResponse"`
	ApplicationID string   `xml:"applicationid,omitempty"`
	Status        Status   `xml:"status,omitempty"`
}

type ProcessingApplication struct {
	XMLName          xml.Name         `xml:"http://equifax.ru/FPSPartner processingApplication"`
	Login            string           `xml:"login,omitempty"`
	Password         string           `xml:"password,omitempty"`
	PartnerID        string           `xml:"partnerid,omitempty"`
	ApplicationID    string           `xml:"applicationid,omitempty"`
	ApplicationDate  time.Time        `xml:"applicationdate,omitempty"`
	ApplicantType    ApplicantType    `xml:"applicanttype,omitempty"`
	ApplicantTypeNum string           `xml:"applicanttypenum,omitempty"`
	ResponseIsNeeded ResponseIsNeeded `xml:"responseisneeded,omitempty"`
}

type ProcessingApplicationResponse struct {
	XMLName       xml.Name `xml:"http://equifax.ru/FPSPartner processingApplicationResponse"`
	ApplicationID string   `xml:"applicationid,omitempty"`
	Status        Status   `xml:"status,omitempty"`
}

type DeleteApplication struct {
	XMLName          xml.Name      `xml:"http://equifax.ru/FPSPartner deleteApplication"`
	Login            string        `xml:"login,omitempty"`
	Password         string        `xml:"password,omitempty"`
	PartnerID        string        `xml:"partnerid,omitempty"`
	ApplicationID    string        `xml:"applicationid,omitempty"`
	ApplicationDate  time.Time     `xml:"applicationdate,omitempty"`
	ApplicantType    ApplicantType `xml:"applicanttype,omitempty"`
	ApplicantTypeNum string        `xml:"applicanttypenum,omitempty"`
}

type DeleteApplicationResponse struct {
	XMLName       xml.Name `xml:"http://equifax.ru/FPSPartner deleteApplicationResponse"`
	ApplicationID string   `xml:"applicationid,omitempty"`
	Status        Status   `xml:"status,omitempty"`
}

type FPSPartnerClient struct {
	client *SOAPClient
}

func NewFPSPartnerClient(url string, enabledTLS bool, timeout time.Duration, auth *BasicAuth) *FPSPartnerClient {
	client := NewSOAPClient(url, enabledTLS, timeout, auth)
	return &FPSPartnerClient{
		client: client,
	}
}

func (service *FPSPartnerClient) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

func (service *FPSPartnerClient) NewApplication(req *NewApplication) (*NewApplicationResponse, error) {
	response := new(NewApplicationResponse)
	err := service.client.Call("#newApplication", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *FPSPartnerClient) OutputVector(req *OutputVector) (*OutputVectorResponse, error) {
	response := new(OutputVectorResponse)
	err := service.client.Call("#outputVector", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *FPSPartnerClient) UpdateCreditStatus(req *UpdateCreditStatus) (*UpdateCreditStatusResponse, error) {
	response := new(UpdateCreditStatusResponse)
	err := service.client.Call("#updateCreditStatus", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *FPSPartnerClient) UpdateFraudStatus(req *UpdateFraudStatus) (*UpdateFraudStatusResponse, error) {
	response := new(UpdateFraudStatusResponse)
	err := service.client.Call("#updateFraudStatus", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *FPSPartnerClient) UpdateDefaultStatus(req *UpdateDefaultStatus) (*UpdateDefaultStatusResponse, error) {
	response := new(UpdateDefaultStatusResponse)
	err := service.client.Call("#updateDefaultStatus", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *FPSPartnerClient) ProcessingApplication(req *ProcessingApplication) (*ProcessingApplicationResponse, error) {
	response := new(ProcessingApplicationResponse)
	err := service.client.Call("#processingApplication", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *FPSPartnerClient) DeleteApplication(req *DeleteApplication) (*DeleteApplicationResponse, error) {
	response := new(DeleteApplicationResponse)
	err := service.client.Call("#deleteApplication", req, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
