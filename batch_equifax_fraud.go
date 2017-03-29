package equifax

import (
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
)

type batchFraud struct {
}

func NewBatchFraud() *batchFraud {
	return &batchFraud{}
}

func csvWriter(out io.Writer) *csv.Writer {
	writer := csv.NewWriter(out)
	writer.Comma = '\t'
	return writer
}

func (*batchFraud) ToCSV(req []*NewApplication) (string, error) {
	gocsv.SetCSVWriter(csvWriter)

	res, err := gocsv.MarshalString(&req)
	if err != nil {
		return "", err
	}
	return res, nil
}
