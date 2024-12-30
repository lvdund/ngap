package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type NGReset struct {
	Cause     Cause
	ResetType ResetType
}

func (msg *NGReset) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_NGReset, Criticality_PresentReject, msg.toIes())
}
func (msg *NGReset) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_Cause},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.Cause,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ResetType},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.ResetType,
	})
	return
}
func (msg *NGReset) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *NGReset) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_ResetType:
		var tmp ResetType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ResetType = tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
