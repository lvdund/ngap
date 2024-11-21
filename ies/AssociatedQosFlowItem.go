package ies

import "github.com/lvdund/ngap/aper"

type AssociatedQosFlowItem struct {
	QosFlowIdentifier        *QosFlowIdentifier `False,`
	QosFlowMappingIndication *aper.Enumerated   `True,OPTIONAL,0,1"`
	// IEExtensions AssociatedQosFlowItemExtIEs `False,OPTIONAL`
}

func (ie *AssociatedQosFlowItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 2)
	if ie.QosFlowIdentifier != nil {
		if err = ie.QosFlowIdentifier.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowMappingIndication != nil {
		err = w.WriteEnumerate(uint64(*ie.QosFlowMappingIndication), aper.Constraint{Lb: 0, Ub: 1}, true)
	}
	return
}
func (ie *AssociatedQosFlowItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.QosFlowIdentifier = new(QosFlowIdentifier)
	if err = ie.QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	v, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	*ie.QosFlowMappingIndication = aper.Enumerated(v)
	return
}
