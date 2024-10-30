package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type AMFConfigurationUpdate struct {
	AMFName                       *AMFName                       `,reject,optional`
	ServedGUAMIList               *ServedGUAMIList               `,reject,optional`
	RelativeAMFCapacity           *RelativeAMFCapacity           `,ignore,optional`
	PLMNSupportList               *PLMNSupportList               `,reject,optional`
	AMFTNLAssociationToAddList    *AMFTNLAssociationToAddList    `,ignore,optional`
	AMFTNLAssociationToRemoveList *AMFTNLAssociationToRemoveList `,ignore,optional`
	AMFTNLAssociationToUpdateList *AMFTNLAssociationToUpdateList `,ignore,optional`
}

func (msg *AMFConfigurationUpdate) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_AMFConfigurationUpdate, Criticality_PresentReject, msg.toIes())
}
func (msg *AMFConfigurationUpdate) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFName != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFName},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFName})
	}
	if msg.ServedGUAMIList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServedGUAMIList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ServedGUAMIList})
	}
	if msg.RelativeAMFCapacity != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RelativeAMFCapacity},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RelativeAMFCapacity})
	}
	if msg.PLMNSupportList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PLMNSupportList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PLMNSupportList})
	}
	if msg.AMFTNLAssociationToAddList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationToAddList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFTNLAssociationToAddList})
	}
	if msg.AMFTNLAssociationToRemoveList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationToRemoveList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFTNLAssociationToRemoveList})
	}
	if msg.AMFTNLAssociationToUpdateList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationToUpdateList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFTNLAssociationToUpdateList})
	}
	return
}
func (msg *AMFConfigurationUpdate) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *AMFConfigurationUpdate) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFName:
		var tmp AMFName
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFName = &tmp
	case ProtocolIEID_ServedGUAMIList:
		var tmp ServedGUAMIList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ServedGUAMIList = &tmp
	case ProtocolIEID_RelativeAMFCapacity:
		var tmp RelativeAMFCapacity
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RelativeAMFCapacity = &tmp
	case ProtocolIEID_PLMNSupportList:
		var tmp PLMNSupportList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PLMNSupportList = &tmp
	case ProtocolIEID_AMFTNLAssociationToAddList:
		var tmp AMFTNLAssociationToAddList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFTNLAssociationToAddList = &tmp
	case ProtocolIEID_AMFTNLAssociationToRemoveList:
		var tmp AMFTNLAssociationToRemoveList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFTNLAssociationToRemoveList = &tmp
	case ProtocolIEID_AMFTNLAssociationToUpdateList:
		var tmp AMFTNLAssociationToUpdateList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFTNLAssociationToUpdateList = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
