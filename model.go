package main

import (
	"github.com/j03hanafi/bankiso/iso20022/head"
)

type BusMsg struct {
	AppHdr *head.BusinessApplicationHeaderV01 `xml:"AppHdr"`
}
