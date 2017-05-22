package equifax

import (
    "encoding/xml"
    "time"
)

const timeEquifaxFormat = "02.01.2006 15:04:05"
const dateEquifaxFormat = "02.01.2006"

const strEmpty = "EMPTY"
const Null = "NULL"

var emptyDateDefault = time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)

type EmptyString string

func (t EmptyString) getValue() string {
    if t == "" {
        return strEmpty
    }
    return (string)(t)
}

func (t EmptyString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    e.EncodeElement(t.getValue(), start)
    return nil
}

func (t EmptyString) MarshalCSV() (string, error) {
    return t.getValue(), nil
}

type Time struct {
    time.Time
}

func (et *Time) getValue() string {
    return et.Format(timeEquifaxFormat)
}

func (et *Time) MarshalCSV() (string, error) {
    return et.getValue(), nil
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

func (et *Date) getValue() string {
    if et.IsZero() {
        return ""
    }
    return et.Format(dateEquifaxFormat)
}

func (et *Date) MarshalCSV() (string, error) {
    return et.getValue(), nil
}

func (et *Date) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
    return xml.Attr{Name: name, Value: et.getValue()}, nil
}

func (et *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    e.EncodeElement(et.getValue(), start)
    return nil
}
