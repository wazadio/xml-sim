package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/j03hanafi/bankiso/iso20022/pacs"

	"github.com/antchfx/xmlquery"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Found error in log ", err)
	}
	log.SetOutput(file)
	log.Println("Log setup")

	path := pathHandler()

	address := ":6067"
	log.Printf("Biller started at %v", address)
	err = http.ListenAndServe(address, path)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func pathHandler() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/biller", biller).Methods("POST")
	router.HandleFunc("/biller3", biller3).Methods("POST")

	return router
}

func biller3(w http.ResponseWriter, r *http.Request) {

	log.Println("New Request from BIFast Connector XML")
	fmt.Println("New Request from BIFast Connector XML")

	body, _ := ioutil.ReadAll(r.Body)
	rawRequest := string(body)
	log.Println(rawRequest)

	request := BusMsg{}
	err := xml.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
	}
	fmt.Println(request)
	appHdr := request.AppHdr
	bizMsgIdr := fmt.Sprintf("%v", *appHdr.BusinessMessageIdentifier)
	trxType := bizMsgIdr[16:19]
	fmt.Println(trxType)
	var fileName string

	switch trxType {
	case "510":
		type Doc struct {
			XMLName  xml.Name              `xml:"BusMsg"`
			Document pacs.Document00800108 `xml:"Document"`
		}
		document := Doc{}
		err := xml.Unmarshal(body, &document)
		if err != nil {
			log.Printf("Error unmarshal JSON: %s", err.Error())
		}

		CrAccId := *document.Document.Message.CreditTransferTransactionInformation[0].CdtrAcct.Id.Other.Identification
		fmt.Println(CrAccId)
		fileName = "accEnqRes.xml"
	}
	fileName = "xml/" + fileName
	fmt.Println("filename:", fileName)
	file, _ := os.Open(fileName)
	defer file.Close()

	response, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	responseFormatter2(w, response, 200)

}

