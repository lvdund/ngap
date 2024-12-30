package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type InitialUEMessage struct {
	RANUENGAPID                         int64
	NASPDU                              []byte
	UserLocationInformation             UserLocationInformation
	RRCEstablishmentCause               RRCEstablishmentCause
	FiveGSTMSI                          *FiveGSTMSI
	AMFSetID                            *[]byte
	UEContextRequest                    *UEContextRequest
	AllowedNSSAI                        *[]AllowedNSSAIItem
	SourceToTargetAMFInformationReroute *SourceToTargetAMFInformationReroute
	SelectedPLMNIdentity                *[]byte
}

func (msg *InitialUEMessage) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_InitialUEMessage, Criticality_PresentIgnore, msg.toIes())
}
func (msg *InitialUEMessage) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.NASPDU,
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_UserLocationInformation},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.UserLocationInformation,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RRCEstablishmentCause},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &msg.RRCEstablishmentCause,
	})
	if msg.FiveGSTMSI != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_FiveGSTMSI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.FiveGSTMSI,
		})
	}
	if msg.AMFSetID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFSetID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &BITSTRING{
				Value: aper.BitString{
					Bytes:   *msg.AMFSetID,
					NumBits: uint64(len(*msg.AMFSetID))},
			}})
	}
	if msg.UEContextRequest != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEContextRequest},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UEContextRequest,
		})
	}
	if msg.AllowedNSSAI != nil {
		tmp_AllowedNSSAI := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *msg.AllowedNSSAI {
			tmp_AllowedNSSAI.Value = append(tmp_AllowedNSSAI.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AllowedNSSAI},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_AllowedNSSAI,
		})
	}
	if msg.SourceToTargetAMFInformationReroute != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SourceToTargetAMFInformationReroute},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SourceToTargetAMFInformationReroute,
		})
	}
	if msg.SelectedPLMNIdentity != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SelectedPLMNIdentity},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 3, Ub: 3},
				ext:   false,
				Value: *msg.SelectedPLMNIdentity,
			}})
	}
	return
}
func (msg *InitialUEMessage) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *InitialUEMessage) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_RANUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = int64(tmp.Value)
	case ProtocolIEID_NASPDU:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NASPDU = tmp.Value
	case ProtocolIEID_UserLocationInformation:
		var tmp UserLocationInformation
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UserLocationInformation = tmp
	case ProtocolIEID_RRCEstablishmentCause:
		var tmp RRCEstablishmentCause
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RRCEstablishmentCause = tmp
	case ProtocolIEID_FiveGSTMSI:
		var tmp FiveGSTMSI
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.FiveGSTMSI = &tmp
	case ProtocolIEID_AMFSetID:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 10, Ub: 10},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.AMFSetID = tmp.Value.Bytes
	case ProtocolIEID_UEContextRequest:
		var tmp UEContextRequest
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEContextRequest = &tmp
	case ProtocolIEID_AllowedNSSAI:
		tmp := Sequence[*AllowedNSSAIItem]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AllowedNSSAI = &[]AllowedNSSAIItem{}
		for _, i := range tmp.Value {
			*msg.AllowedNSSAI = append(*msg.AllowedNSSAI, *i)
		}
	case ProtocolIEID_SourceToTargetAMFInformationReroute:
		var tmp SourceToTargetAMFInformationReroute
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SourceToTargetAMFInformationReroute = &tmp
	case ProtocolIEID_SelectedPLMNIdentity:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 3, Ub: 3},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.SelectedPLMNIdentity = tmp.Value
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
