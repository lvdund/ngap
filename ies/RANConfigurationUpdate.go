package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type RANConfigurationUpdate struct {
	RANNodeName                     *[]byte
	SupportedTAList                 *[]SupportedTAItem
	DefaultPagingDRX                *PagingDRX
	GlobalRANNodeID                 *GlobalRANNodeID
	NGRANTNLAssociationToRemoveList *[]NGRANTNLAssociationToRemoveItem
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
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 150},
				ext:   true,
				Value: *msg.RANNodeName,
			}})
	}
	if msg.SupportedTAList != nil {
		tmp_SupportedTAList := Sequence[*SupportedTAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTACs},
			ext: false,
		}
		for _, i := range *msg.SupportedTAList {
			tmp_SupportedTAList.Value = append(tmp_SupportedTAList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SupportedTAList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_SupportedTAList,
		})
	}
	if msg.DefaultPagingDRX != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DefaultPagingDRX},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DefaultPagingDRX,
		})
	}
	if msg.GlobalRANNodeID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.GlobalRANNodeID,
		})
	}
	if msg.NGRANTNLAssociationToRemoveList != nil {
		tmp_NGRANTNLAssociationToRemoveList := Sequence[*NGRANTNLAssociationToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		for _, i := range *msg.NGRANTNLAssociationToRemoveList {
			tmp_NGRANTNLAssociationToRemoveList.Value = append(tmp_NGRANTNLAssociationToRemoveList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NGRANTNLAssociationToRemoveList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_NGRANTNLAssociationToRemoveList,
		})
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
	case ProtocolIEID_RANNodeName:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.RANNodeName = tmp.Value
	case ProtocolIEID_SupportedTAList:
		tmp := Sequence[*SupportedTAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTACs},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SupportedTAList = &[]SupportedTAItem{}
		for _, i := range tmp.Value {
			*msg.SupportedTAList = append(*msg.SupportedTAList, *i)
		}
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
		tmp := Sequence[*NGRANTNLAssociationToRemoveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTNLAssociations},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NGRANTNLAssociationToRemoveList = &[]NGRANTNLAssociationToRemoveItem{}
		for _, i := range tmp.Value {
			*msg.NGRANTNLAssociationToRemoveList = append(*msg.NGRANTNLAssociationToRemoveList, *i)
		}
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
