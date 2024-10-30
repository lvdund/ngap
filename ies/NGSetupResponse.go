package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type NGSetupResponse struct {
	AMFName                *AMFName                `,reject,mandatory`
	ServedGUAMIList        *ServedGUAMIList        `,reject,mandatory`
	RelativeAMFCapacity    *RelativeAMFCapacity    `,ignore,mandatory`
	PLMNSupportList        *PLMNSupportList        `,reject,mandatory`
	CriticalityDiagnostics *CriticalityDiagnostics `,ignore,optional`
	UERetentionInformation *UERetentionInformation `,ignore,optional`
}

func (msg *NGSetupResponse) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_NGSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *NGSetupResponse) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFName != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFName},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFName})
	}
	if msg.ServedGUAMIList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServedGUAMIList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ServedGUAMIList})
	}
	if msg.RelativeAMFCapacity != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RelativeAMFCapacity},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RelativeAMFCapacity})
	}
	if msg.PLMNSupportList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PLMNSupportList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PLMNSupportList})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics})
	}
	if msg.UERetentionInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UERetentionInformation},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UERetentionInformation})
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
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false)
	if err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	var buf []byte
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	ieR := aper.NewReader(bytes.NewReader(buf))
	switch msgIe.Id.Value {
	case ProtocolIEID_AMFName:
		var tmp AMFName
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFName = &tmp
	case ProtocolIEID_ServedGUAMIList:
		var tmp ServedGUAMIList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ServedGUAMIList = &tmp
	case ProtocolIEID_RelativeAMFCapacity:
		var tmp RelativeAMFCapacity
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RelativeAMFCapacity = &tmp
	case ProtocolIEID_PLMNSupportList:
		var tmp PLMNSupportList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PLMNSupportList = &tmp
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
