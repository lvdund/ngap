package ies

import "github.com/lvdund/ngap/aper"

type AdditionalDLUPTNLInformationForHOItem struct {
	AdditionalDLNGUUPTNLInformation        *UPTransportLayerInformation   `False,`
	AdditionalQosFlowSetupResponseList     *QosFlowListWithDataForwarding `False,`
	AdditionalDLForwardingUPTNLInformation *UPTransportLayerInformation   `False,OPTIONAL`
	// IEExtensions AdditionalDLUPTNLInformationForHOItemExtIEs `False,OPTIONAL`
}

func (ie *AdditionalDLUPTNLInformationForHOItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AdditionalDLForwardingUPTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.AdditionalDLNGUUPTNLInformation != nil {
		if err = ie.AdditionalDLNGUUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.AdditionalQosFlowSetupResponseList != nil {
		if err = ie.AdditionalQosFlowSetupResponseList.Encode(w); err != nil {
			return
		}
	}
	if ie.AdditionalDLForwardingUPTNLInformation != nil {
		if err = ie.AdditionalDLForwardingUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *AdditionalDLUPTNLInformationForHOItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.AdditionalDLNGUUPTNLInformation = new(UPTransportLayerInformation)
	ie.AdditionalQosFlowSetupResponseList = new(QosFlowListWithDataForwarding)
	ie.AdditionalDLForwardingUPTNLInformation = new(UPTransportLayerInformation)
	if err = ie.AdditionalDLNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if err = ie.AdditionalQosFlowSetupResponseList.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.AdditionalDLForwardingUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
