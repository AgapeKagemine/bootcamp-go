package domain

import "encoding/xml"

// Envelope was generated 2024-08-06 16:16:48 by https://xml-to-go.github.io/ in Ukraine.
type EnvelopeRequest struct {
	XMLName xml.Name           `xml:"soap:Envelope"`
	Xsi     string             `xml:"xmlns:xsi,attr"`
	Xsd     string             `xml:"xmlns:xsd,attr"`
	Soap    string             `xml:"xmlns:soap,attr"`
	Body    EnvelopRequestBody `xml:"soap:Body"`
}

type EnvelopRequestBody struct {
	XMLName  xml.Name               `xml:"soap:Body"`
	Multiply EnvelopRequestMultiply `xml:"Multiply"`
}

type EnvelopRequestMultiply struct {
	XMLName xml.Name `xml:"Multiply"`
	Xmlns   string   `xml:"xmlns,attr"`
	IntA    string   `xml:"intA"`
	IntB    string   `xml:"intB"`
}
