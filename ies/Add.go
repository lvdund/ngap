package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupRequestTransfer struct {
	Choice                            uint64
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate
	ULNGUUPTNLInformation             *UPTransportLayerInformation
	AdditionalULNGUUPTNLInformation   *UPTransportLayerInformationList
	DataForwardingNotPossible         *DataForwardingNotPossible
	PDUSessionType                    *PDUSessionType
	SecurityIndication                *SecurityIndication
	NetworkInstance                   *NetworkInstance
	QosFlowSetupRequestList           *QosFlowSetupRequestList
}

func (ie *PDUSessionResourceSetupRequestTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.PDUSessionAggregateMaximumBitRate.Encode(w)
	case 2:
		err = ie.ULNGUUPTNLInformation.Encode(w)
	case 3:
		err = ie.AdditionalULNGUUPTNLInformation.Encode(w)
	case 4:
		err = ie.DataForwardingNotPossible.Encode(w)
	case 5:
		err = ie.PDUSessionType.Encode(w)
	case 6:
		err = ie.SecurityIndication.Encode(w)
	case 7:
		err = ie.NetworkInstance.Encode(w)
	case 8:
		err = ie.QosFlowSetupRequestList.Encode(w)
	}
	return
}
func (ie *PDUSessionResourceSetupRequestTransfer) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp PDUSessionAggregateMaximumBitRate
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PDUSessionAggregateMaximumBitRate = &tmp
	case 2:
		var tmp UPTransportLayerInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ULNGUUPTNLInformation = &tmp
	case 3:
		var tmp UPTransportLayerInformationList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AdditionalULNGUUPTNLInformation = &tmp
	case 4:
		var tmp DataForwardingNotPossible
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DataForwardingNotPossible = &tmp
	case 5:
		var tmp PDUSessionType
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PDUSessionType = &tmp
	case 6:
		var tmp SecurityIndication
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SecurityIndication = &tmp
	case 7:
		var tmp NetworkInstance
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NetworkInstance = &tmp
	case 8:
		var tmp QosFlowSetupRequestList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.QosFlowSetupRequestList = &tmp
	}
	return
}
