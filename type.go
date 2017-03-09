package equifax

import (
	"encoding/xml"
	"time"
)

const timeEquifaxFormat = "02.01.2006 15:04:05"
const dateEquifaxFormat = "02.01.2006"

type Time struct {
	time.Time
}

func (et *Time) getValue() string {
	return et.Format(timeEquifaxFormat)
}

func (et *Time) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: et.getValue()}, nil
}

func (et *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(et.getValue(), start)
	return nil
}

type Date struct {
	time.Time
}

func (et *Date) ToTime() time.Time {
	return time.Time(et)
}

func (et *Date) getValue() string {
	if et.IsZero() {
		return emptyDate.Format(dateEquifaxFormat)
	}
	return et.Format(dateEquifaxFormat)
}

func (et *Date) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: et.getValue()}, nil
}

func (et *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(et.getValue(), start)
	return nil
}
