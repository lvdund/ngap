package ngap

import (
	"bytes"
	"ngap/ie"
)

const (
	NgapPresentNothing uint8 = iota
	NgapPduInitiatingMessage
	NgapPduSuccessfulOutcome
	NgapPduUnsuccessfulOutcome
)

const (
	CriticalityReject uint8 = iota
	CriticalityNotify
	CriticalityIgnore
)

type NgapIE interface {
	Encode(aper.Writer) error
	Decode(aper.Reader) error
}

type NgapMessageIE struct {
	Id          ProtocolIeId
	Criticality Criticality
	Value       NgapIE
}

type NgapPdu struct {
	Present uint8
	Message NgapMessage
}

//represent InitiatingMessage, SuccessfulOutcome or UnsuccessfulOutcome
type NgapMessage struct {
	ProcedureCode ProcedureCode
	Criticality   Criticality
	Value         IeContainer //contain all IEs of an Ngap message
}

type IeContainer struct {
	IEs []NgapMessageIE
}

//refer to free5gc's aper to know how to do encoding/decoding the list of IEs
//it seems that the list is encoded first into a byte array, the array
//is encode as a octetstring (or bitstring)
func (c *IeContainer) Encode(w aper.AperWriter) (err error) {
	//TODO:
	return
}

func (c *IeContainer) Decode(w aper.AperReader) (err error) {
	//TODO:
	return
}

//and ngap example message
type NGSetupRequest struct {
	GlobalRanNodeId        ie.GlobalRanNodeId
	RanNodeName            []byte
	SupportedTaList        []ie.SupportedTaItem
	DefaultPagingDrx       *ie.PagingDrx
	UeRetentionInformation *ie.UeRetentionInformation
	NbIotDefaultPagingDrx  *ie.NbIotDefaultPagingDrx
}

//load message content from decoded IEs of an NgapPdu (code should be generated
//from spec)
//NOTE: criticality of IEs should be handled here, we may need to return a
//CriticalityDianosis object to report to the method caller

func (msg *NGSetupRequest) Load(ies []NgapMessageIE) (err error) {
	//fill data structure fields with IEs
	//check presence (mandotory fields)
	return
}

//create list of IEs for encoding (code should be generated from spec)
func (msg *NGSetupRequest) toIes() (ies []NgapMessageIE) {
	//turn data structure fields into IEs

	//GlobalRanNodeId
	ies = ies.append(ies, NgapMessageIE{
		Id:          ProtocolIeIdGlobalRanNodeId,
		Criticality: CriticalityReject,
		Value:       &msg.GlobalRanNodeId,
	})
	//TODO: add other fields

	return
}

//create NgapPdu for encoding (code should be generated from spec)
func (msg *NGSetupRequest) toNgapPdu() (pdu NgapPdu) {
	pdu = NgapPdu{
		Present: NgapInitiatingMessage, //predefined from spec
	}
	pdu.Message = &NgapMessage{
		ProcedureCode: NgapProcedureNGSetup,
		Criticality:   CriticalityReject, //parse from spec
		Value: NgapIeContainer{
			IEs: msg.toIes(),
		},
	}
	return
}

//API for decode NgapPdu
//Note that after decoding, we must inspect the procedure code and message
//present value to determine the embeding message. After that we should create
//the message and load it content from IEs container in the decoded NgapPdu

func NgapDecode(wire []byte) (pdu NgapPdu, err error) {
	r := aper.NewReader(bytes.NewReader(wire))
	//1. decode extention bit
	//2. decode present
	//3. decode procedure code
	//4. decode criticality
	//5. decode all IEs
	err = pdu.Message.IEs.Decode(r)
	return
}

//API for encode NgapPdu
func NgapEncode(pdu *NgapPdu) (wire []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(buf)
	//1. write extention bit
	//2. write present
	//3. write procedure code
	//4. write criticality
	//5. write all IEs
	if err = pdu.Message.IEs.Encode(w); err != nil {
		return
	}
	wire = w.Bytes()
	return
}

//example how to handle a ngap message
func HandleNgap(wire []byte) {
	if pdu, err = NgapDecode(wire); err != nil {
		return
	}
	switch pdu.Present {
	case NgapPduInitiatingMessage:
		handleInitiatingMessage(&pdu)
	case NgapPduSuccessfulOutcome:
		handleSuccessfulOutcome(&pdu)
	case NgapPduUnsuccessfulOutcome:
	}
}

func handleInitiatingMessage(pdu *NgapPdu) {
	switch pdu.Message.ProcedureCode {
	case NgapProcedureNGSetup: //NGSetupRequest
		msg := new(NGSetupRequest)
		_ = msg.Load(pdu) //load decoded NgapPdu into message data structure
		//then handle the message
	default:
		//TODO:
	}
}

func handleSuccessfulOutcome(pdu *NgapPdu) {
	switch pdu.Message.ProcedureCode {
	case NgapProcedureNGSetup: //NGSetupResponse
		msg := new(NGSetupResponse)
		_ = msg.Load(pdu) //load decoded NgapPdu into message data structure
		//then handle the message
	default:
		//TODO:
	}
}
