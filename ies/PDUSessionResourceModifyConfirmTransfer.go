package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList      []QosFlowModifyConfirmItem            `lb:1,ub:maxnoofQosFlows,madatory`
	ULNGUUPTNLInformation         UPTransportLayerInformation           `madatory`
	AdditionalNGUUPTNLInformation []UPTransportLayerInformationPairItem `lb:1,ub:maxnoofMultiConnectivityMinusOne,optional`
	QosFlowFailedToModifyList     []QosFlowWithCauseItem                `lb:1,ub:maxnoofQosFlows,optional`
	// IEExtensions *PDUSessionResourceModifyConfirmTransferExtIEs `optional`
}

func (ie *PDUSessionResourceModifyConfirmTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AdditionalNGUUPTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.QosFlowFailedToModifyList != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if len(ie.QosFlowModifyConfirmList) > 0 {
		tmp := Sequence[*QosFlowModifyConfirmItem]{
			Value: []*QosFlowModifyConfirmItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowModifyConfirmList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode QosFlowModifyConfirmList", err)
			return
		}
	} else {
		err = utils.WrapError("QosFlowModifyConfirmList is nil", err)
		return
	}
	if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode ULNGUUPTNLInformation", err)
		return
	}
	if len(ie.AdditionalNGUUPTNLInformation) > 0 {
		tmp := Sequence[*UPTransportLayerInformationPairItem]{
			Value: []*UPTransportLayerInformationPairItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext:   false,
		}
		for _, i := range ie.AdditionalNGUUPTNLInformation {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode AdditionalNGUUPTNLInformation", err)
			return
		}
	}
	if len(ie.QosFlowFailedToModifyList) > 0 {
		tmp := Sequence[*QosFlowWithCauseItem]{
			Value: []*QosFlowWithCauseItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowFailedToModifyList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode QosFlowFailedToModifyList", err)
			return
		}
	}
	return
}
func (ie *PDUSessionResourceModifyConfirmTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	tmp_QosFlowModifyConfirmList := Sequence[*QosFlowModifyConfirmItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	fn := func() *QosFlowModifyConfirmItem { return new(QosFlowModifyConfirmItem) }
	if err = tmp_QosFlowModifyConfirmList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read QosFlowModifyConfirmList", err)
		return
	}
	ie.QosFlowModifyConfirmList = []QosFlowModifyConfirmItem{}
	for _, i := range tmp_QosFlowModifyConfirmList.Value {
		ie.QosFlowModifyConfirmList = append(ie.QosFlowModifyConfirmList, *i)
	}
	if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
		err = utils.WrapError("Read ULNGUUPTNLInformation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_AdditionalNGUUPTNLInformation := Sequence[*UPTransportLayerInformationPairItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		fn := func() *UPTransportLayerInformationPairItem { return new(UPTransportLayerInformationPairItem) }
		if err = tmp_AdditionalNGUUPTNLInformation.Decode(r, fn); err != nil {
			err = utils.WrapError("Read AdditionalNGUUPTNLInformation", err)
			return
		}
		ie.AdditionalNGUUPTNLInformation = []UPTransportLayerInformationPairItem{}
		for _, i := range tmp_AdditionalNGUUPTNLInformation.Value {
			ie.AdditionalNGUUPTNLInformation = append(ie.AdditionalNGUUPTNLInformation, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_QosFlowFailedToModifyList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowWithCauseItem { return new(QosFlowWithCauseItem) }
		if err = tmp_QosFlowFailedToModifyList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read QosFlowFailedToModifyList", err)
			return
		}
		ie.QosFlowFailedToModifyList = []QosFlowWithCauseItem{}
		for _, i := range tmp_QosFlowFailedToModifyList.Value {
			ie.QosFlowFailedToModifyList = append(ie.QosFlowFailedToModifyList, *i)
		}
	}
	return
}
