package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AMFConfigurationUpdate struct {
	AMFName                       []byte                          `lb:1,ub:150,optional,reject,valueExt`
	ServedGUAMIList               []ServedGUAMIItem               `lb:1,ub:maxnoofServedGUAMIs,optional,reject`
	RelativeAMFCapacity           *int64                          `lb:0,ub:255,optional,ignore`
	PLMNSupportList               []PLMNSupportItem               `lb:1,ub:maxnoofPLMNs,optional,reject`
	AMFTNLAssociationToAddList    []AMFTNLAssociationToAddItem    `lb:1,ub:maxnoofTNLAssociations,optional,ignore`
	AMFTNLAssociationToRemoveList []AMFTNLAssociationToRemoveItem `lb:1,ub:maxnoofTNLAssociations,optional,ignore`
	AMFTNLAssociationToUpdateList []AMFTNLAssociationToUpdateItem `lb:1,ub:maxnoofTNLAssociations,optional,ignore`
}

func (msg *AMFConfigurationUpdate) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_AMFConfigurationUpdate, Criticality_PresentReject, ies)
}
func (msg *AMFConfigurationUpdate) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	if msg.AMFName != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFName},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 150},
				ext:   true,
				Value: msg.AMFName,
			}})
	}
	if len(msg.ServedGUAMIList) > 0 {
		tmp_ServedGUAMIList := Sequence[*ServedGUAMIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofServedGUAMIs},
			ext: false,
		}
		for _, i := range msg.ServedGUAMIList {
			tmp_ServedGUAMIList.Value = append(tmp_ServedGUAMIList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ServedGUAMIList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_ServedGUAMIList,
		})
	}
	if msg.RelativeAMFCapacity != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RelativeAMFCapacity},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 255},
				ext:   false,
				Value: aper.Integer(*msg.RelativeAMFCapacity),
			}})
	}
	if len(msg.PLMNSupportList) > 0 {
		tmp_PLMNSupportList := Sequence[*PLMNSupportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPLMNs},
			ext: false,
		}
		for _, i := range msg.PLMNSupportList {
			tmp_PLMNSupportList.Value = append(tmp_PLMNSupportList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PLMNSupportList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_PLMNSupportList,
		})
	}
	if len(msg.AMFTNLAssociationToAddList) > 0 {
		tmp_AMFTNLAssociationToAddList := Sequence[*AMFTNLAssociationToAddItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		for _, i := range msg.AMFTNLAssociationToAddList {
			tmp_AMFTNLAssociationToAddList.Value = append(tmp_AMFTNLAssociationToAddList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationToAddList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_AMFTNLAssociationToAddList,
		})
	}
	if len(msg.AMFTNLAssociationToRemoveList) > 0 {
		tmp_AMFTNLAssociationToRemoveList := Sequence[*AMFTNLAssociationToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		for _, i := range msg.AMFTNLAssociationToRemoveList {
			tmp_AMFTNLAssociationToRemoveList.Value = append(tmp_AMFTNLAssociationToRemoveList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationToRemoveList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_AMFTNLAssociationToRemoveList,
		})
	}
	if len(msg.AMFTNLAssociationToUpdateList) > 0 {
		tmp_AMFTNLAssociationToUpdateList := Sequence[*AMFTNLAssociationToUpdateItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		for _, i := range msg.AMFTNLAssociationToUpdateList {
			tmp_AMFTNLAssociationToUpdateList.Value = append(tmp_AMFTNLAssociationToUpdateList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationToUpdateList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_AMFTNLAssociationToUpdateList,
		})
	}
	return
}
func (msg *AMFConfigurationUpdate) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := AMFConfigurationUpdateDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	return
}

type AMFConfigurationUpdateDecoder struct {
	msg      *AMFConfigurationUpdate
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *AMFConfigurationUpdateDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFName:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read AMFName", err)
			return
		}
		msg.AMFName = tmp.Value
	case ProtocolIEID_ServedGUAMIList:
		tmp := Sequence[*ServedGUAMIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofServedGUAMIs},
			ext: false,
		}
		fn := func() *ServedGUAMIItem { return new(ServedGUAMIItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read ServedGUAMIList", err)
			return
		}
		msg.ServedGUAMIList = []ServedGUAMIItem{}
		for _, i := range tmp.Value {
			msg.ServedGUAMIList = append(msg.ServedGUAMIList, *i)
		}
	case ProtocolIEID_RelativeAMFCapacity:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RelativeAMFCapacity", err)
			return
		}
		*msg.RelativeAMFCapacity = int64(tmp.Value)
	case ProtocolIEID_PLMNSupportList:
		tmp := Sequence[*PLMNSupportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPLMNs},
			ext: false,
		}
		fn := func() *PLMNSupportItem { return new(PLMNSupportItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read PLMNSupportList", err)
			return
		}
		msg.PLMNSupportList = []PLMNSupportItem{}
		for _, i := range tmp.Value {
			msg.PLMNSupportList = append(msg.PLMNSupportList, *i)
		}
	case ProtocolIEID_AMFTNLAssociationToAddList:
		tmp := Sequence[*AMFTNLAssociationToAddItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		fn := func() *AMFTNLAssociationToAddItem { return new(AMFTNLAssociationToAddItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AMFTNLAssociationToAddList", err)
			return
		}
		msg.AMFTNLAssociationToAddList = []AMFTNLAssociationToAddItem{}
		for _, i := range tmp.Value {
			msg.AMFTNLAssociationToAddList = append(msg.AMFTNLAssociationToAddList, *i)
		}
	case ProtocolIEID_AMFTNLAssociationToRemoveList:
		tmp := Sequence[*AMFTNLAssociationToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		fn := func() *AMFTNLAssociationToRemoveItem { return new(AMFTNLAssociationToRemoveItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AMFTNLAssociationToRemoveList", err)
			return
		}
		msg.AMFTNLAssociationToRemoveList = []AMFTNLAssociationToRemoveItem{}
		for _, i := range tmp.Value {
			msg.AMFTNLAssociationToRemoveList = append(msg.AMFTNLAssociationToRemoveList, *i)
		}
	case ProtocolIEID_AMFTNLAssociationToUpdateList:
		tmp := Sequence[*AMFTNLAssociationToUpdateItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		fn := func() *AMFTNLAssociationToUpdateItem { return new(AMFTNLAssociationToUpdateItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AMFTNLAssociationToUpdateList", err)
			return
		}
		msg.AMFTNLAssociationToUpdateList = []AMFTNLAssociationToUpdateItem{}
		for _, i := range tmp.Value {
			msg.AMFTNLAssociationToUpdateList = append(msg.AMFTNLAssociationToUpdateList, *i)
		}
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
