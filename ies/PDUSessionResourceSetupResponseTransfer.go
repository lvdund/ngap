package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceSetupResponseTransfer struct {
	DLQosFlowPerTNLInformation           QosFlowPerTNLInformation
	AdditionalDLQosFlowPerTNLInformation []QosFlowPerTNLInformationItem `optional`
	SecurityResult                       *SecurityResult                `optional`
	QosFlowFailedToSetupList             []QosFlowWithCauseItem         `optional`
	// IEExtensions *PDUSessionResourceSetupResponseTransferExtIEs `optional`
}

func (ie *PDUSessionResourceSetupResponseTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AdditionalDLQosFlowPerTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SecurityResult != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.QosFlowFailedToSetupList != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if err = ie.DLQosFlowPerTNLInformation.Encode(w); err != nil {
		err = utils.WrapError("Read DLQosFlowPerTNLInformation", err)
		return
	}
	if ie.AdditionalDLQosFlowPerTNLInformation != nil {
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
				err = utils.WrapError("Read AdditionalDLQosFlowPerTNLInformation", err)
				return
			}
		}
	}
	if ie.SecurityResult != nil {
		if err = ie.SecurityResult.Encode(w); err != nil {
			err = utils.WrapError("Read SecurityResult", err)
			return
		}
	}
	if ie.QosFlowFailedToSetupList != nil {
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
				err = utils.WrapError("Read QosFlowFailedToSetupList", err)
				return
			}
		}
	}
	return
}
func (ie *PDUSessionResourceSetupResponseTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if err = ie.DLQosFlowPerTNLInformation.Decode(r); err != nil {
		err = utils.WrapError("Read DLQosFlowPerTNLInformation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
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
	if aper.IsBitSet(optionals, 2) {
		if err = ie.SecurityResult.Decode(r); err != nil {
			err = utils.WrapError("Read SecurityResult", err)
			return
		}
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
	return
}
