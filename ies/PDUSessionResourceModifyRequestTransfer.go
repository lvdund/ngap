package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceModifyRequestTransfer struct {
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate `optional,mandatory,reject`
	ULNGUUPTNLModifyList              []ULNGUUPTNLModifyItem             `lb:1,ub:maxnoofMultiConnectivity,optional,mandatory,reject`
	NetworkInstance                   *int64                             `lb:1,ub:256,optional,mandatory,reject,valueExt`
	QosFlowAddOrModifyRequestList     []QosFlowAddOrModifyRequestItem    `lb:1,ub:maxnoofQosFlows,optional,mandatory,reject`
	QosFlowToReleaseList              []QosFlowWithCauseItem             `lb:1,ub:maxnoofQosFlows,optional,mandatory,reject`
	AdditionalULNGUUPTNLInformation   []UPTransportLayerInformationItem  `lb:1,ub:maxnoofMultiConnectivityMinusOne,optional,mandatory,reject`
	CommonNetworkInstance             []byte                             `lb:0,ub:0,optional,mandatory,ignore`
}

func (msg *PDUSessionResourceModifyRequestTransfer) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceModify, Criticality_PresentReject, ies)
}
func (msg *PDUSessionResourceModifyRequestTransfer) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	if msg.PDUSessionAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionAggregateMaximumBitRate,
		})
	}
	if len(msg.ULNGUUPTNLModifyList) > 0 {
		tmp_ULNGUUPTNLModifyList := Sequence[*ULNGUUPTNLModifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivity},
			ext: false,
		}
		for _, i := range msg.ULNGUUPTNLModifyList {
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
	if len(msg.QosFlowAddOrModifyRequestList) > 0 {
		tmp_QosFlowAddOrModifyRequestList := Sequence[*QosFlowAddOrModifyRequestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		for _, i := range msg.QosFlowAddOrModifyRequestList {
			tmp_QosFlowAddOrModifyRequestList.Value = append(tmp_QosFlowAddOrModifyRequestList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_QosFlowAddOrModifyRequestList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_QosFlowAddOrModifyRequestList,
		})
	}
	if len(msg.QosFlowToReleaseList) > 0 {
		tmp_QosFlowToReleaseList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		for _, i := range msg.QosFlowToReleaseList {
			tmp_QosFlowToReleaseList.Value = append(tmp_QosFlowToReleaseList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_QosFlowToReleaseList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_QosFlowToReleaseList,
		})
	}
	if len(msg.AdditionalULNGUUPTNLInformation) > 0 {
		tmp_AdditionalULNGUUPTNLInformation := Sequence[*UPTransportLayerInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		for _, i := range msg.AdditionalULNGUUPTNLInformation {
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
				Value: msg.CommonNetworkInstance,
			}})
	}
	return
}
func (msg *PDUSessionResourceModifyRequestTransfer) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("PDUSessionResourceModifyRequestTransfer"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PDUSessionResourceModifyRequestTransferDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	return
}

type PDUSessionResourceModifyRequestTransferDecoder struct {
	msg      *PDUSessionResourceModifyRequestTransfer
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *PDUSessionResourceModifyRequestTransferDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	ieId := msgIe.Id.Value
	if _, ok := decoder.list[ieId]; ok {
		err = fmt.Errorf("Duplicated protocol IEID[%d] found", ieId)
		return
	}
	decoder.list[ieId] = msgIe
	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg
	switch msgIe.Id.Value {
	case ProtocolIEID_PDUSessionAggregateMaximumBitRate:
		var tmp PDUSessionAggregateMaximumBitRate
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PDUSessionAggregateMaximumBitRate", err)
			return
		}
		msg.PDUSessionAggregateMaximumBitRate = &tmp
	case ProtocolIEID_ULNGUUPTNLModifyList:
		tmp := Sequence[*ULNGUUPTNLModifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivity},
			ext: false,
		}
		fn := func() *ULNGUUPTNLModifyItem { return new(ULNGUUPTNLModifyItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ULNGUUPTNLModifyList", err)
			return
		}
		msg.ULNGUUPTNLModifyList = []ULNGUUPTNLModifyItem{}
		for _, i := range tmp.Value {
			msg.ULNGUUPTNLModifyList = append(msg.ULNGUUPTNLModifyList, *i)
		}
	case ProtocolIEID_NetworkInstance:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NetworkInstance", err)
			return
		}
		*msg.NetworkInstance = int64(tmp.Value)
	case ProtocolIEID_QosFlowAddOrModifyRequestList:
		tmp := Sequence[*QosFlowAddOrModifyRequestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowAddOrModifyRequestItem { return new(QosFlowAddOrModifyRequestItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read QosFlowAddOrModifyRequestList", err)
			return
		}
		msg.QosFlowAddOrModifyRequestList = []QosFlowAddOrModifyRequestItem{}
		for _, i := range tmp.Value {
			msg.QosFlowAddOrModifyRequestList = append(msg.QosFlowAddOrModifyRequestList, *i)
		}
	case ProtocolIEID_QosFlowToReleaseList:
		tmp := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowWithCauseItem { return new(QosFlowWithCauseItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read QosFlowToReleaseList", err)
			return
		}
		msg.QosFlowToReleaseList = []QosFlowWithCauseItem{}
		for _, i := range tmp.Value {
			msg.QosFlowToReleaseList = append(msg.QosFlowToReleaseList, *i)
		}
	case ProtocolIEID_AdditionalULNGUUPTNLInformation:
		tmp := Sequence[*UPTransportLayerInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofMultiConnectivityMinusOne},
			ext: false,
		}
		fn := func() *UPTransportLayerInformationItem { return new(UPTransportLayerInformationItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AdditionalULNGUUPTNLInformation", err)
			return
		}
		msg.AdditionalULNGUUPTNLInformation = []UPTransportLayerInformationItem{}
		for _, i := range tmp.Value {
			msg.AdditionalULNGUUPTNLInformation = append(msg.AdditionalULNGUUPTNLInformation, *i)
		}
	case ProtocolIEID_CommonNetworkInstance:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CommonNetworkInstance", err)
			return
		}
		msg.CommonNetworkInstance = tmp.Value
	default:
		switch msgIe.Criticality.Value {
		case Criticality_PresentReject:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: reject)", msgIe.Id.Value)
		case Criticality_PresentIgnore:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: ignore)", msgIe.Id.Value)
		case Criticality_PresentNotify:
			fmt.Errorf("Not comprehended IE ID 0x%04x (criticality: notify)", msgIe.Id.Value)
		}
		if msgIe.Criticality.Value != Criticality_PresentIgnore {
			decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
				IECriticality: msgIe.Criticality,
				IEID:          msgIe.Id,
				TypeOfError:   TypeOfError{Value: TypeOfErrorNotunderstood},
			})
		}
	}
	return
}
