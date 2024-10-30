package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type WriteReplaceWarningRequest struct {
	MessageIdentifier           *MessageIdentifier           `,reject,mandatory`
	SerialNumber                *SerialNumber                `,reject,mandatory`
	WarningAreaList             *WarningAreaList             `,ignore,optional`
	RepetitionPeriod            *RepetitionPeriod            `,reject,mandatory`
	NumberOfBroadcastsRequested *NumberOfBroadcastsRequested `,reject,mandatory`
	WarningType                 *WarningType                 `,ignore,optional`
	WarningSecurityInfo         *WarningSecurityInfo         `,ignore,optional`
	DataCodingScheme            *DataCodingScheme            `,ignore,optional`
	WarningMessageContents      *WarningMessageContents      `,ignore,optional`
	ConcurrentWarningMessageInd *ConcurrentWarningMessageInd `,reject,optional`
	WarningAreaCoordinates      *WarningAreaCoordinates      `,ignore,optional`
}

func (msg *WriteReplaceWarningRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_WriteReplaceWarning, Criticality_PresentReject, msg.toIes())
}
func (msg *WriteReplaceWarningRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.MessageIdentifier != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_MessageIdentifier},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.MessageIdentifier})
	}
	if msg.SerialNumber != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SerialNumber},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.SerialNumber})
	}
	if msg.WarningAreaList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningAreaList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.WarningAreaList})
	}
	if msg.RepetitionPeriod != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RepetitionPeriod},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RepetitionPeriod})
	}
	if msg.NumberOfBroadcastsRequested != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NumberOfBroadcastsRequested},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NumberOfBroadcastsRequested})
	}
	if msg.WarningType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningType},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.WarningType})
	}
	if msg.WarningSecurityInfo != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningSecurityInfo},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.WarningSecurityInfo})
	}
	if msg.DataCodingScheme != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DataCodingScheme},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.DataCodingScheme})
	}
	if msg.WarningMessageContents != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningMessageContents},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.WarningMessageContents})
	}
	if msg.ConcurrentWarningMessageInd != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ConcurrentWarningMessageInd},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ConcurrentWarningMessageInd})
	}
	if msg.WarningAreaCoordinates != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningAreaCoordinates},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.WarningAreaCoordinates})
	}
	return
}
func (msg *WriteReplaceWarningRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *WriteReplaceWarningRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_MessageIdentifier:
		var tmp MessageIdentifier
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MessageIdentifier = &tmp
	case ProtocolIEID_SerialNumber:
		var tmp SerialNumber
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SerialNumber = &tmp
	case ProtocolIEID_WarningAreaList:
		var tmp WarningAreaList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.WarningAreaList = &tmp
	case ProtocolIEID_RepetitionPeriod:
		var tmp RepetitionPeriod
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RepetitionPeriod = &tmp
	case ProtocolIEID_NumberOfBroadcastsRequested:
		var tmp NumberOfBroadcastsRequested
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NumberOfBroadcastsRequested = &tmp
	case ProtocolIEID_WarningType:
		var tmp WarningType
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.WarningType = &tmp
	case ProtocolIEID_WarningSecurityInfo:
		var tmp WarningSecurityInfo
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.WarningSecurityInfo = &tmp
	case ProtocolIEID_DataCodingScheme:
		var tmp DataCodingScheme
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.DataCodingScheme = &tmp
	case ProtocolIEID_WarningMessageContents:
		var tmp WarningMessageContents
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.WarningMessageContents = &tmp
	case ProtocolIEID_ConcurrentWarningMessageInd:
		var tmp ConcurrentWarningMessageInd
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ConcurrentWarningMessageInd = &tmp
	case ProtocolIEID_WarningAreaCoordinates:
		var tmp WarningAreaCoordinates
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.WarningAreaCoordinates = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
