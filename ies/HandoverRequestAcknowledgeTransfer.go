package ies

import "github.com/lvdund/ngap/aper"

type HandoverRequestAcknowledgeTransfer struct {
	DLNGUUPTNLInformation         UPTransportLayerInformation
	DLForwardingUPTNLInformation  *UPTransportLayerInformation
	SecurityResult                *SecurityResult
	QosFlowSetupResponseList      *[]QosFlowItemWithDataForwarding
	QosFlowFailedToSetupList      *[]QosFlowWithCauseItem
	DataForwardingResponseDRBList *[]DataForwardingResponseDRBItem
	// IEExtensions  *HandoverRequestAcknowledgeTransferExtIEs
}

func (ie *HandoverRequestAcknowledgeTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLForwardingUPTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SecurityResult != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.QosFlowSetupResponseList != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.QosFlowFailedToSetupList != nil {
		aper.SetBit(optionals, 4)
	}
	if ie.DataForwardingResponseDRBList != nil {
		aper.SetBit(optionals, 5)
	}
	w.WriteBits(optionals, 6)
	if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
		return
	}
	if ie.DLForwardingUPTNLInformation != nil {
		if err = ie.DLForwardingUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.SecurityResult != nil {
		if err = ie.SecurityResult.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowSetupResponseList != nil {
		if len(*ie.QosFlowSetupResponseList) > 0 {
			tmp := Sequence[*QosFlowItemWithDataForwarding]{
				Value: []*QosFlowItemWithDataForwarding{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range *ie.QosFlowSetupResponseList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.QosFlowFailedToSetupList != nil {
		if len(*ie.QosFlowFailedToSetupList) > 0 {
			tmp := Sequence[*QosFlowWithCauseItem]{
				Value: []*QosFlowWithCauseItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range *ie.QosFlowFailedToSetupList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.DataForwardingResponseDRBList != nil {
		if len(*ie.DataForwardingResponseDRBList) > 0 {
			tmp := Sequence[*DataForwardingResponseDRBItem]{
				Value: []*DataForwardingResponseDRBItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
				ext:   false,
			}
			for _, i := range *ie.DataForwardingResponseDRBList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	return
}
func (ie *HandoverRequestAcknowledgeTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(6); err != nil {
		return
	}
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DLForwardingUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.SecurityResult.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_QosFlowSetupResponseList := Sequence[*QosFlowItemWithDataForwarding]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp_QosFlowSetupResponseList.Decode(r); err != nil {
			return
		}
		ie.QosFlowSetupResponseList = &[]QosFlowItemWithDataForwarding{}
		for _, i := range tmp_QosFlowSetupResponseList.Value {
			*ie.QosFlowSetupResponseList = append(*ie.QosFlowSetupResponseList, *i)
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_QosFlowFailedToSetupList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp_QosFlowFailedToSetupList.Decode(r); err != nil {
			return
		}
		ie.QosFlowFailedToSetupList = &[]QosFlowWithCauseItem{}
		for _, i := range tmp_QosFlowFailedToSetupList.Value {
			*ie.QosFlowFailedToSetupList = append(*ie.QosFlowFailedToSetupList, *i)
		}
	}
	if aper.IsBitSet(optionals, 5) {
		tmp_DataForwardingResponseDRBList := Sequence[*DataForwardingResponseDRBItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		if err = tmp_DataForwardingResponseDRBList.Decode(r); err != nil {
			return
		}
		ie.DataForwardingResponseDRBList = &[]DataForwardingResponseDRBItem{}
		for _, i := range tmp_DataForwardingResponseDRBList.Value {
			*ie.DataForwardingResponseDRBList = append(*ie.DataForwardingResponseDRBList, *i)
		}
	}
	return
}
