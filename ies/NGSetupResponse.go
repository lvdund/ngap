package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type NGSetupResponse struct {
	AMFName                []byte
	ServedGUAMIList        []ServedGUAMIItem
	RelativeAMFCapacity    int64
	PLMNSupportList        []PLMNSupportItem
	CriticalityDiagnostics *CriticalityDiagnostics
	UERetentionInformation *UERetentionInformation
}

func (msg *NGSetupResponse) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_NGSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *NGSetupResponse) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFName},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 1, Ub: 150},
			ext:   true,
			Value: msg.AMFName,
		}})
	tmp_ServedGUAMIList := Sequence[*ServedGUAMIItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofServedGUAMIs},
		ext: false,
	}
	for _, i := range msg.ServedGUAMIList {
		tmp_ServedGUAMIList.Value = append(tmp_ServedGUAMIList.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ServedGUAMIList},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_ServedGUAMIList,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RelativeAMFCapacity},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 255},
			ext:   false,
			Value: aper.Integer(msg.RelativeAMFCapacity),
		}})
	tmp_PLMNSupportList := Sequence[*PLMNSupportItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPLMNs},
		ext: false,
	}
	for _, i := range msg.PLMNSupportList {
		tmp_PLMNSupportList.Value = append(tmp_PLMNSupportList.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PLMNSupportList},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_PLMNSupportList,
	})
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	if msg.UERetentionInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERetentionInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERetentionInformation,
		})
	}
	return
}
func (msg *NGSetupResponse) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *NGSetupResponse) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFName:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFName = tmp.Value
	case ProtocolIEID_ServedGUAMIList:
		tmp := Sequence[*ServedGUAMIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofServedGUAMIs},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ServedGUAMIList = []ServedGUAMIItem{}
		for _, i := range tmp.Value {
			msg.ServedGUAMIList = append(msg.ServedGUAMIList, *i)
		}
	case ProtocolIEID_RelativeAMFCapacity:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RelativeAMFCapacity = int64(tmp.Value)
	case ProtocolIEID_PLMNSupportList:
		tmp := Sequence[*PLMNSupportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPLMNs},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PLMNSupportList = []PLMNSupportItem{}
		for _, i := range tmp.Value {
			msg.PLMNSupportList = append(msg.PLMNSupportList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CriticalityDiagnostics = &tmp
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
