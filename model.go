package main

import (
	"encoding/xml"

	"github.com/j03hanafi/bankiso/iso20022/head"
)

type BusMsg struct {
	AppHdr *head.BusinessApplicationHeaderV01 `xml:"AppHdr"`
}

type Message interface {
	String() (result string, ok bool)
}

type Document struct {
	Namespaces map[string]string
	BusMsg     BusMsg `xml:"BusMsg"`
}

func (a *Document) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a.Namespaces = map[string]string{}
	for _, attr := range start.Attr {
		if attr.Name.Space == "xmlns" {
			a.Namespaces[attr.Name.Local] = attr.Value
		}
	}

	// Go on with unmarshalling.
	type app Document
	aa := (*app)(a)
	return d.DecodeElement(aa, &start)
}
