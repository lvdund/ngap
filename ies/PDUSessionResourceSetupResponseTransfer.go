package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupResponseTransfer struct {
	DLQosFlowPerTNLInformation           *QosFlowPerTNLInformation     `True,`
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList `False,OPTIONAL`
	SecurityResult                       *SecurityResult               `True,OPTIONAL`
	QosFlowFailedToSetupList             *QosFlowListWithCause         `False,OPTIONAL`
	// IEExtensions PDUSessionResourceSetupResponseTransferExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceSetupResponseTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
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
	if ie.DLQosFlowPerTNLInformation != nil {
		if err = ie.DLQosFlowPerTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.AdditionalDLQosFlowPerTNLInformation != nil {
		if err = ie.AdditionalDLQosFlowPerTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.SecurityResult != nil {
		if err = ie.SecurityResult.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowFailedToSetupList != nil {
		if err = ie.QosFlowFailedToSetupList.Encode(w); err != nil {
			return
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
	ie.DLQosFlowPerTNLInformation = new(QosFlowPerTNLInformation)
	ie.AdditionalDLQosFlowPerTNLInformation = new(QosFlowPerTNLInformationList)
	ie.SecurityResult = new(SecurityResult)
	ie.QosFlowFailedToSetupList = new(QosFlowListWithCause)
	if err = ie.DLQosFlowPerTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.AdditionalDLQosFlowPerTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.SecurityResult.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.QosFlowFailedToSetupList.Decode(r); err != nil {
			return
		}
	}
	return
}
