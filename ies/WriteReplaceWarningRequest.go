package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type WriteReplaceWarningRequest struct {
	MessageIdentifier           []byte
	SerialNumber                []byte
	WarningAreaList             *WarningAreaList
	RepetitionPeriod            int64
	NumberOfBroadcastsRequested int64
	WarningType                 *[]byte
	WarningSecurityInfo         *[]byte
	DataCodingScheme            *[]byte
	WarningMessageContents      *[]byte
	ConcurrentWarningMessageInd *ConcurrentWarningMessageInd
	WarningAreaCoordinates      *[]byte
}

func (msg *WriteReplaceWarningRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_WriteReplaceWarning, Criticality_PresentReject, msg.toIes())
}
func (msg *WriteReplaceWarningRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_MessageIdentifier},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			Value: aper.BitString{
				Bytes:   msg.MessageIdentifier,
				NumBits: uint64(len(msg.MessageIdentifier))},
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SerialNumber},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			Value: aper.BitString{
				Bytes:   msg.SerialNumber,
				NumBits: uint64(len(msg.SerialNumber))},
		}})
	if msg.WarningAreaList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningAreaList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.WarningAreaList,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RepetitionPeriod},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 131071},
			ext:   false,
			Value: aper.Integer(msg.RepetitionPeriod),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NumberOfBroadcastsRequested},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 65535},
			ext:   false,
			Value: aper.Integer(msg.NumberOfBroadcastsRequested),
		}})
	if msg.WarningType != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningType},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 2, Ub: 2},
				ext:   false,
				Value: *msg.WarningType,
			}})
	}
	if msg.WarningSecurityInfo != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningSecurityInfo},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 50, Ub: 50},
				ext:   false,
				Value: *msg.WarningSecurityInfo,
			}})
	}
	if msg.DataCodingScheme != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DataCodingScheme},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				Value: aper.BitString{
					Bytes:   *msg.DataCodingScheme,
					NumBits: uint64(len(*msg.DataCodingScheme))},
			}})
	}
	if msg.WarningMessageContents != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningMessageContents},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 9600},
				ext:   false,
				Value: *msg.WarningMessageContents,
			}})
	}
	if msg.ConcurrentWarningMessageInd != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ConcurrentWarningMessageInd},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.ConcurrentWarningMessageInd,
		})
	}
	if msg.WarningAreaCoordinates != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningAreaCoordinates},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 1024},
				ext:   false,
				Value: *msg.WarningAreaCoordinates,
			}})
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
	case ProtocolIEID_MessageIdentifier:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.MessageIdentifier = tmp.Value.Bytes
	case ProtocolIEID_SerialNumber:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SerialNumber = tmp.Value.Bytes
	case ProtocolIEID_WarningAreaList:
		var tmp WarningAreaList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.WarningAreaList = &tmp
	case ProtocolIEID_RepetitionPeriod:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 131071},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RepetitionPeriod = int64(tmp.Value)
	case ProtocolIEID_NumberOfBroadcastsRequested:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 65535},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NumberOfBroadcastsRequested = int64(tmp.Value)
	case ProtocolIEID_WarningType:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 2, Ub: 2},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.WarningType = tmp.Value
	case ProtocolIEID_WarningSecurityInfo:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 50, Ub: 50},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.WarningSecurityInfo = tmp.Value
	case ProtocolIEID_DataCodingScheme:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 8, Ub: 8},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.DataCodingScheme = tmp.Value.Bytes
	case ProtocolIEID_WarningMessageContents:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 9600},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.WarningMessageContents = tmp.Value
	case ProtocolIEID_ConcurrentWarningMessageInd:
		var tmp ConcurrentWarningMessageInd
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ConcurrentWarningMessageInd = &tmp
	case ProtocolIEID_WarningAreaCoordinates:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 1024},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.WarningAreaCoordinates = tmp.Value
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
