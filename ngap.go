package ngap

import (
	"bytes"
	"fmt"
	"ngap/aper"
	"ngap/ie"

	"github.com/sirupsen/logrus"
)

const (
	NgapPresentNothing uint8 = iota
	NgapPduInitiatingMessage
	NgapPduSuccessfulOutcome
	NgapPduUnsuccessfulOutcome
)

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

func (ie NgapMessageIE) Encode(w aper.AperWriter) (err error) {
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
	err = w.WriteOpenType(buf.Bytes())
	// err = ie.Value.Encode(w)
	return
}

// encode a sequence of Ngap message IE (IEs container)
func encodeIes(ies []NgapMessageIE) (wire []byte, err error) {
	var buff bytes.Buffer
	w := aper.NewWriter(&buff)
	// Encoding Value Extensive Bit
	w.WriteBool(aper.Zero)
	if err = aper.WriteSequenceOf[NgapMessageIE](ies, w, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, true); err != nil {
		return
	}
	wire = buff.Bytes()
	return
}

// an ngap example message that implement MessageUnmarshaller and encoding
// method
type NGSetupRequest struct {
	GlobalRanNodeId        *ie.GlobalRanNodeId
	RanNodeName            *ie.RANNodeName
	SupportedTaList        []ie.SupportedTaItem
	DefaultPagingDrx       *ie.PagingDrx
	UeRetentionInformation *ie.UeRetentionInformation
	// NbIotDefaultPagingDrx  *ie.NbIotDefaultPagingDrx
}

// implement MessageUnmarshaller (code should be generated from spec)
func (msg *NGSetupRequest) decode(wire []byte) (err error, diagList []ie.CriticalityDiagnostics) {

	r := aper.NewReader(bytes.NewReader(wire))

	// //fill data structure fields with IEs
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, false); err != nil {
		return
	}
	_ = ies

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
	logrus.Infoln("Decode IE - NGSetupRequest msg========")
	//1. decode protocol Ie Id
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return
	}
	msgIe.Id.NgapProtocolIeId = aper.Integer(id)
	//2. decode criticality
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
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
		if err = msg.RanNodeName.Decode(r); err != nil {
			return
		}
		fmt.Println("=decode RanNodeName:", msg.RanNodeName.Value)
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
		fmt.Println("GlobalRanNodeId")
	}
	//RanNodeName
	if msg.RanNodeName != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDRANNodeName},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentIgnore},
			Value:       msg.RanNodeName,
		})
		fmt.Println("RanNodeName")
	}
	//SupportedTaList
	if len(msg.SupportedTaList) > 0 {
		var SupportedTaList []NgapIE
		for _, ie := range msg.SupportedTaList {
			SupportedTaList = append(SupportedTaList, &ie)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDSupportedTAList},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentReject},
			Value:       NewIEs(SupportedTaList),
		})
		fmt.Println("SupportedTaList")
	}
	//DefaultPagingDrx
	if msg.DefaultPagingDrx != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDDefaultPagingDRX},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentIgnore},
			Value:       msg.DefaultPagingDrx,
		})
		fmt.Println("DefaultPagingDrx")
	}
	//UeRetentionInformation
	if msg.UeRetentionInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDUERetentionInformation},
			Criticality: ie.Criticality{Value: ie.CriticalityPresentIgnore},
			Value:       msg.UeRetentionInformation,
		})
		fmt.Println("UeRetentionInformation")
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
	present := NgapPduInitiatingMessage //predefined from spec
	procedureCode := ie.ProcedureCode{Value: aper.Integer(ie.ProcedureCodeNGSetup)}
	criticality := ie.Criticality{Value: ie.CriticalityPresentReject} //parse from spec
	//1. TODO: write present
	//2. TODO:write procedure code
	//3. TODO: write criticality
	//4. write message content
	if err = w.WriteChoice(uint64(present), 2, true); err != nil {
		return
	}
	if err = procedureCode.Encode(w); err != nil {
		return
	}
	if err = criticality.Encode(w); err != nil {
		return
	}

	ies := msg.toIes()
	if len(ies) == 0 {
		fmt.Println("Can not load NGSetupRequest")
		return
	}
	var containerBytes []byte
	// buf := bytes.NewBuffer(containerBytes)
	// msgBytes := aper.NewWriter(buf)
	// // Encoding Value Extensive Bit
	// msgBytes.WriteBool(aper.Zero)
	// if err = w.WriteChoice(7, 2, false); err != nil {
	// 	return
	// }
	//first encode message content into byte array
	if containerBytes, err = encodeIes(ies); err != nil {
		return
	}
	//then write the byte array
	// if err = w.WriteOpenType(containerBytes); err != nil { //write OpenType bytes
	// 	return
	// }
	err = w.WriteOpenType(containerBytes)

	return
}

func NgapEncode(pdu NgapPdu) (w aper.AperWriter, err error) {
	var buff bytes.Buffer
	w = aper.NewWriter(&buff)
	// Encoding Value Extensive Bit
	w.WriteBool(aper.Zero)
	if err = pdu.Message.Msg.Encode(w); err != nil {
		return
	}
	w.WriteFlush()
	// fmt.Printf("\n\tEncoded: %v", w.GetBuf())
	return
}
