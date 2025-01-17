package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceModifyResponseTransfer struct {
	DLNGUUPTNLInformation                *UPTransportLayerInformation     `optional`
	ULNGUUPTNLInformation                *UPTransportLayerInformation     `optional`
	QosFlowAddOrModifyResponseList       []QosFlowAddOrModifyResponseItem `lb:1,ub:maxnoofQosFlows,optional`
	AdditionalDLQosFlowPerTNLInformation []QosFlowPerTNLInformationItem   `lb:1,ub:maxnoofMultiConnectivityMinusOne,optional`
	QosFlowFailedToAddOrModifyList       []QosFlowWithCauseItem           `lb:1,ub:maxnoofQosFlows,optional`
	// IEExtensions *PDUSessionResourceModifyResponseTransferExtIEs `optional`
}

func (ie *PDUSessionResourceModifyResponseTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLNGUUPTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ULNGUUPTNLInformation != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.QosFlowAddOrModifyResponseList != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.AdditionalDLQosFlowPerTNLInformation != nil {
		aper.SetBit(optionals, 4)
	}
	if ie.QosFlowFailedToAddOrModifyList != nil {
		aper.SetBit(optionals, 5)
	}
	w.WriteBits(optionals, 6)
	if ie.DLNGUUPTNLInformation != nil {
		if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode DLNGUUPTNLInformation", err)
			return
		}
	}
	if ie.ULNGUUPTNLInformation != nil {
		if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode ULNGUUPTNLInformation", err)
			return
		}
	}
	if len(ie.QosFlowAddOrModifyResponseList) > 0 {
		tmp := Sequence[*QosFlowAddOrModifyResponseItem]{
			Value: []*QosFlowAddOrModifyResponseItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowAddOrModifyResponseList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode QosFlowAddOrModifyResponseList", err)
			return
		}
	}
	if len(ie.AdditionalDLQosFlowPerTNLInformation) > 0 {
		tmp := Sequence[*QosFlowPerTNLInformationItem]{
			Value: []*QosFlowPerTNLInformationItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext:   false,
		}
		for _, i := range ie.AdditionalDLQosFlowPerTNLInformation {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode AdditionalDLQosFlowPerTNLInformation", err)
			return
		}
	}
	if len(ie.QosFlowFailedToAddOrModifyList) > 0 {
		tmp := Sequence[*QosFlowWithCauseItem]{
			Value: []*QosFlowWithCauseItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowFailedToAddOrModifyList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode QosFlowFailedToAddOrModifyList", err)
			return
		}
	}
	return
}
func (ie *PDUSessionResourceModifyResponseTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(6); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
			err = utils.WrapError("Read DLNGUUPTNLInformation", err)
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
			err = utils.WrapError("Read ULNGUUPTNLInformation", err)
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_QosFlowAddOrModifyResponseList := Sequence[*QosFlowAddOrModifyResponseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowAddOrModifyResponseItem { return new(QosFlowAddOrModifyResponseItem) }
		if err = tmp_QosFlowAddOrModifyResponseList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read QosFlowAddOrModifyResponseList", err)
			return
		}
		ie.QosFlowAddOrModifyResponseList = []QosFlowAddOrModifyResponseItem{}
		for _, i := range tmp_QosFlowAddOrModifyResponseList.Value {
			ie.QosFlowAddOrModifyResponseList = append(ie.QosFlowAddOrModifyResponseList, *i)
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_AdditionalDLQosFlowPerTNLInformation := Sequence[*QosFlowPerTNLInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		fn := func() *QosFlowPerTNLInformationItem { return new(QosFlowPerTNLInformationItem) }
		if err = tmp_AdditionalDLQosFlowPerTNLInformation.Decode(r, fn); err != nil {
			err = utils.WrapError("Read AdditionalDLQosFlowPerTNLInformation", err)
			return
		}
		ie.AdditionalDLQosFlowPerTNLInformation = []QosFlowPerTNLInformationItem{}
		for _, i := range tmp_AdditionalDLQosFlowPerTNLInformation.Value {
			ie.AdditionalDLQosFlowPerTNLInformation = append(ie.AdditionalDLQosFlowPerTNLInformation, *i)
		}
	}
	if aper.IsBitSet(optionals, 5) {
		tmp_QosFlowFailedToAddOrModifyList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowWithCauseItem { return new(QosFlowWithCauseItem) }
		if err = tmp_QosFlowFailedToAddOrModifyList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read QosFlowFailedToAddOrModifyList", err)
			return
		}
		ie.QosFlowFailedToAddOrModifyList = []QosFlowWithCauseItem{}
		for _, i := range tmp_QosFlowFailedToAddOrModifyList.Value {
			ie.QosFlowFailedToAddOrModifyList = append(ie.QosFlowFailedToAddOrModifyList, *i)
		}
	}
	return
}
