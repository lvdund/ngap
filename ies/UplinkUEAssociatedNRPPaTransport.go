package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type UplinkUEAssociatedNRPPaTransport struct {
	AMFUENGAPID *AMFUENGAPID `,reject,mandatory`
	RANUENGAPID *RANUENGAPID `,reject,mandatory`
	RoutingID   *RoutingID   `,reject,mandatory`
	NRPPaPDU    *NRPPaPDU    `,reject,mandatory`
}

func (msg *UplinkUEAssociatedNRPPaTransport) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_UplinkUEAssociatedNRPPaTransport, Criticality_PresentIgnore, msg.toIes())
}
func (msg *UplinkUEAssociatedNRPPaTransport) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RANUENGAPID})
	}
	if msg.RoutingID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RoutingID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RoutingID})
	}
	if msg.NRPPaPDU != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NRPPaPDU},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NRPPaPDU})
	}
	return
}
func (msg *UplinkUEAssociatedNRPPaTransport) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *UplinkUEAssociatedNRPPaTransport) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFUENGAPID:
		var tmp AMFUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = &tmp
	case ProtocolIEID_RANUENGAPID:
		var tmp RANUENGAPID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = &tmp
	case ProtocolIEID_RoutingID:
		var tmp RoutingID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RoutingID = &tmp
	case ProtocolIEID_NRPPaPDU:
		var tmp NRPPaPDU
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NRPPaPDU = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
