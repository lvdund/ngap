package ies

import (
	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceSetupRequestTransfer struct {
	// Choice                            uint64
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate
	ULNGUUPTNLInformation             *UPTransportLayerInformation
	AdditionalULNGUUPTNLInformation   *UPTransportLayerInformationList
	DataForwardingNotPossible         *DataForwardingNotPossible
	PDUSessionType                    *PDUSessionType
	SecurityIndication                *SecurityIndication
	NetworkInstance                   *NetworkInstance
	QosFlowSetupRequestList           *QosFlowSetupRequestList
	CommonNetworkInstance             *CommonNetworkInstance
}

func (ie *PDUSessionResourceSetupRequestTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PDUSessionAggregateMaximumBitRate != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.AdditionalULNGUUPTNLInformation != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.DataForwardingNotPossible != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.SecurityIndication != nil {
		aper.SetBit(optionals, 4)
	}
	if ie.NetworkInstance != nil {
		aper.SetBit(optionals, 5)
	}
	if ie.CommonNetworkInstance != nil {
		aper.SetBit(optionals, 6)
	}
	w.WriteBits(optionals, 6)
	if ie.PDUSessionAggregateMaximumBitRate != nil {
		if err = ie.PDUSessionAggregateMaximumBitRate.Encode(w); err != nil {
			return
		}
	}
	if ie.ULNGUUPTNLInformation != nil {
		if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.AdditionalULNGUUPTNLInformation != nil {
		if err = ie.AdditionalULNGUUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.DataForwardingNotPossible != nil {
		if err = ie.DataForwardingNotPossible.Encode(w); err != nil {
			return
		}
	}
	if ie.PDUSessionType != nil {
		if err = ie.PDUSessionType.Encode(w); err != nil {
			return
		}
	}
	if ie.SecurityIndication != nil {
		if err = ie.SecurityIndication.Encode(w); err != nil {
			return
		}
	}
	if ie.NetworkInstance != nil {
		if err = ie.NetworkInstance.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowSetupRequestList != nil {
		if err = ie.QosFlowSetupRequestList.Encode(w); err != nil {
			return
		}
	}
	if ie.CommonNetworkInstance != nil {
		if err = ie.CommonNetworkInstance.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceSetupRequestTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(6); err != nil {
		return
	}
	ie.PDUSessionAggregateMaximumBitRate = new(PDUSessionAggregateMaximumBitRate)
	ie.ULNGUUPTNLInformation = new(UPTransportLayerInformation)
	ie.AdditionalULNGUUPTNLInformation = new(UPTransportLayerInformationList)
	ie.DataForwardingNotPossible = new(DataForwardingNotPossible)
	ie.PDUSessionType = new(PDUSessionType)
	ie.CommonNetworkInstance = new(CommonNetworkInstance)
	ie.NetworkInstance = new(NetworkInstance)
	ie.QosFlowSetupRequestList = new(QosFlowSetupRequestList)
	ie.CommonNetworkInstance = new(CommonNetworkInstance)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.PDUSessionAggregateMaximumBitRate.Decode(r); err != nil {
			return
		}
	}
	if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.AdditionalULNGUUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.DataForwardingNotPossible.Decode(r); err != nil {
			return
		}
	}
	if err = ie.PDUSessionType.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.SecurityIndication.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 5) {
		if err = ie.NetworkInstance.Decode(r); err != nil {
			return
		}
	}
	if err = ie.QosFlowSetupRequestList.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 6) {
		if err = ie.CommonNetworkInstance.Decode(r); err != nil {
			return
		}
	}
	return
}
