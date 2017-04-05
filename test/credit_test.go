package test

import (
	"testing"
	"time"

	"github.com/l-vitaly/cryptopro"
	"github.com/l-vitaly/equifax"
	"github.com/l-vitaly/gounit"
)

func TestCredit(t *testing.T) {
	u := gounit.New(t)

	store, err := cryptopro.SystemStore("MY")
	defer store.Close()
	u.AssertNotError(err, "Open Store")

	crt, err := store.GetBySHA1("5f08160e7dca8db7b8b3fd1b055a6c4300c37ba6")
	defer crt.Close()
	u.AssertNotError(err, "Get Cert")

	c := equifax.NewEquifaxCredit("http://10.130.1.2/xml.php", "90J", crt, false)

	resp, err := c.Get(&equifax.CreditRequest{
		Num:          1,
		Type:         "30033",
		DateOfReport: equifax.Date{time.Now()},
		Reason:       1,
		Individual: &equifax.Individual{
			FirstName:  "СЕРГЕЙ",
			LastName:   "СЕРГЕЕВ",
			MiddleName: "СЕРГЕЕВИЧ",
			Gender:     equifax.GenderType1,
			Birthday:   equifax.Date{time.Date(1975, 1, 20, 0, 0, 0, 0, time.UTC)},
			Birthplace: "МОСКВА",
			IdentityDocument: &equifax.IdentityDocument{
				DocType:  equifax.DocType1,
				DocNO:    "2000000000",
				DocDate:  equifax.Date{time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)},
				DocPlace: "ОВД УЛЬЯНОВСКА",
			},
		},
		Application: &equifax.Application{
			Consent:        equifax.ConsentType1,
			ConsentDate:    equifax.Date{time.Now()},
			ConsentEndDate: equifax.Date{time.Now().Add(time.Hour * 24)},
			AdmCodeInForm:  equifax.AdmCodeInFormType1,
		},
		AddressFact: &equifax.AddressFact{
			Index:   "000000",
			Country: equifax.CountryTypeRU,
			City:    "МОСКВА",
			Region:  equifax.RegionType00,
			Street:  "6 КВАРТАЛ",
			House:   "17",
			Flat:    "48",
		},
		AddressReg: &equifax.AddressReg{
			Index:   "000000",
			Country: equifax.CountryTypeRU,
			City:    "МОСКВА",
			Region:  equifax.RegionType00,
			Street:  "6 КВАРТАЛ",
			House:   "17",
			Flat:    "48",
		},
	})

	u.AssertNotError(err, "Get Credit History")
	u.AssertContains(resp.Response.Code, []equifax.ResponseCode{
		equifax.ResponseCodeType1, equifax.ResponseCodeType0,
	}, "Response Code")
	u.AssertGreaterThan(0, len(resp.Response.TitlePart.Data), "TitlePart")
	u.AssertGreaterThan(0, len(resp.Response.BasePart.Data), "BasePart")
	u.AssertGreaterThan(0, len(resp.Response.AddPart.Data), "AddPart")
}
