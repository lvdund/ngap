package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type UplinkRANConfigurationTransfer struct {
	SONConfigurationTransferUL     *SONConfigurationTransfer     `,ignore,optional`
	ENDCSONConfigurationTransferUL *ENDCSONConfigurationTransfer `,ignore,optional`
}

func (msg *UplinkRANConfigurationTransfer) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_UplinkRANConfigurationTransfer, Criticality_PresentIgnore, msg.toIes())
}
func (msg *UplinkRANConfigurationTransfer) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.SONConfigurationTransferUL != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SONConfigurationTransferUL},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SONConfigurationTransferUL})
	}
	if msg.ENDCSONConfigurationTransferUL != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ENDCSONConfigurationTransferUL},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.ENDCSONConfigurationTransferUL})
	}
	return
}
func (msg *UplinkRANConfigurationTransfer) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *UplinkRANConfigurationTransfer) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_SONConfigurationTransferUL:
		var tmp SONConfigurationTransfer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SONConfigurationTransferUL = &tmp
	case ProtocolIEID_ENDCSONConfigurationTransferUL:
		var tmp ENDCSONConfigurationTransfer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ENDCSONConfigurationTransferUL = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
