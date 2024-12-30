package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type NGSetupRequest struct {
	GlobalRANNodeID        GlobalRANNodeID
	RANNodeName            *[]byte
	SupportedTAList        []SupportedTAItem
	DefaultPagingDRX       PagingDRX
	UERetentionInformation *UERetentionInformation
}

func (msg *NGSetupRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_NGSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *NGSetupRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.GlobalRANNodeID,
	})
	if msg.RANNodeName != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANNodeName},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 150},
				ext:   true,
				Value: *msg.RANNodeName,
			}})
	}
	tmp_SupportedTAList := Sequence[*SupportedTAItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofTACs},
		ext: false,
	}
	for _, i := range msg.SupportedTAList {
		tmp_SupportedTAList.Value = append(tmp_SupportedTAList.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SupportedTAList},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_SupportedTAList,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_DefaultPagingDRX},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.DefaultPagingDRX,
	})
	if msg.UERetentionInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERetentionInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERetentionInformation,
		})
	}
	return
}
func (msg *NGSetupRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *NGSetupRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false); err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	ieR := aper.NewReader(bytes.NewReader(buf))
	switch msgIe.Id.Value {
	case ProtocolIEID_GlobalRANNodeID:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.GlobalRANNodeID = tmp
	case ProtocolIEID_RANNodeName:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.RANNodeName = tmp.Value
	case ProtocolIEID_SupportedTAList:
		tmp := Sequence[*SupportedTAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTACs},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SupportedTAList = []SupportedTAItem{}
		for _, i := range tmp.Value {
			msg.SupportedTAList = append(msg.SupportedTAList, *i)
		}
	case ProtocolIEID_DefaultPagingDRX:
		var tmp PagingDRX
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.DefaultPagingDRX = tmp
	case ProtocolIEID_UERetentionInformation:
		var tmp UERetentionInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UERetentionInformation = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
