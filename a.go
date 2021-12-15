package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

func tes() {
	s := `<?xml version="1.0" encoding="UTF-8"?>
	<!-- pacs.008 - example request for account enquiry from INDOIDJA into CI Hub (FASTIDJA) -->
	<ns:BusMsg xmlns:ns="urn:iso" xmlns:ns1="urn:iso:std:iso:20022:tech:xsd:head.001.001.01" xmlns:ns2="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:iso ../../../xsd/phase1/MainCIHub.xsd ">
	   <ns:AppHdr>
		  <ns1:Fr>
			 <ns1:FIId>
				<ns1:FinInstnId>
				   <ns1:Othr>
					  <ns1:Id>INDOIDJA</ns1:Id>
					  <!-- Sending system -->
				   </ns1:Othr>
				</ns1:FinInstnId>
			 </ns1:FIId>
		  </ns1:Fr>
		  <ns1:To>
			 <ns1:FIId>
				<ns1:FinInstnId>
				   <ns1:Othr>
					  <ns1:Id>FASTIDJA</ns1:Id>
					  <!-- Receiving system -->
				   </ns1:Othr>
				</ns1:FinInstnId>
			 </ns1:FIId>
		  </ns1:To>
		  <ns1:BizMsgIdr>20210301INDOIDJA510ORB12345678</ns1:BizMsgIdr>
		  <ns1:MsgDefIdr>pacs.008.001.08</ns1:MsgDefIdr>
		  <ns1:CreDt>2021-03-01T12:00:00Z</ns1:CreDt>
	   </ns:AppHdr>
	   <ns:Document>
		  <ns:FIToFICstmrCdtTrf>
			 <ns2:GrpHdr>
				<ns2:MsgId>20210301INDOIDJA51012345678</ns2:MsgId>
				<!-- YYYYMMDDBBBBBBBBBTTTSSSSSSSS -->
				<ns2:CreDtTm>2021-03-01T19:00:00.000</ns2:CreDtTm>
				<ns2:NbOfTxs>1</ns2:NbOfTxs>
				<ns2:SttlmInf>
				   <ns2:SttlmMtd>CLRG</ns2:SttlmMtd>
				</ns2:SttlmInf>
			 </ns2:GrpHdr>
			 <ns2:CdtTrfTxInf>
				<ns2:PmtId>
				   <ns2:EndToEndId>20210301INDOIDJA510ORB12345678</ns2:EndToEndId>
				   <ns2:TxId>20210301INDOIDJA11012345678</ns2:TxId>
				</ns2:PmtId>
				<ns2:PmtTpInf>
				   <ns2:CtgyPurp>
					  <ns2:Prtry>51099</ns2:Prtry>
				   </ns2:CtgyPurp>
				</ns2:PmtTpInf>
				<ns2:IntrBkSttlmAmt Ccy="IDR">1234.56</ns2:IntrBkSttlmAmt>
				<ns2:ChrgBr>DEBT</ns2:ChrgBr>
				<!-- Dbtr info not needed for AE, but Dbtr tag must be present for XSD validation -->
				<ns2:Dbtr />
				<ns2:DbtrAgt>
				   <ns2:FinInstnId>
					  <ns2:Othr>
						 <ns2:Id>INDOIDJA</ns2:Id>
						 <!-- Debiting Agent Bank ID -->
					  </ns2:Othr>
				   </ns2:FinInstnId>
				</ns2:DbtrAgt>
				<ns2:CdtrAgt>
				   <ns2:FinInstnId>
					  <ns2:Othr>
						 <ns2:Id>CENAIDJA</ns2:Id>
						 <!-- Crediting Agent Bank ID -->
					  </ns2:Othr>
				   </ns2:FinInstnId>
				</ns2:CdtrAgt>
				<!-- Cdtr info not needed for AE, but Cdtr tag must be present for XSD validation -->
				<ns2:Cdtr />
				<ns2:CdtrAcct>
				   <ns2:Id>
					  <ns2:Othr>
						 <ns2:Id>987654321</ns2:Id>
						 <!-- Crediting Account Number - needed for AE -->
					  </ns2:Othr>
				   </ns2:Id>
				</ns2:CdtrAcct>
			 </ns2:CdtTrfTxInf>
		  </ns:FIToFICstmrCdtTrf>
	   </ns:Document>
	</ns:BusMsg>`
	fmt.Println("here")

	// xmlFile, err := os.Open("xml/x.xml")

	doc, err := xmlquery.Parse(strings.NewReader(s))
	fmt.Println("doc")
	if err != nil {
		panic(err)
	}

	root := xmlquery.FindOne(doc, "//ns:BusMsg")
	if n := root.SelectElement("//ns:Document/ns:FIToFICstmrCdtTrf/ns2:CdtTrfTxInf/ns2:CdtrAcct/ns2:Id/ns2:Othr/ns2:Id"); n != nil {
		fmt.Printf("Name #%s\n", n.InnerText())
	}

}
