package ies

import "github.com/lvdund/ngap/aper"

type HandoverRequestAcknowledgeTransfer struct {
	DLNGUUPTNLInformation         *UPTransportLayerInformation   `False,`
	DLForwardingUPTNLInformation  *UPTransportLayerInformation   `False,OPTIONAL`
	SecurityResult                *SecurityResult                `True,OPTIONAL`
	QosFlowSetupResponseList      *QosFlowListWithDataForwarding `False,`
	QosFlowFailedToSetupList      *QosFlowListWithCause          `False,OPTIONAL`
	DataForwardingResponseDRBList *DataForwardingResponseDRBList `False,OPTIONAL`
	// IEExtensions HandoverRequestAcknowledgeTransferExtIEs `False,OPTIONAL`
}

func (ie *HandoverRequestAcknowledgeTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
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
	if ie.DLNGUUPTNLInformation != nil {
		if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
			return
		}
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
		if err = ie.QosFlowSetupResponseList.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowFailedToSetupList != nil {
		if err = ie.QosFlowFailedToSetupList.Encode(w); err != nil {
			return
		}
	}
	if ie.DataForwardingResponseDRBList != nil {
		if err = ie.DataForwardingResponseDRBList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *HandoverRequestAcknowledgeTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	ie.DLNGUUPTNLInformation = new(UPTransportLayerInformation)
	ie.DLForwardingUPTNLInformation = new(UPTransportLayerInformation)
	ie.SecurityResult = new(SecurityResult)
	ie.QosFlowSetupResponseList = new(QosFlowListWithDataForwarding)
	ie.QosFlowFailedToSetupList = new(QosFlowListWithCause)
	ie.DataForwardingResponseDRBList = new(DataForwardingResponseDRBList)
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.DLForwardingUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.SecurityResult.Decode(r); err != nil {
			return
		}
	}
	if err = ie.QosFlowSetupResponseList.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.QosFlowFailedToSetupList.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 5) {
		if err = ie.DataForwardingResponseDRBList.Decode(r); err != nil {
			return
		}
	}
	return
}
