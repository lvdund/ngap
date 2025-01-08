package ies

import "github.com/lvdund/ngap/aper"

type AdditionalDLUPTNLInformationForHOItem struct {
	AdditionalDLNGUUPTNLInformation        UPTransportLayerInformation
	AdditionalQosFlowSetupResponseList     []QosFlowItemWithDataForwarding
	AdditionalDLForwardingUPTNLInformation *UPTransportLayerInformation `optional`
	// IEExtensions *AdditionalDLUPTNLInformationForHOItemExtIEs `optional`
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
	if err = ie.AdditionalDLNGUUPTNLInformation.Encode(w); err != nil {
		return
	}
	if len(ie.AdditionalQosFlowSetupResponseList) > 0 {
		tmp := Sequence[*QosFlowItemWithDataForwarding]{
			Value: []*QosFlowItemWithDataForwarding{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.AdditionalQosFlowSetupResponseList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
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
	if err = ie.AdditionalDLNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	tmp_AdditionalQosFlowSetupResponseList := Sequence[*QosFlowItemWithDataForwarding]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	fn := func() *QosFlowItemWithDataForwarding { return new(QosFlowItemWithDataForwarding) }
	if err = tmp_AdditionalQosFlowSetupResponseList.Decode(r, fn); err != nil {
		return
	}
	ie.AdditionalQosFlowSetupResponseList = []QosFlowItemWithDataForwarding{}
	for _, i := range tmp_AdditionalQosFlowSetupResponseList.Value {
		ie.AdditionalQosFlowSetupResponseList = append(ie.AdditionalQosFlowSetupResponseList, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.AdditionalDLForwardingUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
