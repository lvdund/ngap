package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyResponseTransfer struct {
	DLNGUUPTNLInformation                *UPTransportLayerInformation    `False,OPTIONAL`
	ULNGUUPTNLInformation                *UPTransportLayerInformation    `False,OPTIONAL`
	QosFlowAddOrModifyResponseList       *QosFlowAddOrModifyResponseList `False,OPTIONAL`
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList   `False,OPTIONAL`
	QosFlowFailedToAddOrModifyList       *QosFlowListWithCause           `False,OPTIONAL`
	// IEExtensions PDUSessionResourceModifyResponseTransferExtIEs `False,OPTIONAL`
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
		if err = ie.QosFlowAddOrModifyResponseList.Encode(w); err != nil {
			return
		}
	}
	if ie.AdditionalDLQosFlowPerTNLInformation != nil {
		if err = ie.AdditionalDLQosFlowPerTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowFailedToAddOrModifyList != nil {
		if err = ie.QosFlowFailedToAddOrModifyList.Encode(w); err != nil {
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
	ie.DLNGUUPTNLInformation = new(UPTransportLayerInformation)
	ie.ULNGUUPTNLInformation = new(UPTransportLayerInformation)
	ie.QosFlowAddOrModifyResponseList = new(QosFlowAddOrModifyResponseList)
	ie.AdditionalDLQosFlowPerTNLInformation = new(QosFlowPerTNLInformationList)
	ie.QosFlowFailedToAddOrModifyList = new(QosFlowListWithCause)
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
		if err = ie.QosFlowAddOrModifyResponseList.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.AdditionalDLQosFlowPerTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 5) {
		if err = ie.QosFlowFailedToAddOrModifyList.Decode(r); err != nil {
			return
		}
	}
	return
}
