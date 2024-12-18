package ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceSetupRequestTransfer struct {
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

func (msg *PDUSessionResourceSetupRequestTransfer) Encode() ([]byte, error) {
	return encodeTransferMessage(msg.toIes())
}

func (msg *PDUSessionResourceSetupRequestTransfer) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.PDUSessionAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionAggregateMaximumBitRate})
	}
	if msg.ULNGUUPTNLInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ULNGUUPTNLInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ULNGUUPTNLInformation})
	}
	if msg.AdditionalULNGUUPTNLInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AdditionalULNGUUPTNLInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AdditionalULNGUUPTNLInformation})
	}
	if msg.DataForwardingNotPossible != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DataForwardingNotPossible},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.DataForwardingNotPossible})
	}
	if msg.PDUSessionType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionType},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionType})
	}
	if msg.SecurityIndication != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SecurityIndication},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SecurityIndication})
	}
	if msg.NetworkInstance != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NetworkInstance},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NetworkInstance})
	}
	if msg.QosFlowSetupRequestList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_QosFlowSetupRequestList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.QosFlowSetupRequestList})
	}
	if msg.CommonNetworkInstance != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CommonNetworkInstance},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CommonNetworkInstance})
	}
	return
}

func (msg *PDUSessionResourceSetupRequestTransfer) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}

func (msg *PDUSessionResourceSetupRequestTransfer) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
	id, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false)
	if err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	c, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
	var buf []byte
	if buf, err = r.ReadOpenType(); err != nil {
		return
	}
	ieR := aper.NewReader(bytes.NewReader(buf))
	switch msgIe.Id.Value {
	case ProtocolIEID_PDUSessionAggregateMaximumBitRate:
		var tmp PDUSessionAggregateMaximumBitRate
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionAggregateMaximumBitRate = &tmp
	case ProtocolIEID_ULNGUUPTNLInformation:
		var tmp UPTransportLayerInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ULNGUUPTNLInformation = &tmp
	case ProtocolIEID_AdditionalULNGUUPTNLInformation:
		var tmp UPTransportLayerInformationList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AdditionalULNGUUPTNLInformation = &tmp
	case ProtocolIEID_DataForwardingNotPossible:
		var tmp DataForwardingNotPossible
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.DataForwardingNotPossible = &tmp
	case ProtocolIEID_PDUSessionType:
		var tmp PDUSessionType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionType = &tmp
	case ProtocolIEID_SecurityIndication:
		var tmp SecurityIndication
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SecurityIndication = &tmp
	case ProtocolIEID_NetworkInstance:
		var tmp NetworkInstance
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NetworkInstance = &tmp
	case ProtocolIEID_QosFlowSetupRequestList:
		var tmp QosFlowSetupRequestList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.QosFlowSetupRequestList = &tmp
	case ProtocolIEID_CommonNetworkInstance:
		var tmp CommonNetworkInstance
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CommonNetworkInstance = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}