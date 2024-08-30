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

type ProcedureCode uint8

//add AperEncode/AperDecode method

type NgapIE interface {
	Encode(aper.Writer) error
	Decode(aper.Reader) error
}

//represent an IE in Ngap messages
type NgapMessageIE struct {
	Id          ProtocolIeId //protocol IE identity
	Criticality Criticality
	Value       NgapIE //open type
}

func (ie *NgapMessageIE) Encode(w aper.AperWriter) (err error) {
	//TODO:
	//1. TODO: encode protocol Ie Id
	//2. TODO: encode criticality
	//3. encode NgapIE
	//encode IE into a byte array first
	var buf bytes.Buffer
	ieW := aper.NewWriter(buf)
	if err = ie.Value.Encode(ieW); err != nil {
		return
	}
	//then write the array as open type
	err = w.WriteOpenType(buf.Bytes)
	return
}

//encode a sequence of Ngap message IE (IEs container)
func encodeIes(ies []NgapMessageIE) (wire []byte, err error) {
	var buf bytes.Buffer
	w = aper.NewWriter(buf)
	numItems := len(ies)
	//1. TODO: write length of the list
	w.WriteLength(numItems) //TODO: check AperWriter's APIs

	//2. write every item on the list
	for _, ie := range ies {
		if err = ie.Encode(w); err != nil {
			return
		}
	}
	wire = buf.Bytes()
	return
}

//an ngap example message that implement MessageUnmarshaller and encoding
//method
type NGSetupRequest struct {
	GlobalRanNodeId        ie.GlobalRanNodeId
	RanNodeName            []byte
	SupportedTaList        []ie.SupportedTaItem
	DefaultPagingDrx       *ie.PagingDrx
	UeRetentionInformation *ie.UeRetentionInformation
	NbIotDefaultPagingDrx  *ie.NbIotDefaultPagingDrx
}

//implement MessageUnmarshaller (code should be generated from spec)
func (msg *NGSetupRequest) decode(wire []byte) (err error, diagList []IeDiagnosticsItem) {
	r = aper.NewReader(wire)
	//fill data structure fields with IEs
	//1. get num of IEs
	var numIEs int
	if numIEs, err = r.ReadConstrainNumber(); err != nil { //TODO: check AperReader's APIs
		return
	}
	//2. decode all IEs
	ies := make([]NgapMessageIE, numIEs)
	for i := 0; i < numIEs; i++ {
		if ies, ieErr = msg.decodeIE(r); ieErr != nil {
			//depending on error/criticality value, we may continue or abort (check spec again)
		}
	}

	//NOTE: after decode all IEs, now let's assign them to the message fields.
	//Alternatively, we can also assign the fields while decoding Ie in the previous
	//step

	//TODO:check presence (mandatory fields)
	//TODO: check for duplicated IEs
	return
}

//decode a single IE in the message (code should be generated from spec
func (msg *NGSetupRequest) decodeIE(r aper.AperReader) (*NgapMessageIE, error) {
	var ie NgapMessageIE
	//1. decode protocol Ie Id
	ie.Id = 0 //TODO: read from reader
	//2. decode criticality
	ie.Criticality = CriticalityReject //TODO: read from reader

	//3. decode NgapIE
	var buf []byte
	//read IE byte array
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	//prepare IE data structure for decoding
	switch ie.Id { //list of cases are generated from spec
	case ProtocolIdRanNodeName:
		ie.Value = new(ie.RanNodeName) //prepare IE for decoding
		//TODO: other cases
	default:
		//TODO: set an error
		err = fmt.Errorf("temporary error")
		return
	}

	//then decode IE
	ieReader := aper.NewReader(bytes.NewReader(buf))
	err = ie.Value.Decode(ieReader)
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

//write message to AperWriter (code should be generated from spec)
func (msg *NGSetupRequest) Encode(w aper.AperWriter) (err error) {
	procedureCode := NgapProcedureNGSetup
	criticality := CriticalityReject //parse from spec
	present := NgapInitiatingMessage //predefined from spec
	ies := msg.toIes()
	//1. TODO: write present
	//2. TODO:write procedure code
	//3. TODO: write criticality
	//4. write message content
	var containerBytes []byte
	//first encode message content into byte array
	if err, containerBytes = encodeIes(ies); err != nil {
		return
	}
	//then write the byte array
	if err = w.WriteOpenType(containerBytes); err != nil { //write OpenType bytes
		return
	}

	return
}
