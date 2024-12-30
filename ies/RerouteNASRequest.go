package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type RerouteNASRequest struct {
	RANUENGAPID                         int64
	AMFUENGAPID                         *int64
	NGAPMessage                         []byte
	AMFSetID                            []byte
	AllowedNSSAI                        *[]AllowedNSSAIItem
	SourceToTargetAMFInformationReroute *SourceToTargetAMFInformationReroute
}

func (msg *RerouteNASRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_RerouteNASRequest, Criticality_PresentReject, msg.toIes())
}
func (msg *RerouteNASRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
				ext:   false,
				Value: aper.Integer(*msg.AMFUENGAPID),
			}})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_NGAPMessage},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.NGAPMessage,
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFSetID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &BITSTRING{
			Value: aper.BitString{
				Bytes:   msg.AMFSetID,
				NumBits: uint64(len(msg.AMFSetID))},
		}})
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
	return
}
func (msg *RerouteNASRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *RerouteNASRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.AMFUENGAPID = int64(tmp.Value)
	case ProtocolIEID_NGAPMessage:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NGAPMessage = tmp.Value
	case ProtocolIEID_AMFSetID:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 10, Ub: 10},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFSetID = tmp.Value.Bytes
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
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
