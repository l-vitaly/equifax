package test

import (
	"testing"
	"time"

	"github.com/l-vitaly/cryptopro"
	"github.com/l-vitaly/equifax"
)

func TestCredit(t *testing.T) {
	store, err := cryptopro.SystemStore("MY")
	defer store.Close()
	if err != nil {
		t.Error(err)
	}
	crt, err := store.GetByThumb("5f08160e7dca8db7b8b3fd1b055a6c4300c37ba6")
	defer crt.Close()
	if err != nil {
		t.Error(err)
	}

	c := equifax.NewEquifaxCredit("http://10.130.1.2/xml.php", "90J", crt)

	resp, err := c.Get(&equifax.CreditRequest{
		Num:          1,
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
			Country: equifax.CitizenshipTypeRU,
			City:    "МОСКВА",
			Region:  equifax.RegionType00,
			Street:  "6 КВАРТАЛ",
			House:   "17",
			Flat:    "48",
		},
		AddressReg: &equifax.AddressReg{
			Index:   "000000",
			Country: equifax.CitizenshipTypeRU,
			City:    "МОСКВА",
			Region:  equifax.RegionType00,
			Street:  "6 КВАРТАЛ",
			House:   "17",
			Flat:    "48",
		},
	})

    if resp.Code != equifax.ResponseCodeType1 {
        t.Errorf("Expected code 1 actual %d", resp.Code)
    }
}
