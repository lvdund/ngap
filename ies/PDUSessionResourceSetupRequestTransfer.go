package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceSetupRequestTransfer struct {
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate
	ULNGUUPTNLInformation             UPTransportLayerInformation
	AdditionalULNGUUPTNLInformation   *[]UPTransportLayerInformationItem
	DataForwardingNotPossible         *DataForwardingNotPossible
	PDUSessionType                    PDUSessionType
	SecurityIndication                *SecurityIndication
	NetworkInstance                   *int64
	QosFlowSetupRequestList           []QosFlowSetupRequestItem
	CommonNetworkInstance             *[]byte
}

func (msg *PDUSessionResourceSetupRequestTransfer) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceSetupRequestTransfer) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.PDUSessionAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionAggregateMaximumBitRate,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_ULNGUUPTNLInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.ULNGUUPTNLInformation,
	})
	if msg.AdditionalULNGUUPTNLInformation != nil {
		tmp_AdditionalULNGUUPTNLInformation := Sequence[*UPTransportLayerInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		for _, i := range *msg.AdditionalULNGUUPTNLInformation {
			tmp_AdditionalULNGUUPTNLInformation.Value = append(tmp_AdditionalULNGUUPTNLInformation.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AdditionalULNGUUPTNLInformation},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_AdditionalULNGUUPTNLInformation,
		})
	}
	if msg.DataForwardingNotPossible != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DataForwardingNotPossible},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.DataForwardingNotPossible,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionType},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.PDUSessionType,
	})
	if msg.SecurityIndication != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SecurityIndication},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SecurityIndication,
		})
	}
	if msg.NetworkInstance != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NetworkInstance},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 256},
				ext:   true,
				Value: aper.Integer(*msg.NetworkInstance),
			}})
	}
	tmp_QosFlowSetupRequestList := Sequence[*QosFlowSetupRequestItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	for _, i := range msg.QosFlowSetupRequestList {
		tmp_QosFlowSetupRequestList.Value = append(tmp_QosFlowSetupRequestList.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_QosFlowSetupRequestList},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_QosFlowSetupRequestList,
	})
	if msg.CommonNetworkInstance != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CommonNetworkInstance},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: *msg.CommonNetworkInstance,
			}})
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
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: int64(aper.POW_16) - 1}, false); err != nil {
		return
	}
	msgIe = new(NgapMessageIE)
	msgIe.Id.Value = aper.Integer(id)
	if c, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return
	}
	msgIe.Criticality.Value = aper.Enumerated(c)
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
		msg.ULNGUUPTNLInformation = tmp
	case ProtocolIEID_AdditionalULNGUUPTNLInformation:
		tmp := Sequence[*UPTransportLayerInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AdditionalULNGUUPTNLInformation = &[]UPTransportLayerInformationItem{}
		for _, i := range tmp.Value {
			*msg.AdditionalULNGUUPTNLInformation = append(*msg.AdditionalULNGUUPTNLInformation, *i)
		}
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
		msg.PDUSessionType = tmp
	case ProtocolIEID_SecurityIndication:
		var tmp SecurityIndication
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SecurityIndication = &tmp
	case ProtocolIEID_NetworkInstance:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.NetworkInstance = int64(tmp.Value)
	case ProtocolIEID_QosFlowSetupRequestList:
		tmp := Sequence[*QosFlowSetupRequestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.QosFlowSetupRequestList = []QosFlowSetupRequestItem{}
		for _, i := range tmp.Value {
			msg.QosFlowSetupRequestList = append(msg.QosFlowSetupRequestList, *i)
		}
	case ProtocolIEID_CommonNetworkInstance:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.CommonNetworkInstance = tmp.Value
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
