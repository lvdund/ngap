package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList      *QosFlowModifyConfirmList            `False,`
	ULNGUUPTNLInformation         *UPTransportLayerInformation         `False,`
	AdditionalNGUUPTNLInformation *UPTransportLayerInformationPairList `False,OPTIONAL`
	QosFlowFailedToModifyList     *QosFlowListWithCause                `False,OPTIONAL`
	// IEExtensions PDUSessionResourceModifyConfirmTransferExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceModifyConfirmTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
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
	if ie.QosFlowModifyConfirmList != nil {
		if err = ie.QosFlowModifyConfirmList.Encode(w); err != nil {
			return
		}
	}
	if ie.ULNGUUPTNLInformation != nil {
		if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.AdditionalNGUUPTNLInformation != nil {
		if err = ie.AdditionalNGUUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowFailedToModifyList != nil {
		if err = ie.QosFlowFailedToModifyList.Encode(w); err != nil {
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
	ie.QosFlowModifyConfirmList = new(QosFlowModifyConfirmList)
	ie.ULNGUUPTNLInformation = new(UPTransportLayerInformation)
	ie.AdditionalNGUUPTNLInformation = new(UPTransportLayerInformationPairList)
	ie.QosFlowFailedToModifyList = new(QosFlowListWithCause)
	if err = ie.QosFlowModifyConfirmList.Decode(r); err != nil {
		return
	}
	if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.AdditionalNGUUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.QosFlowFailedToModifyList.Decode(r); err != nil {
			return
		}
	}
	return
}
