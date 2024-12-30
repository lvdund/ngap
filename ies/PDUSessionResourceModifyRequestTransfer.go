package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceModifyRequestTransfer struct {
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate
	ULNGUUPTNLModifyList              *[]ULNGUUPTNLModifyItem
	NetworkInstance                   *int64
	QosFlowAddOrModifyRequestList     *[]QosFlowAddOrModifyRequestItem
	QosFlowToReleaseList              *[]QosFlowWithCauseItem
	AdditionalULNGUUPTNLInformation   *[]UPTransportLayerInformationItem
	CommonNetworkInstance             *[]byte
}

func (msg *PDUSessionResourceModifyRequestTransfer) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceModify, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceModifyRequestTransfer) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.PDUSessionAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionAggregateMaximumBitRate,
		})
	}
	if msg.ULNGUUPTNLModifyList != nil {
		tmp_ULNGUUPTNLModifyList := Sequence[*ULNGUUPTNLModifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivity},
			ext: false,
		}
		for _, i := range *msg.ULNGUUPTNLModifyList {
			tmp_ULNGUUPTNLModifyList.Value = append(tmp_ULNGUUPTNLModifyList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ULNGUUPTNLModifyList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_ULNGUUPTNLModifyList,
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
	if msg.QosFlowAddOrModifyRequestList != nil {
		tmp_QosFlowAddOrModifyRequestList := Sequence[*QosFlowAddOrModifyRequestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		for _, i := range *msg.QosFlowAddOrModifyRequestList {
			tmp_QosFlowAddOrModifyRequestList.Value = append(tmp_QosFlowAddOrModifyRequestList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_QosFlowAddOrModifyRequestList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_QosFlowAddOrModifyRequestList,
		})
	}
	if msg.QosFlowToReleaseList != nil {
		tmp_QosFlowToReleaseList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		for _, i := range *msg.QosFlowToReleaseList {
			tmp_QosFlowToReleaseList.Value = append(tmp_QosFlowToReleaseList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_QosFlowToReleaseList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_QosFlowToReleaseList,
		})
	}
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
	case ProtocolIEID_ULNGUUPTNLModifyList:
		tmp := Sequence[*ULNGUUPTNLModifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivity},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ULNGUUPTNLModifyList = &[]ULNGUUPTNLModifyItem{}
		for _, i := range tmp.Value {
			*msg.ULNGUUPTNLModifyList = append(*msg.ULNGUUPTNLModifyList, *i)
		}
	case ProtocolIEID_NetworkInstance:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.NetworkInstance = int64(tmp.Value)
	case ProtocolIEID_QosFlowAddOrModifyRequestList:
		tmp := Sequence[*QosFlowAddOrModifyRequestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.QosFlowAddOrModifyRequestList = &[]QosFlowAddOrModifyRequestItem{}
		for _, i := range tmp.Value {
			*msg.QosFlowAddOrModifyRequestList = append(*msg.QosFlowAddOrModifyRequestList, *i)
		}
	case ProtocolIEID_QosFlowToReleaseList:
		tmp := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.QosFlowToReleaseList = &[]QosFlowWithCauseItem{}
		for _, i := range tmp.Value {
			*msg.QosFlowToReleaseList = append(*msg.QosFlowToReleaseList, *i)
		}
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
