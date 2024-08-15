package domain

import "encoding/xml"

type EnvelopeResponse struct {
	XMLName xml.Name    `xml:"Envelope"`
	Body    EnvelopBody `xml:"Body"`
}

type EnvelopBody struct {
	XMLName          xml.Name                `xml:"Body"`
	MultiplyResponse EnvelopMultiplyResponse `xml:"MultiplyResponse"`
}

type EnvelopMultiplyResponse struct {
	XMLName        xml.Name `xml:"MultiplyResponse"`
	MultiplyResult string   `xml:"MultiplyResult"`
}
