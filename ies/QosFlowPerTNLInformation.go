package ies

import "github.com/lvdund/ngap/aper"

type QosFlowPerTNLInformation struct {
	UPTransportLayerInformation *UPTransportLayerInformation `False,`
	AssociatedQosFlowList       *AssociatedQosFlowList       `False,`
	// IEExtensions QosFlowPerTNLInformationExtIEs `False,OPTIONAL`
}

func (ie *QosFlowPerTNLInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.UPTransportLayerInformation != nil {
		if err = ie.UPTransportLayerInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.AssociatedQosFlowList != nil {
		if err = ie.AssociatedQosFlowList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QosFlowPerTNLInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.UPTransportLayerInformation = new(UPTransportLayerInformation)
	ie.AssociatedQosFlowList = new(AssociatedQosFlowList)
	if err = ie.UPTransportLayerInformation.Decode(r); err != nil {
		return
	}
	if err = ie.AssociatedQosFlowList.Decode(r); err != nil {
		return
	}
	return
}
