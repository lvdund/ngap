package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type RANConfigurationUpdate struct {
	RANNodeName                     *RANNodeName                     `,ignore,optional`
	SupportedTAList                 *SupportedTAList                 `,reject,optional`
	DefaultPagingDRX                *PagingDRX                       `,ignore,optional`
	GlobalRANNodeID                 *GlobalRANNodeID                 `,ignore,optional`
	NGRANTNLAssociationToRemoveList *NGRANTNLAssociationToRemoveList `,reject,optional`
}

func (msg *RANConfigurationUpdate) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_RANConfigurationUpdate, Criticality_PresentReject, msg.toIes())
}
func (msg *RANConfigurationUpdate) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.RANNodeName != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANNodeName},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RANNodeName})
	}
	if msg.SupportedTAList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SupportedTAList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SupportedTAList})
	}
	if msg.DefaultPagingDRX != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DefaultPagingDRX},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DefaultPagingDRX})
	}
	if msg.GlobalRANNodeID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.GlobalRANNodeID})
	}
	if msg.NGRANTNLAssociationToRemoveList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NGRANTNLAssociationToRemoveList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NGRANTNLAssociationToRemoveList})
	}
	return
}
func (msg *RANConfigurationUpdate) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *RANConfigurationUpdate) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_RANNodeName:
		var tmp RANNodeName
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANNodeName = &tmp
	case ProtocolIEID_SupportedTAList:
		var tmp SupportedTAList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SupportedTAList = &tmp
	case ProtocolIEID_DefaultPagingDRX:
		var tmp PagingDRX
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.DefaultPagingDRX = &tmp
	case ProtocolIEID_GlobalRANNodeID:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.GlobalRANNodeID = &tmp
	case ProtocolIEID_NGRANTNLAssociationToRemoveList:
		var tmp NGRANTNLAssociationToRemoveList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NGRANTNLAssociationToRemoveList = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
