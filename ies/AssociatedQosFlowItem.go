package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AssociatedQosFlowItem struct {
	QosFlowIdentifier        int64                     `lb:0,ub:63,madatory,valExt`
	QosFlowMappingIndication *QosFlowMappingIndication `optional`
	// IEExtensions *AssociatedQosFlowItemExtIEs `optional`
}

func (ie *AssociatedQosFlowItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.QosFlowMappingIndication != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_QosFlowIdentifier := NewINTEGER(ie.QosFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, true)
	if err = tmp_QosFlowIdentifier.Encode(w); err != nil {
		err = utils.WrapError("Encode QosFlowIdentifier", err)
		return
	}
	if ie.QosFlowMappingIndication != nil {
		if err = ie.QosFlowMappingIndication.Encode(w); err != nil {
			err = utils.WrapError("Encode QosFlowMappingIndication", err)
			return
		}
	}
	return
}
func (ie *AssociatedQosFlowItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_QosFlowIdentifier := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 63},
		ext: true,
	}
	if err = tmp_QosFlowIdentifier.Decode(r); err != nil {
		err = utils.WrapError("Read QosFlowIdentifier", err)
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(QosFlowMappingIndication)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read QosFlowMappingIndication", err)
			return
		}
		ie.QosFlowMappingIndication = tmp
	}
	return
}
