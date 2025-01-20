package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type HandoverRequestAcknowledgeTransfer struct {
	DLNGUUPTNLInformation         UPTransportLayerInformation     `madatory`
	DLForwardingUPTNLInformation  *UPTransportLayerInformation    `optional`
	SecurityResult                *SecurityResult                 `optional`
	QosFlowSetupResponseList      []QosFlowItemWithDataForwarding `lb:1,ub:maxnoofQosFlows,madatory`
	QosFlowFailedToSetupList      []QosFlowWithCauseItem          `lb:1,ub:maxnoofQosFlows,optional`
	DataForwardingResponseDRBList []DataForwardingResponseDRBItem `lb:1,ub:maxnoofDRBs,optional`
	// IEExtensions *HandoverRequestAcknowledgeTransferExtIEs `optional`
}

func (ie *HandoverRequestAcknowledgeTransfer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
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
	if ie.QosFlowFailedToSetupList != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.DataForwardingResponseDRBList != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode DLNGUUPTNLInformation", err)
		return
	}
	if ie.DLForwardingUPTNLInformation != nil {
		if err = ie.DLForwardingUPTNLInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode DLForwardingUPTNLInformation", err)
			return
		}
	}
	if ie.SecurityResult != nil {
		if err = ie.SecurityResult.Encode(w); err != nil {
			err = utils.WrapError("Encode SecurityResult", err)
			return
		}
	}
	if len(ie.QosFlowSetupResponseList) > 0 {
		tmp := Sequence[*QosFlowItemWithDataForwarding]{
			Value: []*QosFlowItemWithDataForwarding{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowSetupResponseList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode QosFlowSetupResponseList", err)
			return
		}
	} else {
		err = utils.WrapError("QosFlowSetupResponseList is nil", err)
		return
	}
	if len(ie.QosFlowFailedToSetupList) > 0 {
		tmp := Sequence[*QosFlowWithCauseItem]{
			Value: []*QosFlowWithCauseItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowFailedToSetupList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode QosFlowFailedToSetupList", err)
			return
		}
	}
	if len(ie.DataForwardingResponseDRBList) > 0 {
		tmp := Sequence[*DataForwardingResponseDRBItem]{
			Value: []*DataForwardingResponseDRBItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext:   false,
		}
		for _, i := range ie.DataForwardingResponseDRBList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode DataForwardingResponseDRBList", err)
			return
		}
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *HandoverRequestAcknowledgeTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		err = utils.WrapError("Read DLNGUUPTNLInformation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DLForwardingUPTNLInformation.Decode(r); err != nil {
			err = utils.WrapError("Read DLForwardingUPTNLInformation", err)
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.SecurityResult.Decode(r); err != nil {
			err = utils.WrapError("Read SecurityResult", err)
			return
		}
	}
	tmp_QosFlowSetupResponseList := Sequence[*QosFlowItemWithDataForwarding]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	fn := func() *QosFlowItemWithDataForwarding { return new(QosFlowItemWithDataForwarding) }
	if err = tmp_QosFlowSetupResponseList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read QosFlowSetupResponseList", err)
		return
	}
	ie.QosFlowSetupResponseList = []QosFlowItemWithDataForwarding{}
	for _, i := range tmp_QosFlowSetupResponseList.Value {
		ie.QosFlowSetupResponseList = append(ie.QosFlowSetupResponseList, *i)
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_QosFlowFailedToSetupList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowWithCauseItem { return new(QosFlowWithCauseItem) }
		if err = tmp_QosFlowFailedToSetupList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read QosFlowFailedToSetupList", err)
			return
		}
		ie.QosFlowFailedToSetupList = []QosFlowWithCauseItem{}
		for _, i := range tmp_QosFlowFailedToSetupList.Value {
			ie.QosFlowFailedToSetupList = append(ie.QosFlowFailedToSetupList, *i)
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_DataForwardingResponseDRBList := Sequence[*DataForwardingResponseDRBItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DataForwardingResponseDRBItem { return new(DataForwardingResponseDRBItem) }
		if err = tmp_DataForwardingResponseDRBList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read DataForwardingResponseDRBList", err)
			return
		}
		ie.DataForwardingResponseDRBList = []DataForwardingResponseDRBItem{}
		for _, i := range tmp_DataForwardingResponseDRBList.Value {
			ie.DataForwardingResponseDRBList = append(ie.DataForwardingResponseDRBList, *i)
		}
	}
	return
}
