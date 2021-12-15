package main

import (
	"encoding/xml"
	"github.com/j03hanafi/bankiso/iso20022/acmt"
	"github.com/j03hanafi/bankiso/iso20022/admi"
	"github.com/j03hanafi/bankiso/iso20022/admn"
	"github.com/j03hanafi/bankiso/iso20022/auth"
	"github.com/j03hanafi/bankiso/iso20022/caaa"
	"github.com/j03hanafi/bankiso/iso20022/caam"
	"github.com/j03hanafi/bankiso/iso20022/cain"
	"github.com/j03hanafi/bankiso/iso20022/camt"
	"github.com/j03hanafi/bankiso/iso20022/catm"
	"github.com/j03hanafi/bankiso/iso20022/catp"
	"github.com/j03hanafi/bankiso/iso20022/head"
	"github.com/j03hanafi/bankiso/iso20022/pacs"
	"github.com/j03hanafi/bankiso/iso20022/pain"
	"github.com/j03hanafi/bankiso/iso20022/prxy"
	"github.com/j03hanafi/bankiso/iso20022/remt"
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
	Document       interface{}                        `xml:"Document"`
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

var ISO20022Registry = map[string]interface{}{
	"acmt.007.001.02": &acmt.Document00700102{},
	"acmt.008.001.02": &acmt.Document00800102{},
	"acmt.009.001.02": &acmt.Document00900102{},
	"acmt.010.001.02": &acmt.Document01000102{},
	"acmt.011.001.02": &acmt.Document01100102{},
	"acmt.012.001.02": &acmt.Document01200102{},
	"acmt.013.001.02": &acmt.Document01300102{},
	"acmt.014.001.02": &acmt.Document01400102{},
	"acmt.015.001.02": &acmt.Document01500102{},
	"acmt.016.001.02": &acmt.Document01600102{},
	"acmt.017.001.02": &acmt.Document01700102{},
	"acmt.018.001.02": &acmt.Document01800102{},
	"acmt.019.001.02": &acmt.Document01900102{},
	"acmt.020.001.02": &acmt.Document02000102{},
	"acmt.021.001.02": &acmt.Document02100102{},
	"acmt.022.001.02": &acmt.Document02200102{},
	"acmt.023.001.02": &acmt.Document02300102{},
	"acmt.024.001.02": &acmt.Document02400102{},
	"admi.002.001.01": &admi.Document00200101{},
	"admi.004.001.02": &admi.Document00400102{},
	"admi.011.001.01": &admi.Document01100101{},
	"admn.001.001.01": &admn.Document00100101{},
	"admn.002.001.01": &admn.Document00200101{},
	"auth.001.001.01": &auth.Document00100101{},
	"auth.002.001.01": &auth.Document00200101{},
	"auth.003.001.01": &auth.Document00300101{},
	"auth.018.001.01": &auth.Document01800101{},
	"auth.019.001.01": &auth.Document01900101{},
	"auth.020.001.01": &auth.Document02000101{},
	"auth.021.001.01": &auth.Document02100101{},
	"auth.022.001.01": &auth.Document02200101{},
	"auth.023.001.01": &auth.Document02300101{},
	"auth.024.001.01": &auth.Document02400101{},
	"auth.025.001.01": &auth.Document02500101{},
	"auth.026.001.01": &auth.Document02600101{},
	"auth.027.001.01": &auth.Document02700101{},
	"caaa.001.001.05": &caaa.Document00100105{},
	"caaa.002.001.05": &caaa.Document00200105{},
	"caaa.003.001.05": &caaa.Document00300105{},
	"caaa.004.001.05": &caaa.Document00400105{},
	"caaa.005.001.05": &caaa.Document00500105{},
	"caaa.006.001.05": &caaa.Document00600105{},
	"caaa.007.001.05": &caaa.Document00700105{},
	"caaa.008.001.05": &caaa.Document00800105{},
	"caaa.009.001.05": &caaa.Document00900105{},
	"caaa.010.001.05": &caaa.Document01000105{},
	"caaa.011.001.05": &caaa.Document01100105{},
	"caaa.012.001.05": &caaa.Document01200105{},
	"caaa.013.001.05": &caaa.Document01300105{},
	"caaa.014.001.05": &caaa.Document01400105{},
	"caaa.015.001.05": &caaa.Document01500105{},
	"caaa.016.001.03": &caaa.Document01600103{},
	"caaa.017.001.03": &caaa.Document01700103{},
	"caam.001.001.02": &caam.Document00100102{},
	"caam.002.001.02": &caam.Document00200102{},
	"caam.003.001.02": &caam.Document00300102{},
	"caam.004.001.02": &caam.Document00400102{},
	"caam.005.001.02": &caam.Document00500102{},
	"caam.006.001.02": &caam.Document00600102{},
	"caam.007.001.01": &caam.Document00700101{},
	"caam.008.001.01": &caam.Document00800101{},
	"caam.009.001.02": &caam.Document00900102{},
	"caam.010.001.02": &caam.Document01000102{},
	"caam.011.001.01": &caam.Document01100101{},
	"caam.012.001.01": &caam.Document01200101{},
	"cain.001.001.01": &cain.Document00100101{},
	"cain.002.001.01": &cain.Document00200101{},
	"cain.003.001.01": &cain.Document00300101{},
	"cain.004.001.01": &cain.Document00400101{},
	"cain.005.001.01": &cain.Document00500101{},
	"cain.006.001.01": &cain.Document00600101{},
	"cain.007.001.01": &cain.Document00700101{},
	"cain.008.001.01": &cain.Document00800101{},
	"cain.009.001.01": &cain.Document00900101{},
	"cain.010.001.01": &cain.Document01000101{},
	"cain.011.001.01": &cain.Document01100101{},
	"cain.012.001.01": &cain.Document01200101{},
	"cain.013.001.01": &cain.Document01300101{},
	"camt.026.001.04": &camt.Document02600104{},
	"camt.027.001.04": &camt.Document02700104{},
	"camt.028.001.06": &camt.Document02800106{},
	"camt.029.001.06": &camt.Document02900106{},
	"camt.030.001.04": &camt.Document03000104{},
	"camt.031.001.04": &camt.Document03100104{},
	"camt.032.001.03": &camt.Document03200103{},
	"camt.033.001.04": &camt.Document03300104{},
	"camt.034.001.04": &camt.Document03400104{},
	"camt.035.001.03": &camt.Document03500103{},
	"camt.036.001.03": &camt.Document03600103{},
	"camt.037.001.04": &camt.Document03700104{},
	"camt.038.001.03": &camt.Document03800103{},
	"camt.039.001.04": &camt.Document03900104{},
	"camt.052.001.06": &camt.Document05200106{},
	"camt.053.001.06": &camt.Document05300106{},
	"camt.054.001.06": &camt.Document05400106{},
	"camt.055.001.05": &camt.Document05500105{},
	"camt.056.001.05": &camt.Document05600105{},
	"camt.057.001.05": &camt.Document05700105{},
	"camt.058.001.05": &camt.Document05800105{},
	"camt.059.001.05": &camt.Document05900105{},
	"camt.060.001.03": &camt.Document06000103{},
	"camt.086.001.02": &camt.Document08600102{},
	//"camt.087.001.03": &camt.Document08700103{}, //had a problem generating message definition. Got amigous message definition RequestToModifyPaymentV03
	"catm.001.001.05": &catm.Document00100105{},
	"catm.002.001.05": &catm.Document00200105{},
	"catm.003.001.05": &catm.Document00300105{},
	"catm.004.001.04": &catm.Document00400104{},
	"catm.005.001.02": &catm.Document00500102{},
	"catm.006.001.02": &catm.Document00600102{},
	"catm.007.001.01": &catm.Document00700101{},
	"catm.008.001.01": &catm.Document00800101{},
	"catp.001.001.02": &catp.Document00100102{},
	"catp.002.001.02": &catp.Document00200102{},
	"catp.003.001.02": &catp.Document00300102{},
	"catp.004.001.02": &catp.Document00400102{},
	"catp.005.001.02": &catp.Document00500102{},
	"catp.006.001.02": &catp.Document00600102{},
	"catp.007.001.02": &catp.Document00700102{},
	"catp.008.001.02": &catp.Document00800102{},
	"catp.009.001.02": &catp.Document00900102{},
	"catp.010.001.02": &catp.Document01000102{},
	"catp.011.001.02": &catp.Document01100102{},
	"catp.012.001.01": &catp.Document01200101{},
	"catp.013.001.01": &catp.Document01300101{},
	"catp.014.001.01": &catp.Document01400101{},
	"catp.015.001.01": &catp.Document01500101{},
	"catp.016.001.01": &catp.Document01600101{},
	"catp.017.001.01": &catp.Document01700101{},
	"head.001.001.01": &head.Document00100101{},
	"pacs.002.001.07": &pacs.Document00200107{},
	"pacs.002.001.10": &pacs.Document00200110{},
	"pacs.003.001.06": &pacs.Document00300106{},
	"pacs.004.001.06": &pacs.Document00400106{},
	"pacs.007.001.06": &pacs.Document00700106{},
	"pacs.008.001.06": &pacs.Document00800106{},
	"pacs.008.001.08": &pacs.Document00800108{},
	"pacs.009.001.06": &pacs.Document00900106{},
	"pacs.009.001.09": &pacs.Document00900109{},
	"pacs.010.001.02": &pacs.Document01000102{},
	"pacs.028.001.04": &pacs.Document02800104{},
	"pain.001.001.07": &pain.Document00100107{},
	"pain.002.001.07": &pain.Document00200107{},
	"pain.007.001.06": &pain.Document00700106{},
	"pain.008.001.06": &pain.Document00800106{},
	"pain.009.001.04": &pain.Document00900104{},
	"pain.010.001.04": &pain.Document01000104{},
	"pain.011.001.04": &pain.Document01100104{},
	"pain.012.001.04": &pain.Document01200104{},
	"pain.013.001.05": &pain.Document01300105{},
	"pain.014.001.05": &pain.Document01400105{},
	"prxy.001.001.01": &prxy.Document00100101{},
	"prxy.002.001.01": &prxy.Document00200101{},
	"prxy.003.001.01": &prxy.Document00300101{},
	"prxy.004.001.01": &prxy.Document00400101{},
	"prxy.005.001.01": &prxy.Document00500101{},
	"prxy.006.001.01": &prxy.Document00600101{},
	"prxy.901.001.01": &prxy.Document90100101{},
	"remt.001.001.02": &remt.Document00100102{},
	"remt.002.001.01": &remt.Document00200101{},
}
