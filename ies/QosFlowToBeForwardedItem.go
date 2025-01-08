package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier int64
	// IEExtensions *QosFlowToBeForwardedItemExtIEs `optional`
}

func (ie *QosFlowToBeForwardedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_QosFlowIdentifier := NewINTEGER(ie.QosFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_QosFlowIdentifier.Encode(w); err != nil {
		err = utils.WrapError("Read QosFlowIdentifier", err)
		return
	}
	return
}
func (ie *QosFlowToBeForwardedItem) Decode(r *aper.AperReader) (err error) {
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
		err = utils.WrapError("Read QosFlowIdentifier", err)
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	return
}
