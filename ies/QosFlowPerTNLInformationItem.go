package ies

import "github.com/lvdund/ngap/aper"

type QosFlowPerTNLInformationItem struct {
	QosFlowPerTNLInformation *QosFlowPerTNLInformation `True,`
	// IEExtensions QosFlowPerTNLInformationItemExtIEs `False,OPTIONAL`
}

func (ie *QosFlowPerTNLInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.QosFlowPerTNLInformation != nil {
		if err = ie.QosFlowPerTNLInformation.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QosFlowPerTNLInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.QosFlowPerTNLInformation = new(QosFlowPerTNLInformation)
	if err = ie.QosFlowPerTNLInformation.Decode(r); err != nil {
		return
	}
	return
}
