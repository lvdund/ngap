package ies

import "github.com/lvdund/ngap/aper"

type QosFlowWithCauseItem struct {
	QosFlowIdentifier int64
	Cause             Cause
	// IEExtensions *QosFlowWithCauseItemExtIEs `optional`
}

func (ie *QosFlowWithCauseItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_QosFlowIdentifier := NewINTEGER(ie.QosFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_QosFlowIdentifier.Encode(w); err != nil {
		return
	}
	if err = ie.Cause.Encode(w); err != nil {
		return
	}
	return
}
func (ie *QosFlowWithCauseItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_QosFlowIdentifier := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 63},
		ext: false,
	}
	if err = tmp_QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	if err = ie.Cause.Decode(r); err != nil {
		return
	}
	return
}
