package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type UEContextReleaseCommand struct {
	UENGAPIDs *UENGAPIDs `,reject,mandatory`
	Cause     *Cause     `,ignore,mandatory`
}

func (msg *UEContextReleaseCommand) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_UEContextRelease, Criticality_PresentReject, msg.toIes())
}
func (msg *UEContextReleaseCommand) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.UENGAPIDs != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UENGAPIDs},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.UENGAPIDs})
	}
	if msg.Cause != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_Cause},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.Cause})
	}
	return
}
func (msg *UEContextReleaseCommand) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *UEContextReleaseCommand) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_UENGAPIDs:
		var tmp UENGAPIDs
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UENGAPIDs = &tmp
	case ProtocolIEID_Cause:
		var tmp Cause
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.Cause = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
