package ies

import (
	"bytes"
	"fmt"
	"io"

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

func (msg *PDUSessionResourceSetupRequestTransfer) Encode(w io.Writer) (err error) {
	return encodeTransferMessage(w, msg.toIes())
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
	if msg.PDUSessionType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionType},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionType})
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
	case ProtocolIEID_PDUSessionType:
		var tmp PDUSessionType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionType = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}