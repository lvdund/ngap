package ies

import "github.com/lvdund/ngap/aper"

const (
	PDUSessionResourceSetupRequestTransferIEsPresentNothing uint64 = iota /* No components present */
	PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionAggregateMaximumBitRate
	PDUSessionResourceSetupRequestTransferIEsPresentULNGUUPTNLInformation
	PDUSessionResourceSetupRequestTransferIEsPresentAdditionalULNGUUPTNLInformation
	PDUSessionResourceSetupRequestTransferIEsPresentDataForwardingNotPossible
	PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionType
	PDUSessionResourceSetupRequestTransferIEsPresentSecurityIndication
	PDUSessionResourceSetupRequestTransferIEsPresentNetworkInstance
	PDUSessionResourceSetupRequestTransferIEsPresentQosFlowSetupRequestList
)

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
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return
	}
	switch ie.Choice {
	case PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionAggregateMaximumBitRate:
		err = ie.PDUSessionAggregateMaximumBitRate.Encode(w)
	case PDUSessionResourceSetupRequestTransferIEsPresentULNGUUPTNLInformation:
		err = ie.ULNGUUPTNLInformation.Encode(w)
	case PDUSessionResourceSetupRequestTransferIEsPresentAdditionalULNGUUPTNLInformation:
		err = ie.AdditionalULNGUUPTNLInformation.Encode(w)
	case PDUSessionResourceSetupRequestTransferIEsPresentDataForwardingNotPossible:
		err = ie.DataForwardingNotPossible.Encode(w)
	case PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionType:
		err = ie.PDUSessionType.Encode(w)
	case PDUSessionResourceSetupRequestTransferIEsPresentSecurityIndication:
		err = ie.SecurityIndication.Encode(w)
	case PDUSessionResourceSetupRequestTransferIEsPresentNetworkInstance:
		err = ie.NetworkInstance.Encode(w)
	case PDUSessionResourceSetupRequestTransferIEsPresentQosFlowSetupRequestList:
		err = ie.QosFlowSetupRequestList.Encode(w)
	}
	return
}
func (ie *PDUSessionResourceSetupRequestTransfer) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return
	}
	switch ie.Choice {
	case PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionAggregateMaximumBitRate:
		var tmp PDUSessionAggregateMaximumBitRate
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PDUSessionAggregateMaximumBitRate = &tmp
	case PDUSessionResourceSetupRequestTransferIEsPresentULNGUUPTNLInformation:
		var tmp UPTransportLayerInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.ULNGUUPTNLInformation = &tmp
	case PDUSessionResourceSetupRequestTransferIEsPresentAdditionalULNGUUPTNLInformation:
		var tmp UPTransportLayerInformationList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.AdditionalULNGUUPTNLInformation = &tmp
	case PDUSessionResourceSetupRequestTransferIEsPresentDataForwardingNotPossible:
		var tmp DataForwardingNotPossible
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.DataForwardingNotPossible = &tmp
	case PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionType:
		var tmp PDUSessionType
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PDUSessionType = &tmp
	case PDUSessionResourceSetupRequestTransferIEsPresentSecurityIndication:
		var tmp SecurityIndication
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SecurityIndication = &tmp
	case PDUSessionResourceSetupRequestTransferIEsPresentNetworkInstance:
		var tmp NetworkInstance
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NetworkInstance = &tmp
	case PDUSessionResourceSetupRequestTransferIEsPresentQosFlowSetupRequestList:
		var tmp QosFlowSetupRequestList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.QosFlowSetupRequestList = &tmp
	}
	return
}
