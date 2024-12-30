package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type AMFConfigurationUpdateFailure struct {
	Cause                  Cause
	TimeToWait             *TimeToWait
	CriticalityDiagnostics *CriticalityDiagnostics
}

func (msg *AMFConfigurationUpdateFailure) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduUnsuccessfulOutcome, ProcedureCode_AMFConfigurationUpdate, Criticality_PresentReject, msg.toIes())
}
func (msg *AMFConfigurationUpdateFailure) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_Cause},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.Cause,
	})
	if msg.TimeToWait != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TimeToWait},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.TimeToWait,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return
}
func (msg *AMFConfigurationUpdateFailure) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *AMFConfigurationUpdateFailure) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_Cause:
		var tmp Cause
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.Cause = tmp
	case ProtocolIEID_TimeToWait:
		var tmp TimeToWait
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TimeToWait = &tmp
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CriticalityDiagnostics = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
