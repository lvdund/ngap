package ies

import "github.com/lvdund/ngap/aper"

type AssociatedQosFlowItem struct {
	QosFlowIdentifier        int64
	QosFlowMappingIndication *int64
	// IEExtensions  *AssociatedQosFlowItemExtIEs
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
		return
	}
	if ie.QosFlowMappingIndication != nil {
		tmp_QosFlowMappingIndication := NewENUMERATED(*ie.QosFlowMappingIndication, aper.Constraint{Lb: 0, Ub: 0}, true)
		if err = tmp_QosFlowMappingIndication.Encode(w); err != nil {
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
		ext: false,
	}
	if err = tmp_QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_QosFlowMappingIndication := ENUMERATED{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_QosFlowMappingIndication.Decode(r); err != nil {
			return
		}
		*ie.QosFlowMappingIndication = int64(tmp_QosFlowMappingIndication.Value)
	}
	return
}