func biller(w http.ResponseWriter, r *http.Request) {
	log.Println("New Request from BIFast Connector")
	fmt.Println("New Request from BIFast Connector")

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

	requestRaw := BusMsg{}
	err := xml.Unmarshal(body, &requestRaw)
	if err != nil {
		log.Printf("Error unmarshal JSON: %s", err.Error())
	}
	fmt.Println("request: ", requestRaw)

	// var msgID string
	var fileName string
	bzMsgID := fmt.Sprintf("%v", *requestRaw.AppHdr.BusinessMessageIdentifier)
	trxType := bzMsgID[16:19]
	fmt.Println("trxType:", trxType)

	switch trxType {
	// ##################### Account Enquiry ##################################
	case "510":
		var CrAccId string
		doc, err := xmlquery.Parse(strings.NewReader(string(body)))
		if err != nil {
			panic(err)
		}

		root := xmlquery.FindOne(doc, "//ns:BusMsg")
		if n := root.SelectElement("//ns:Document/ns:FIToFICstmrCdtTrf/ns2:CdtTrfTxInf/ns2:CdtrAcct/ns2:Id/ns2:Othr/ns2:Id"); n != nil {
			fmt.Printf("Name #%s\n", n.InnerText())
			CrAccId = n.InnerText()
		}

		switch CrAccId {
		case "510654300":
			fileName = "accountEnquiryResponse.xml"
		case "511654182":
			fileName = "sampleAccountEnquiry2.json"
		}

		//##################### Credit Transfer ###################################
		// case "010": // Credit Transfer

		// 	CrAccId := *document.Message.CreditTransferTransactionInformation[0].CdtrAcct.Id.Other.Identification
		// 	switch CrAccId {
		// 	case "0102345600":
		// 		fileName = "sampleCreditTransferResponse.json"
		// 	case "0102345184":
		// 		fileName = "sampleCreditTransferResponse2.json"
		// 	}
		// case "012":
		// 	fileName = "sampleCreditTransferResponse012.json"
		// 	fmt.Println("012")
		// case "110":
		// 	fileName = "sampleCreditTransferResponsewithProxy.json"
		// 	fmt.Println("110")
		// //==========================================================================

		// case "019":
		// 	fileName = "sampleFItoFICreditTransfer.json"
		// 	fmt.Println("019")
		// case "011":
		// 	fileName = "sampleReverseCreditTransfer.json"
		// 	fmt.Println("011")

		// // ################# Proxy Resolution #####################################
		// case "610":

		// 	PxValue := *document.Message.LookUp.PrxyOnly.PrxyRtrvl.Val
		// 	switch PxValue {
		// 	case "086102345000":
		// 		fileName = "sampleProxyResolution.json"
		// 	case "086112345101":
		// 		fileName = "sampleProxyResolution2.json"
		// 	case "086112345804":
		// 		fileName = "sampleProxyResolution3.json"
		// 	case "086132345600":
		// 		fileName = "sampleProxyResolution4.json"
		// 	case "086142345804":
		// 		fileName = "sampleProxyResolution5.json"
		// 	case "08615234804":
		// 		fileName = "sampleProxyResolution6.json"
		// 	case "08616234811":
		// 		fileName = "sampleProxyResolution7.json"
		// 	case "08617234805":
		// 		fileName = "sampleProxyResolution8.json"
		// 	}
		// case "611":
		// 	fileName = "sampleProxyResolution611.json"
		// 	fmt.Println("611")
		// case "612":
		// 	fileName = "sampleProxyResolution612.json"
		// 	fmt.Println(("612"))
		// // =========================================================================

		// // ################# Proxy Registration Inquiry ############################
		// case "620":

		// 	CsAccId := *document.Message.GroupHeader.MessageSender.Account.Identification.Other.Identification
		// 	fmt.Println(CsAccId)
		// 	switch CsAccId {
		// 	case "6202345600":
		// 		fileName = "sampleProxyRegistrationInquiry.json"
		// 	case "6212345101":
		// 		fileName = "sampleProxyRegistrationInquiry2.json"
		// 	case "6222345808":
		// 		fileName = "sampleProxyRegistrationInquiry3.json"
		// 	case "6232345600":
		// 		fileName = "sampleProxyRegistrationInquiry4.json"
		// 	case "6242345600":
		// 		fileName = "sampleProxyRegistrationInquiry5.json"
		// 	case "6252345808":
		// 		fileName = "sampleProxyRegistrationInquiry6.json"
		// 	case "6262345806":
		// 		fileName = "sampleProxyRegistrationInquiry7.json"
		// 	}
		// case "621":
		// 	fileName = "sampleProxyRegistrationInquiry621.json"
		// 	fmt.Println("621")
		// case "622":
		// 	fileName = "sampleProxyRegistrationInquiry622.json"
		// 	fmt.Println("622")
		// // =========================================================================

		// case "710":
		// 	fileName = "sampleRegisterNewProxy.json"
		// 	fmt.Println("710")

		// //#################### Proxy Maintenance ###################################
		// case "720":

		// 	SdAccNum := *document.Message.SupplementaryData[0].Envlp.Dtl.Cstmr.Id
		// 	switch SdAccNum {
		// 	case "7202345600":
		// 		fileName = "sampleProxyMaintenance.json"
		// 	case "7212345101":
		// 		fileName = "sampleProxyMaintenance2.json"
		// 	}
		// case "721":
		// 	fileName = "sampleProxyMaintenance721.json"
		// 	fmt.Println("721")
		//============================================================================
	}

	fileName = "xml/" + fileName
	fmt.Println("filename:", fileName)

	file, _ := os.Open(fileName)
	defer file.Close()

	response, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	responseFormatter(w, response, 200)
}

func responseFormatter(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func responseFormatter2(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(statusCode)
	w.Write(data)
}
