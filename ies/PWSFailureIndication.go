package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PWSFailureIndication struct {
	PWSFailedCellIDList *PWSFailedCellIDList `,reject,mandatory`
	GlobalRANNodeID     *GlobalRANNodeID     `,reject,mandatory`
}

func (msg *PWSFailureIndication) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PWSFailureIndication, Criticality_PresentIgnore, msg.toIes())
}
func (msg *PWSFailureIndication) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.PWSFailedCellIDList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PWSFailedCellIDList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PWSFailedCellIDList})
	}
	if msg.GlobalRANNodeID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.GlobalRANNodeID})
	}
	return
}
func (msg *PWSFailureIndication) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PWSFailureIndication) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PWSFailedCellIDList:
		var tmp PWSFailedCellIDList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PWSFailedCellIDList = &tmp
	case ProtocolIEID_GlobalRANNodeID:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.GlobalRANNodeID = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
