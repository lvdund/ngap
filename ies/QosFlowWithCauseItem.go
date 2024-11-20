package ies

import "github.com/lvdund/ngap/aper"

type QosFlowWithCauseItem struct {
	QosFlowIdentifier *QosFlowIdentifier `False,`
	Cause             *Cause             `False,`
	// IEExtensions QosFlowWithCauseItemExtIEs `False,OPTIONAL`
}

func (ie *QosFlowWithCauseItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.QosFlowIdentifier != nil {
		if err = ie.QosFlowIdentifier.Encode(w); err != nil {
			return
		}
	}
	if ie.Cause != nil {
		if err = ie.Cause.Encode(w); err != nil {
			return
		}
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
	ie.QosFlowIdentifier = new(QosFlowIdentifier)
	ie.Cause = new(Cause)
	if err = ie.QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	if err = ie.Cause.Decode(r); err != nil {
		return
	}
	return
}
