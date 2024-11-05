package ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

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
	// Choice                            uint64
	PDUSessionAggregateMaximumBitRate *PDUSessionAggregateMaximumBitRate
	ULNGUUPTNLInformation             *UPTransportLayerInformation
	AdditionalULNGUUPTNLInformation   *UPTransportLayerInformationList
	DataForwardingNotPossible         *DataForwardingNotPossible
	PDUSessionType                    *PDUSessionType
	SecurityIndication                *SecurityIndication
	NetworkInstance                   *NetworkInstance
	QosFlowSetupRequestList           *QosFlowSetupRequestList
}

func (msg *PDUSessionResourceSetupRequestTransfer) Encode() (wire []byte, err error) {
	var buf bytes.Buffer
	aw := aper.NewWriter(&buf)
	if err = aw.WriteBool(aper.Zero); err != nil {
		return
	}
	ies := msg.toIes()
	if len(ies) == 0 {
		err = fmt.Errorf("empty message")
		return
	}

	var subbuf bytes.Buffer
	cW := aper.NewWriter(&subbuf) //container writer
	if err = aper.WriteSequenceOf[NgapMessageIE](ies, cW, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, false); err != nil {
		return
	}
	if err = cW.Close(); err != nil { //finalize
		return
	}

	if err = aw.WriteOpenType(subbuf.Bytes()); err != nil {
		return
	}
	err = aw.Close()
	wire = buf.Bytes()
	return
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
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.QosFlowSetupRequestList})
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
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}

// func (ie *PDUSessionResourceSetupRequestTransfer) Encode(w *aper.AperWriter) (err error) {
// 	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
// 		return
// 	}
// 	switch ie.Choice {
// 	case PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionAggregateMaximumBitRate:
// 		err = ie.PDUSessionAggregateMaximumBitRate.Encode(w)
// 	case PDUSessionResourceSetupRequestTransferIEsPresentULNGUUPTNLInformation:
// 		err = ie.ULNGUUPTNLInformation.Encode(w)
// 	case PDUSessionResourceSetupRequestTransferIEsPresentAdditionalULNGUUPTNLInformation:
// 		err = ie.AdditionalULNGUUPTNLInformation.Encode(w)
// 	case PDUSessionResourceSetupRequestTransferIEsPresentDataForwardingNotPossible:
// 		err = ie.DataForwardingNotPossible.Encode(w)
// 	case PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionType:
// 		err = ie.PDUSessionType.Encode(w)
// 	case PDUSessionResourceSetupRequestTransferIEsPresentSecurityIndication:
// 		err = ie.SecurityIndication.Encode(w)
// 	case PDUSessionResourceSetupRequestTransferIEsPresentNetworkInstance:
// 		err = ie.NetworkInstance.Encode(w)
// 	case PDUSessionResourceSetupRequestTransferIEsPresentQosFlowSetupRequestList:
// 		err = ie.QosFlowSetupRequestList.Encode(w)
// 	}
// 	return
// }
// func (ie *PDUSessionResourceSetupRequestTransfer) Decode(r *aper.AperReader) (err error) {
// 	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
// 		return
// 	}
// 	switch ie.Choice {
// 	case PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionAggregateMaximumBitRate:
// 		var tmp PDUSessionAggregateMaximumBitRate
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.PDUSessionAggregateMaximumBitRate = &tmp
// 	case PDUSessionResourceSetupRequestTransferIEsPresentULNGUUPTNLInformation:
// 		var tmp UPTransportLayerInformation
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.ULNGUUPTNLInformation = &tmp
// 	case PDUSessionResourceSetupRequestTransferIEsPresentAdditionalULNGUUPTNLInformation:
// 		var tmp UPTransportLayerInformationList
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.AdditionalULNGUUPTNLInformation = &tmp
// 	case PDUSessionResourceSetupRequestTransferIEsPresentDataForwardingNotPossible:
// 		var tmp DataForwardingNotPossible
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.DataForwardingNotPossible = &tmp
// 	case PDUSessionResourceSetupRequestTransferIEsPresentPDUSessionType:
// 		var tmp PDUSessionType
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.PDUSessionType = &tmp
// 	case PDUSessionResourceSetupRequestTransferIEsPresentSecurityIndication:
// 		var tmp SecurityIndication
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.SecurityIndication = &tmp
// 	case PDUSessionResourceSetupRequestTransferIEsPresentNetworkInstance:
// 		var tmp NetworkInstance
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.NetworkInstance = &tmp
// 	case PDUSessionResourceSetupRequestTransferIEsPresentQosFlowSetupRequestList:
// 		var tmp QosFlowSetupRequestList
// 		if err = tmp.Decode(r); err != nil {
// 			return
// 		}
// 		ie.QosFlowSetupRequestList = &tmp
// 	}
// 	return
// }
