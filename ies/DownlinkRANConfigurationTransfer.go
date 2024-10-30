package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type DownlinkRANConfigurationTransfer struct {
	SONConfigurationTransferDL     *SONConfigurationTransfer     `,ignore,optional`
	ENDCSONConfigurationTransferDL *ENDCSONConfigurationTransfer `,ignore,optional`
}

func (msg *DownlinkRANConfigurationTransfer) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_DownlinkRANConfigurationTransfer, Criticality_PresentIgnore, msg.toIes())
}
func (msg *DownlinkRANConfigurationTransfer) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.SONConfigurationTransferDL != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SONConfigurationTransferDL},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SONConfigurationTransferDL})
	}
	if msg.ENDCSONConfigurationTransferDL != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ENDCSONConfigurationTransferDL},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.ENDCSONConfigurationTransferDL})
	}
	return
}
func (msg *DownlinkRANConfigurationTransfer) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *DownlinkRANConfigurationTransfer) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_SONConfigurationTransferDL:
		var tmp SONConfigurationTransfer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.SONConfigurationTransferDL = &tmp
	case ProtocolIEID_ENDCSONConfigurationTransferDL:
		var tmp ENDCSONConfigurationTransfer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.ENDCSONConfigurationTransferDL = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
