package ngap

import (
	"bytes"
	"fmt"
	"ngap/aper"
	"ngap/ie"
)

const (
	NgapPresentNothing uint8 = iota
	NgapPduInitiatingMessage
	NgapPduSuccessfulOutcome
	NgapPduUnsuccessfulOutcome
)

type ProcedureCode uint8

//add AperEncode/AperDecode method

type NgapIE interface {
	Encode(aper.AperWriter) error
	Decode(aper.AperReader) error
}

// represent an IE in Ngap messages
type NgapMessageIE struct {
	Id          ie.NgapProtocolIeId //protocol IE identity
	Criticality ie.Criticality
	Value       NgapIE //open type
}

func (ie *NgapMessageIE) Encode(w aper.AperWriter) (err error) {
	//TODO:
	//1. TODO: encode protocol Ie Id
	//2. TODO: encode criticality
	if err = ie.Id.Encode(w); err != nil {
		return
	}
	if err = ie.Criticality.Encode(w); err != nil {
		return
	}
	//3. encode NgapIE
	//encode IE into a byte array first
	var buf bytes.Buffer
	ieW := aper.NewWriter(&buf)
	if err = ie.Value.Encode(ieW); err != nil {
		return
	}
	//then write the array as open type
	// err = w.WriteOpenType(buf.Bytes())
	err = ie.Value.Encode(w)
	return
}

// encode a sequence of Ngap message IE (IEs container)
func encodeIes(ies []NgapMessageIE) (wire []byte, err error) {
	var buff bytes.Buffer
	w := aper.NewWriter(&buff)
	// numItems := len(ies)
	//1. TODO: write length of the list
	// w.WriteLength(numItems) //TODO: check AperWriter's APIs

	//2. write every item on the list
	for _, ie := range ies {
		if err = ie.Encode(w); err != nil {
			return
		}
	}
	wire = buff.Bytes()
	return
}

// an ngap example message that implement MessageUnmarshaller and encoding
// method
type NGSetupRequest struct {
	GlobalRanNodeId        *ie.GlobalRanNodeId
	RanNodeName            []byte
	SupportedTaList        []ie.SupportedTaItem
	DefaultPagingDrx       *ie.PagingDrx
	UeRetentionInformation *ie.UeRetentionInformation
	// NbIotDefaultPagingDrx  *ie.NbIotDefaultPagingDrx
}

// implement MessageUnmarshaller (code should be generated from spec)
func (msg *NGSetupRequest) decode(wire []byte) (err error, diagList []ie.CriticalityDiagnostics) {
	// r := aper.NewReader(bytes.NewReader(wire))
	// //fill data structure fields with IEs
	// //1. get num of IEs
	// var numIEs int
	// if numIEs, err = r.ReadConstrainNumber(); err != nil { //TODO: check AperReader's APIs
	// 	return
	// }
	// //2. decode all IEs
	// // ies := make([]NgapMessageIE, numIEs)
	// for i := 0; i < numIEs; i++ {
	// 	if _, ieErr := msg.decodeIE(r); ieErr != nil {
	// 		//depending on error/criticality value, we may continue or abort (check spec again)
	// 	}
	// }
	// //NOTE: after decode all IEs, now let's assign them to the message fields.
	// //Alternatively, we can also assign the fields while decoding Ie in the previous
	// //step

	// //TODO:check presence (mandatory fields)
	// if msg.GlobalRanNodeId == nil || msg.SupportedTaList == nil || msg.DefaultPagingDrx == nil {
	// 	return
	// }
	// //TODO: check for duplicated IEs
	return
}

