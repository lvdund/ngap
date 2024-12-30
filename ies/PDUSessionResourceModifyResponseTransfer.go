package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyResponseTransfer struct {
	DLNGUUPTNLInformation                *UPTransportLayerInformation
	ULNGUUPTNLInformation                *UPTransportLayerInformation
	QosFlowAddOrModifyResponseList       *[]QosFlowAddOrModifyResponseItem
	AdditionalDLQosFlowPerTNLInformation *[]QosFlowPerTNLInformationItem
	QosFlowFailedToAddOrModifyList       *[]QosFlowWithCauseItem
	// IEExtensions  *PDUSessionResourceModifyResponseTransferExtIEs
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
			return
		}
	}
	if ie.ULNGUUPTNLInformation != nil {
		if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowAddOrModifyResponseList != nil {
		if len(*ie.QosFlowAddOrModifyResponseList) > 0 {
			tmp := Sequence[*QosFlowAddOrModifyResponseItem]{
				Value: []*QosFlowAddOrModifyResponseItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range *ie.QosFlowAddOrModifyResponseList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.AdditionalDLQosFlowPerTNLInformation != nil {
		if len(*ie.AdditionalDLQosFlowPerTNLInformation) > 0 {
			tmp := Sequence[*QosFlowPerTNLInformationItem]{
				Value: []*QosFlowPerTNLInformationItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
				ext:   false,
			}
			for _, i := range *ie.AdditionalDLQosFlowPerTNLInformation {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.QosFlowFailedToAddOrModifyList != nil {
		if len(*ie.QosFlowFailedToAddOrModifyList) > 0 {
			tmp := Sequence[*QosFlowWithCauseItem]{
				Value: []*QosFlowWithCauseItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range *ie.QosFlowFailedToAddOrModifyList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
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
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_QosFlowAddOrModifyResponseList := Sequence[*QosFlowAddOrModifyResponseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp_QosFlowAddOrModifyResponseList.Decode(r); err != nil {
			return
		}
		ie.QosFlowAddOrModifyResponseList = &[]QosFlowAddOrModifyResponseItem{}
		for _, i := range tmp_QosFlowAddOrModifyResponseList.Value {
			*ie.QosFlowAddOrModifyResponseList = append(*ie.QosFlowAddOrModifyResponseList, *i)
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_AdditionalDLQosFlowPerTNLInformation := Sequence[*QosFlowPerTNLInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		if err = tmp_AdditionalDLQosFlowPerTNLInformation.Decode(r); err != nil {
			return
		}
		ie.AdditionalDLQosFlowPerTNLInformation = &[]QosFlowPerTNLInformationItem{}
		for _, i := range tmp_AdditionalDLQosFlowPerTNLInformation.Value {
			*ie.AdditionalDLQosFlowPerTNLInformation = append(*ie.AdditionalDLQosFlowPerTNLInformation, *i)
		}
	}
	if aper.IsBitSet(optionals, 5) {
		tmp_QosFlowFailedToAddOrModifyList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp_QosFlowFailedToAddOrModifyList.Decode(r); err != nil {
			return
		}
		ie.QosFlowFailedToAddOrModifyList = &[]QosFlowWithCauseItem{}
		for _, i := range tmp_QosFlowFailedToAddOrModifyList.Value {
			*ie.QosFlowFailedToAddOrModifyList = append(*ie.QosFlowFailedToAddOrModifyList, *i)
		}
	}
	return
}
