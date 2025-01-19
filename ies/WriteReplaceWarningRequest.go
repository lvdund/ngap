package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type WriteReplaceWarningRequest struct {
	MessageIdentifier           []byte                       `lb:16,ub:16,mandatory,reject`
	SerialNumber                []byte                       `lb:16,ub:16,mandatory,reject`
	WarningAreaList             *WarningAreaList             `optional,mandatory,ignore`
	RepetitionPeriod            int64                        `lb:0,ub:131071,mandatory,reject`
	NumberOfBroadcastsRequested int64                        `lb:0,ub:65535,mandatory,reject`
	WarningType                 []byte                       `lb:2,ub:2,optional,mandatory,ignore`
	WarningSecurityInfo         []byte                       `lb:50,ub:50,optional,mandatory,ignore`
	DataCodingScheme            []byte                       `lb:8,ub:8,optional,mandatory,ignore`
	WarningMessageContents      []byte                       `lb:1,ub:9600,optional,mandatory,ignore`
	ConcurrentWarningMessageInd *ConcurrentWarningMessageInd `optional,mandatory,reject`
	WarningAreaCoordinates      []byte                       `lb:1,ub:1024,optional,mandatory,ignore`
}

func (msg *WriteReplaceWarningRequest) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_WriteReplaceWarning, Criticality_PresentReject, ies)
}
func (msg *WriteReplaceWarningRequest) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_MessageIdentifier},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
			Value: aper.BitString{
				Bytes: msg.MessageIdentifier},
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_SerialNumber},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
			Value: aper.BitString{
				Bytes: msg.SerialNumber},
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
				Value: msg.WarningType,
			}})
	}
	if msg.WarningSecurityInfo != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningSecurityInfo},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 50, Ub: 50},
				ext:   false,
				Value: msg.WarningSecurityInfo,
			}})
	}
	if msg.DataCodingScheme != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_DataCodingScheme},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				c:   aper.Constraint{Lb: 8, Ub: 8},
				ext: false,
				Value: aper.BitString{
					Bytes: msg.DataCodingScheme},
			}})
	}
	if msg.WarningMessageContents != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_WarningMessageContents},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 1, Ub: 9600},
				ext:   false,
				Value: msg.WarningMessageContents,
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
				Value: msg.WarningAreaCoordinates,
			}})
	}
	return
}
func (msg *WriteReplaceWarningRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("WriteReplaceWarningRequest"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := WriteReplaceWarningRequestDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_MessageIdentifier]; !ok {
		err = fmt.Errorf("Mandatory field MessageIdentifier is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_MessageIdentifier},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_SerialNumber]; !ok {
		err = fmt.Errorf("Mandatory field SerialNumber is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_SerialNumber},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_RepetitionPeriod]; !ok {
		err = fmt.Errorf("Mandatory field RepetitionPeriod is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_RepetitionPeriod},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_NumberOfBroadcastsRequested]; !ok {
		err = fmt.Errorf("Mandatory field NumberOfBroadcastsRequested is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_NumberOfBroadcastsRequested},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type WriteReplaceWarningRequestDecoder struct {
	msg      *WriteReplaceWarningRequest
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *WriteReplaceWarningRequestDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_MessageIdentifier:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read MessageIdentifier", err)
			return
		}
		msg.MessageIdentifier = tmp.Value.Bytes
	case ProtocolIEID_SerialNumber:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 16, Ub: 16},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SerialNumber", err)
			return
		}
		msg.SerialNumber = tmp.Value.Bytes
	case ProtocolIEID_WarningAreaList:
		var tmp WarningAreaList
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read WarningAreaList", err)
			return
		}
		msg.WarningAreaList = &tmp
	case ProtocolIEID_RepetitionPeriod:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 131071},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read RepetitionPeriod", err)
			return
		}
		msg.RepetitionPeriod = int64(tmp.Value)
	case ProtocolIEID_NumberOfBroadcastsRequested:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 65535},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read NumberOfBroadcastsRequested", err)
			return
		}
		msg.NumberOfBroadcastsRequested = int64(tmp.Value)
	case ProtocolIEID_WarningType:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 2, Ub: 2},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read WarningType", err)
			return
		}
		msg.WarningType = tmp.Value
	case ProtocolIEID_WarningSecurityInfo:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 50, Ub: 50},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read WarningSecurityInfo", err)
			return
		}
		msg.WarningSecurityInfo = tmp.Value
	case ProtocolIEID_DataCodingScheme:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 8, Ub: 8},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read DataCodingScheme", err)
			return
		}
		msg.DataCodingScheme = tmp.Value.Bytes
	case ProtocolIEID_WarningMessageContents:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 9600},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read WarningMessageContents", err)
			return
		}
		msg.WarningMessageContents = tmp.Value
	case ProtocolIEID_ConcurrentWarningMessageInd:
		var tmp ConcurrentWarningMessageInd
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ConcurrentWarningMessageInd", err)
			return
		}
		msg.ConcurrentWarningMessageInd = &tmp
	case ProtocolIEID_WarningAreaCoordinates:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 1024},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read WarningAreaCoordinates", err)
			return
		}
		msg.WarningAreaCoordinates = tmp.Value
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
