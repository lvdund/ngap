package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type AMFConfigurationUpdateAcknowledge struct {
	AMFTNLAssociationSetupList         *AMFTNLAssociationSetupList `,ignore,optional`
	AMFTNLAssociationFailedToSetupList *TNLAssociationList         `,ignore,optional`
	CriticalityDiagnostics             *CriticalityDiagnostics     `,ignore,optional`
}

func (msg *AMFConfigurationUpdateAcknowledge) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_AMFConfigurationUpdate, Criticality_PresentReject, msg.toIes())
}
func (msg *AMFConfigurationUpdateAcknowledge) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFTNLAssociationSetupList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFTNLAssociationSetupList})
	}
	if msg.AMFTNLAssociationFailedToSetupList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFTNLAssociationFailedToSetupList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFTNLAssociationFailedToSetupList})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics})
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
	case ProtocolIEID_AMFTNLAssociationSetupList:
		var tmp AMFTNLAssociationSetupList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFTNLAssociationSetupList = &tmp
	case ProtocolIEID_AMFTNLAssociationFailedToSetupList:
		var tmp TNLAssociationList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFTNLAssociationFailedToSetupList = &tmp
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
