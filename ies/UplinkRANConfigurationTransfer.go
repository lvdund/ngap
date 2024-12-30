package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type UplinkRANConfigurationTransfer struct {
	SONConfigurationTransferUL     *SONConfigurationTransfer
	ENDCSONConfigurationTransferUL *[]byte
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
			Value:       msg.SONConfigurationTransferUL,
		})
	}
	if msg.ENDCSONConfigurationTransferUL != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ENDCSONConfigurationTransferUL},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: *msg.ENDCSONConfigurationTransferUL,
			}})
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
	case ProtocolIEID_SONConfigurationTransferUL:
		var tmp SONConfigurationTransfer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SONConfigurationTransferUL = &tmp
	case ProtocolIEID_ENDCSONConfigurationTransferUL:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.ENDCSONConfigurationTransferUL = tmp.Value
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
