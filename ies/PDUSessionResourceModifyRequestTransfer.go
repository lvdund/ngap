package ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceModifyRequestTransfer struct {
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate
	ULNGUUPTNLModifyList              *ULNGUUPTNLModifyList
	NetworkInstance                   *NetworkInstance
	QosFlowAddOrModifyRequestList     *QosFlowAddOrModifyRequestList
	QosFlowToReleaseList              *QosFlowListWithCause
	AdditionalULNGUUPTNLInformation   *UPTransportLayerInformationList
	CommonNetworkInstance             *CommonNetworkInstance
}

func (msg *PDUSessionResourceModifyRequestTransfer) Encode() ([]byte, error) {
	return encodeTransferMessage(msg.toIes())
}

func (msg *PDUSessionResourceModifyRequestTransfer) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.PDUSessionAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionAggregateMaximumBitRate})
	}
	if msg.ULNGUUPTNLModifyList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ULNGUUPTNLModifyList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ULNGUUPTNLModifyList})
	}
	if msg.NetworkInstance != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NetworkInstance},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NetworkInstance})
	}
	if msg.QosFlowAddOrModifyRequestList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_QosFlowAddOrModifyRequestList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.QosFlowAddOrModifyRequestList})
	}
	if msg.QosFlowToReleaseList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_QosFlowToReleaseList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.QosFlowToReleaseList})
	}
	if msg.AdditionalULNGUUPTNLInformation != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AdditionalULNGUUPTNLInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AdditionalULNGUUPTNLInformation})
	}
	if msg.CommonNetworkInstance != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CommonNetworkInstance},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CommonNetworkInstance})
	}
	return
}

func (msg *PDUSessionResourceModifyRequestTransfer) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}

func (msg *PDUSessionResourceModifyRequestTransfer) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_ULNGUUPTNLModifyList:
		var tmp ULNGUUPTNLModifyList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ULNGUUPTNLModifyList = &tmp
	case ProtocolIEID_NetworkInstance:
		var tmp NetworkInstance
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NetworkInstance = &tmp
	case ProtocolIEID_QosFlowAddOrModifyRequestList:
		var tmp QosFlowAddOrModifyRequestList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.QosFlowAddOrModifyRequestList = &tmp
	case ProtocolIEID_QosFlowToReleaseList:
		var tmp QosFlowListWithCause
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.QosFlowToReleaseList = &tmp
	case ProtocolIEID_AdditionalULNGUUPTNLInformation:
		var tmp UPTransportLayerInformationList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AdditionalULNGUUPTNLInformation = &tmp
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
