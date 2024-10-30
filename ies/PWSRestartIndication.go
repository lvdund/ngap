package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PWSRestartIndication struct {
	CellIDListForRestart          *CellIDListForRestart          `,reject,mandatory`
	GlobalRANNodeID               *GlobalRANNodeID               `,reject,mandatory`
	TAIListForRestart             *TAIListForRestart             `,reject,mandatory`
	EmergencyAreaIDListForRestart *EmergencyAreaIDListForRestart `,reject,optional`
}

func (msg *PWSRestartIndication) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PWSRestartIndication, Criticality_PresentIgnore, msg.toIes())
}
func (msg *PWSRestartIndication) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.CellIDListForRestart != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CellIDListForRestart},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.CellIDListForRestart})
	}
	if msg.GlobalRANNodeID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.GlobalRANNodeID})
	}
	if msg.TAIListForRestart != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TAIListForRestart},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.TAIListForRestart})
	}
	if msg.EmergencyAreaIDListForRestart != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_EmergencyAreaIDListForRestart},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.EmergencyAreaIDListForRestart})
	}
	return
}
func (msg *PWSRestartIndication) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PWSRestartIndication) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_CellIDListForRestart:
		var tmp CellIDListForRestart
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CellIDListForRestart = &tmp
	case ProtocolIEID_GlobalRANNodeID:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.GlobalRANNodeID = &tmp
	case ProtocolIEID_TAIListForRestart:
		var tmp TAIListForRestart
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TAIListForRestart = &tmp
	case ProtocolIEID_EmergencyAreaIDListForRestart:
		var tmp EmergencyAreaIDListForRestart
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.EmergencyAreaIDListForRestart = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
