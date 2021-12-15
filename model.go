package main

import (
	"encoding/xml"
	"github.com/j03hanafi/bankiso/iso20022/head"
)

type BusMsg struct {
	AppHdr   *head.BusinessApplicationHeaderV01 `xml:"AppHdr" json:"AppHdr"`
	Document interface{}                        `xml:"Document" json:"Document"`
}

type ChannelInput struct {
	BusMsg BusMsg `xml:"BusMsg" json:"BusMsg"`
}

type Channel struct {
	XMLName        xml.Name                           `xml:"BusMsg"`
	Ns             string                             `xml:"ns,attr"`
	Ns1            string                             `xml:"ns1,attr"`
	Ns2            string                             `xml:"ns2,attr"`
	Xsi            string                             `xml:"xsi,attr"`
	SchemaLocation string                             `xml:"schemaLocation,attr"`
	AppHdr         *head.BusinessApplicationHeaderV01 `xml:"AppHdr" json:"AppHdr"`
}
