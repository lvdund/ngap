package ies

import "github.com/lvdund/ngap/aper"

type QosFlowAddOrModifyResponseItem struct {
	QosFlowIdentifier *QosFlowIdentifier `False,`
	// IEExtensions QosFlowAddOrModifyResponseItemExtIEs `False,OPTIONAL`
}

func (ie *QosFlowAddOrModifyResponseItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.QosFlowIdentifier != nil {
		if err = ie.QosFlowIdentifier.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QosFlowAddOrModifyResponseItem) Decode(r *aper.AperReader) (err error) {
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
	return
}
