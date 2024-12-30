package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type AMFConfigurationUpdateAcknowledge struct {
	AMFTNLAssociationSetupList         *[]AMFTNLAssociationSetupItem
	AMFTNLAssociationFailedToSetupList *[]TNLAssociationItem
	CriticalityDiagnostics             *CriticalityDiagnostics
}

func (msg *AMFConfigurationUpdateAcknowledge) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_AMFConfigurationUpdate, Criticality_PresentReject, msg.toIes())
}
func (msg *AMFConfigurationUpdateAcknowledge) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFTNLAssociationSetupList != nil {
		tmp_AMFTNLAssociationSetupList := Sequence[*AMFTNLAssociationSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		for _, i := range *msg.AMFTNLAssociationSetupList {
			tmp_AMFTNLAssociationSetupList.Value = append(tmp_AMFTNLAssociationSetupList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_AMFTNLAssociationSetupList,
		})
	}
	if msg.AMFTNLAssociationFailedToSetupList != nil {
		tmp_AMFTNLAssociationFailedToSetupList := Sequence[*TNLAssociationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		for _, i := range *msg.AMFTNLAssociationFailedToSetupList {
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
func (msg *AMFConfigurationUpdateAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *AMFConfigurationUpdateAcknowledge) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFTNLAssociationSetupList:
		tmp := Sequence[*AMFTNLAssociationSetupItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFTNLAssociationSetupList = &[]AMFTNLAssociationSetupItem{}
		for _, i := range tmp.Value {
			*msg.AMFTNLAssociationSetupList = append(*msg.AMFTNLAssociationSetupList, *i)
		}
	case ProtocolIEID_AMFTNLAssociationFailedToSetupList:
		tmp := Sequence[*TNLAssociationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFTNLAssociationFailedToSetupList = &[]TNLAssociationItem{}
		for _, i := range tmp.Value {
			*msg.AMFTNLAssociationFailedToSetupList = append(*msg.AMFTNLAssociationFailedToSetupList, *i)
		}
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CriticalityDiagnostics = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
