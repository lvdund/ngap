package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyIndicationTransfer struct {
	DLQosFlowPerTNLInformation           *QosFlowPerTNLInformation     `True,`
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList `False,OPTIONAL`
	// IEExtensions PDUSessionResourceModifyIndicationTransferExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceModifyIndicationTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AdditionalDLQosFlowPerTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
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
	return
}
func (ie *PDUSessionResourceModifyIndicationTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.DLQosFlowPerTNLInformation = new(QosFlowPerTNLInformation)
	ie.AdditionalDLQosFlowPerTNLInformation = new(QosFlowPerTNLInformationList)
	if err = ie.DLQosFlowPerTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.AdditionalDLQosFlowPerTNLInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
