package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList      []QosFlowModifyConfirmItem
	ULNGUUPTNLInformation         UPTransportLayerInformation
	AdditionalNGUUPTNLInformation *[]UPTransportLayerInformationPairItem
	QosFlowFailedToModifyList     *[]QosFlowWithCauseItem
	// IEExtensions  *PDUSessionResourceModifyConfirmTransferExtIEs
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
			return
		}
	}
	if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
		return
	}
	if ie.AdditionalNGUUPTNLInformation != nil {
		if len(*ie.AdditionalNGUUPTNLInformation) > 0 {
			tmp := Sequence[*UPTransportLayerInformationPairItem]{
				Value: []*UPTransportLayerInformationPairItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
				ext:   false,
			}
			for _, i := range *ie.AdditionalNGUUPTNLInformation {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.QosFlowFailedToModifyList != nil {
		if len(*ie.QosFlowFailedToModifyList) > 0 {
			tmp := Sequence[*QosFlowWithCauseItem]{
				Value: []*QosFlowWithCauseItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range *ie.QosFlowFailedToModifyList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
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
	if err = tmp_QosFlowModifyConfirmList.Decode(r); err != nil {
		return
	}
	ie.QosFlowModifyConfirmList = []QosFlowModifyConfirmItem{}
	for _, i := range tmp_QosFlowModifyConfirmList.Value {
		ie.QosFlowModifyConfirmList = append(ie.QosFlowModifyConfirmList, *i)
	}
	if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_AdditionalNGUUPTNLInformation := Sequence[*UPTransportLayerInformationPairItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		if err = tmp_AdditionalNGUUPTNLInformation.Decode(r); err != nil {
			return
		}
		ie.AdditionalNGUUPTNLInformation = &[]UPTransportLayerInformationPairItem{}
		for _, i := range tmp_AdditionalNGUUPTNLInformation.Value {
			*ie.AdditionalNGUUPTNLInformation = append(*ie.AdditionalNGUUPTNLInformation, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_QosFlowFailedToModifyList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp_QosFlowFailedToModifyList.Decode(r); err != nil {
			return
		}
		ie.QosFlowFailedToModifyList = &[]QosFlowWithCauseItem{}
		for _, i := range tmp_QosFlowFailedToModifyList.Value {
			*ie.QosFlowFailedToModifyList = append(*ie.QosFlowFailedToModifyList, *i)
		}
	}
	return
}
