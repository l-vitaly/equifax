package test

import (
	"testing"
	"time"

	"github.com/l-vitaly/equifax"
)

func TestClient(t *testing.T) {
	c := equifax.NewFPSPartnerClient(
		"http://10.130.11.151/soap/bank",
		"",
		"",
		"",
		false,
		15*time.Second,
		nil,
		true,
	)
	resp, err := c.NewApplication(&equifax.NewApplication{})
	t.Log(resp, err.(*equifax.SOAPFault).Code)
}
