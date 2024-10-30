package ies

import "github.com/lvdund/ngap/aper"

type QosFlowInformationItem struct {
	QosFlowIdentifier *QosFlowIdentifier `False,`
	DLForwarding      *DLForwarding      `False,OPTIONAL`
	// IEExtensions QosFlowInformationItemExtIEs `False,OPTIONAL`
}

func (ie *QosFlowInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLForwarding != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.QosFlowIdentifier != nil {
		if err = ie.QosFlowIdentifier.Encode(w); err != nil {
			return
		}
	}
	if ie.DLForwarding != nil {
		if err = ie.DLForwarding.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QosFlowInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.QosFlowIdentifier = new(QosFlowIdentifier)
	ie.DLForwarding = new(DLForwarding)
	if err = ie.QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.DLForwarding.Decode(r); err != nil {
			return
		}
	}
	return
}
