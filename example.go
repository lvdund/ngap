package ngap

import (
	"bytes"
	"fmt"
	"io"
	"ngap/aper"
	"ngap/ie"

	"github.com/sirupsen/logrus"
)

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
//it decodes the list of IEs for an Ngap message
func (msg *NGSetupRequest) decode(wire []byte) (err error, diagList []ie.CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))

	//1. container of IE list have an extension bit
	r.ReadBool()

	//2. decode the inner IE list
	fmt.Printf("msg content = %.8b\n", wire)
	// //fill data structure fields with IEs
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, false); err != nil {
		return
	}

	logrus.Infoln("seems fine", len(ies))
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
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false)
	if err != nil {
		return
	}
	fmt.Printf("IeID=%v-%.8b\n", id, id)
	msgIe = new(NgapMessageIE)
	msgIe.Id.NgapProtocolIeId = aper.Integer(id)
	//2. decode criticality
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}

	fmt.Printf("Criticality=%v\n", c)
	msgIe.Criticality.Value = aper.Enumerated(c)
	//3. decode NgapIE
	var buf []byte
	//read IE byte array
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	fmt.Printf("IE=%.8b\n", buf)
	ieR := aper.NewReader(bytes.NewReader(buf))
	//prepare IE data structure for decoding
	switch msgIe.Id.NgapProtocolIeId { //list of cases are generated from spec
	case ie.ProtocolIEIDGlobalRANNodeID:
		if err = msg.GlobalRanNodeId.Decode(ieR); err != nil {
			return
		}
	case ie.ProtocolIEIDRANNodeName:
		var tmp ie.RANNodeName
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RanNodeName = &tmp
		fmt.Printf("=decode RanNodeName:%s-%.8b\n", msg.RanNodeName.Value, []byte{'a'})
	case ie.ProtocolIEIDSupportedTAList:
		// msg.SupportedTaList = append(msg.SupportedTaList)
	case ie.ProtocolIEIDDefaultPagingDRX:
		if err = msg.DefaultPagingDrx.Decode(ieR); err != nil {
			return
		}
	case ie.ProtocolIEIDUERetentionInformation:
		if err = msg.UeRetentionInformation.Decode(ieR); err != nil {
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
		/*
			@DUNG: should we use WriteSequenceOf ? using NgapMessageIE does not seem right to me (inner Item does not need protocol IeId and criticality), pls check the specs

				var SupportedTaList []NgapIE
				for _, ie := range msg.SupportedTaList {
					SupportedTaList = append(SupportedTaList, &ie)
				}
				ies = append(ies, NgapMessageIE{
					Id:          ie.NgapProtocolIeId{NgapProtocolIeId: ie.ProtocolIEIDSupportedTAList},
					Criticality: ie.Criticality{Value: ie.CriticalityPresentReject},
					Value:       NewIEs(SupportedTaList),
				})
		*/
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

// write message to io.Writer (code should be generated from spec)
func (msg *NGSetupRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ie.ProcedureCodeNGSetup, ie.CriticalityPresentReject, msg.toIes())
}
