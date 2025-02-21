package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AMFConfigurationUpdateAcknowledge struct {
	AMFTNLAssociationSetupList         []AMFTNLAssociationSetupItem `lb:1,ub:maxnoofTNLAssociations,optional,ignore`
	AMFTNLAssociationFailedToSetupList []TNLAssociationItem         `lb:1,ub:maxnoofTNLAssociations,optional,ignore`
	CriticalityDiagnostics             *CriticalityDiagnostics      `optional,ignore`
}

func (msg *AMFConfigurationUpdateAcknowledge) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("AMFConfigurationUpdateAcknowledge"), err)
		return
	}
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_AMFConfigurationUpdate, Criticality_PresentReject, ies)
}
func (msg *AMFConfigurationUpdateAcknowledge) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	if len(msg.AMFTNLAssociationSetupList) > 0 {
		tmp_AMFTNLAssociationSetupList := Sequence[*AMFTNLAssociationSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		for _, i := range msg.AMFTNLAssociationSetupList {
			tmp_AMFTNLAssociationSetupList.Value = append(tmp_AMFTNLAssociationSetupList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_AMFTNLAssociationSetupList,
		})
	}
	if len(msg.AMFTNLAssociationFailedToSetupList) > 0 {
		tmp_AMFTNLAssociationFailedToSetupList := Sequence[*TNLAssociationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		for _, i := range msg.AMFTNLAssociationFailedToSetupList {
			tmp_AMFTNLAssociationFailedToSetupList.Value = append(tmp_AMFTNLAssociationFailedToSetupList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationFailedToSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_AMFTNLAssociationFailedToSetupList,
		})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return
}
func (msg *AMFConfigurationUpdateAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("AMFConfigurationUpdateAcknowledge"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := AMFConfigurationUpdateAcknowledgeDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	return
}

type AMFConfigurationUpdateAcknowledgeDecoder struct {
	msg      *AMFConfigurationUpdateAcknowledge
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *AMFConfigurationUpdateAcknowledgeDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFTNLAssociationSetupList:
		tmp := Sequence[*AMFTNLAssociationSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		fn := func() *AMFTNLAssociationSetupItem { return new(AMFTNLAssociationSetupItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AMFTNLAssociationSetupList", err)
			return
		}
		msg.AMFTNLAssociationSetupList = []AMFTNLAssociationSetupItem{}
		for _, i := range tmp.Value {
			msg.AMFTNLAssociationSetupList = append(msg.AMFTNLAssociationSetupList, *i)
		}
	case ProtocolIEID_AMFTNLAssociationFailedToSetupList:
		tmp := Sequence[*TNLAssociationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		fn := func() *TNLAssociationItem { return new(TNLAssociationItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read AMFTNLAssociationFailedToSetupList", err)
			return
		}
		msg.AMFTNLAssociationFailedToSetupList = []TNLAssociationItem{}
		for _, i := range tmp.Value {
			msg.AMFTNLAssociationFailedToSetupList = append(msg.AMFTNLAssociationFailedToSetupList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
		msg.CriticalityDiagnostics = &tmp
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
