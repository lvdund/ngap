package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AdditionalDLUPTNLInformationForHOItem struct {
	AdditionalDLNGUUPTNLInformation        UPTransportLayerInformation     `madatory`
	AdditionalQosFlowSetupResponseList     []QosFlowItemWithDataForwarding `lb:1,ub:maxnoofQosFlows,madatory`
	AdditionalDLForwardingUPTNLInformation *UPTransportLayerInformation    `optional`
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
		err = utils.WrapError("Encode AdditionalDLNGUUPTNLInformation", err)
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
			err = utils.WrapError("Encode AdditionalQosFlowSetupResponseList", err)
			return
		}
	} else {
		err = utils.WrapError("AdditionalQosFlowSetupResponseList is nil", err)
		return
	}
	if ie.AdditionalDLForwardingUPTNLInformation != nil {
		if err = ie.AdditionalDLForwardingUPTNLInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode AdditionalDLForwardingUPTNLInformation", err)
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
		err = utils.WrapError("Read AdditionalDLNGUUPTNLInformation", err)
		return
	}
	tmp_AdditionalQosFlowSetupResponseList := Sequence[*QosFlowItemWithDataForwarding]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	fn := func() *QosFlowItemWithDataForwarding { return new(QosFlowItemWithDataForwarding) }
	if err = tmp_AdditionalQosFlowSetupResponseList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read AdditionalQosFlowSetupResponseList", err)
		return
	}
	ie.AdditionalQosFlowSetupResponseList = []QosFlowItemWithDataForwarding{}
	for _, i := range tmp_AdditionalQosFlowSetupResponseList.Value {
		ie.AdditionalQosFlowSetupResponseList = append(ie.AdditionalQosFlowSetupResponseList, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(UPTransportLayerInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read AdditionalDLForwardingUPTNLInformation", err)
			return
		}
		ie.AdditionalDLForwardingUPTNLInformation = tmp
	}
	return
}