// decode a single IE in the message (code should be generated from spec
func (msg *NGSetupRequest) decodeIE(r aper.AperReader) (msgIe *NgapMessageIE, err error) {
	//1. decode protocol Ie Id
	id, err := r.ReadInteger(&aper.Constrain{Lb: 0, Ub: 65535}, true)
	if err != nil {
		return
	}
	msgIe.Id.NgapProtocolIeId = aper.Integer(id)
	//2. decode criticality
	c, err := r.ReadEnumerate(&aper.Constrain{Lb: 0, Ub: 2}, true)
	if err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	//3. decode NgapIE
	// var buf []byte
	//read IE byte array
	// if buf, err = r.ReadOpenType(); err != nil {
	// 	return
	// }
	// r := aper.NewReader(bytes.NewReader(buf))
	//prepare IE data structure for decoding
	switch msgIe.Id.NgapProtocolIeId { //list of cases are generated from spec
	case ie.ProtocolIEIDGlobalRANNodeID:
		if err = msg.GlobalRanNodeId.Decode(r); err != nil {
			return
		}
	case ie.ProtocolIEIDRANNodeName:
		msg.RanNodeName, _ = r.ReadOctetString(&aper.Constrain{Lb: 1, Ub: 150}, true)
	case ie.ProtocolIEIDSupportedTAList:
		// msg.SupportedTaList = append(msg.SupportedTaList)
	case ie.ProtocolIEIDDefaultPagingDRX:
		if err = msg.DefaultPagingDrx.Decode(r); err != nil {
			return
		}
	case ie.ProtocolIEIDUERetentionInformation:
		if err = msg.UeRetentionInformation.Decode(r); err != nil {
			return
		}
	// case 0:
	// 	if err = msg.NbIotDefaultPagingDrx.Decode(r); err != nil {
	// 		return
	// 	}
	default:
		err = fmt.Errorf("temporary error")
		return
	}

	//then decode IE
	// err = msgIe.Value.Decode(ieReader)
	return
}

// create list of IEs for encoding (code should be generated from spec)
func (msg *NGSetupRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	//GlobalRanNodeId
	if msg.GlobalRanNodeId != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDGlobalRANNodeID},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentReject},
			Value:       msg.GlobalRanNodeId,
		})
	}
	//RanNodeName
	if msg.RanNodeName != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDRANNodeName},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentIgnore},
			Value:       NewIE(msg.RanNodeName),
		})
	}
	//SupportedTaList
	if msg.SupportedTaList != nil {
		var SupportedTaList []NgapIE
		for _, ie := range msg.SupportedTaList {
			SupportedTaList = append(SupportedTaList, &ie)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDSupportedTAList},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentReject},
			Value:       NewIEs(SupportedTaList),
		})
	}
	//DefaultPagingDrx
	if msg.DefaultPagingDrx != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDDefaultPagingDRX},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentIgnore},
			Value:       msg.DefaultPagingDrx,
		})
	}
	//UeRetentionInformation
	if msg.UeRetentionInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDUERetentionInformation},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentIgnore},
			Value:       msg.UeRetentionInformation,
		})
	}
	//NbIotDefaultPagingDrx
	// ies = append(ies, NgapMessageIE{
	// 	Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDDefaultPagingDRX},
	// 	Criticality: ie.Criticality{Value: ie.CriticalityPresentIgnore},
	// 	Value:       msg.NbIotDefaultPagingDrx,
	// })

	return
}

// write message to AperWriter (code should be generated from spec)
func (msg *NGSetupRequest) Encode(w aper.AperWriter) (err error) {
	procedureCode := ie.ProcedureCodeNGSetup
	criticality := ie.CriticalityPresentReject           //parse from spec
	present := ie.InitiatingMessagePresentNGSetupRequest //predefined from spec
	ies := msg.toIes()
	if ies == nil {
		return fmt.Errorf("Cann not load NGSetupRequest")
	}
	//1. TODO: write present
	//2. TODO:write procedure code
	//3. TODO: write criticality
	//4. write message content
	if err = w.WriteInteger(int64(present), &aper.Constrain{Lb: 0, Ub: 255}, true); err != nil {
		return
	}
	if err = w.WriteInteger(int64(procedureCode), &aper.Constrain{Lb: 0, Ub: 255}, true); err != nil {
		return
	}
	if err = w.WriteEnumerate(uint64(criticality), aper.Constrain{Lb: 0, Ub: 255}, true); err != nil {
		return
	}
	var containerBytes []byte
	//first encode message content into byte array
	if containerBytes, err = encodeIes(ies); err != nil {
		return
	}
	//then write the byte array
	// if err = w.WriteOpenType(containerBytes); err != nil { //write OpenType bytes
	// 	return
	// }
	err = w.WriteBytes(containerBytes)

	return
}

func NgapEncode(pdu NgapPdu) (w aper.AperWriter, err error) {
	var buff bytes.Buffer
	w = aper.NewWriter(&buff)
	if err = w.WriteInteger(int64(pdu.Present), &aper.Constrain{Lb: 0, Ub: 2}, true); err != nil {
		return
	}
	if err = pdu.Message.Msg.Encode(w); err != nil {
		return
	}
	return
}
