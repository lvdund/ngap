package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceSetupRequestTransfer struct {
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate `optional`
	ULNGUUPTNLInformation             UPTransportLayerInformation
	AdditionalULNGUUPTNLInformation   []UPTransportLayerInformationItem `optional`
	DataForwardingNotPossible         *DataForwardingNotPossible        `optional`
	PDUSessionType                    PDUSessionType
	SecurityIndication                *SecurityIndication `optional`
	NetworkInstance                   *int64              `optional`
	QosFlowSetupRequestList           []QosFlowSetupRequestItem
	CommonNetworkInstance             []byte `optional`
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
		for _, i := range msg.AdditionalULNGUUPTNLInformation {
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
				Value: msg.CommonNetworkInstance,
			}})
	}
	return
}
func (msg *PDUSessionResourceSetupRequestTransfer) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PDUSessionResourceSetupRequestTransferDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_ULNGUUPTNLInformation]; !ok {
		err = fmt.Errorf("Mandatory field ULNGUUPTNLInformation is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_ULNGUUPTNLInformation},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_PDUSessionType]; !ok {
		err = fmt.Errorf("Mandatory field PDUSessionType is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_PDUSessionType},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_QosFlowSetupRequestList]; !ok {
		err = fmt.Errorf("Mandatory field QosFlowSetupRequestList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_QosFlowSetupRequestList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PDUSessionResourceSetupRequestTransferDecoder struct {
	msg      *PDUSessionResourceSetupRequestTransfer
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *PDUSessionResourceSetupRequestTransferDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_ULNGUUPTNLInformation:
		var tmp UPTransportLayerInformation
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ULNGUUPTNLInformation", err)
			return
		}
		msg.ULNGUUPTNLInformation = tmp
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
	case ProtocolIEID_DataForwardingNotPossible:
		var tmp DataForwardingNotPossible
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DataForwardingNotPossible", err)
			return
		}
		msg.DataForwardingNotPossible = &tmp
	case ProtocolIEID_PDUSessionType:
		var tmp PDUSessionType
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read PDUSessionType", err)
			return
		}
		msg.PDUSessionType = tmp
	case ProtocolIEID_SecurityIndication:
		var tmp SecurityIndication
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SecurityIndication", err)
			return
		}
		msg.SecurityIndication = &tmp
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
	case ProtocolIEID_QosFlowSetupRequestList:
		tmp := Sequence[*QosFlowSetupRequestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowSetupRequestItem { return new(QosFlowSetupRequestItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read QosFlowSetupRequestList", err)
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
