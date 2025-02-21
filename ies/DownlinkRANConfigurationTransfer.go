package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DownlinkRANConfigurationTransfer struct {
	SONConfigurationTransferDL     *SONConfigurationTransfer `optional,ignore`
	ENDCSONConfigurationTransferDL []byte                    `lb:0,ub:0,optional,ignore`
}

func (msg *DownlinkRANConfigurationTransfer) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		err = msgErrors(fmt.Errorf("DownlinkRANConfigurationTransfer"), err)
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_DownlinkRANConfigurationTransfer, Criticality_PresentIgnore, ies)
}
func (msg *DownlinkRANConfigurationTransfer) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	if msg.SONConfigurationTransferDL != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_SONConfigurationTransferDL},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.SONConfigurationTransferDL,
		})
	}
	if msg.ENDCSONConfigurationTransferDL != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_ENDCSONConfigurationTransferDL},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: msg.ENDCSONConfigurationTransferDL,
			}})
	}
	return
}
func (msg *DownlinkRANConfigurationTransfer) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	defer func() {
		if err != nil {
			err = msgErrors(fmt.Errorf("DownlinkRANConfigurationTransfer"), err)
		}
	}()
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := DownlinkRANConfigurationTransferDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	return
}

type DownlinkRANConfigurationTransferDecoder struct {
	msg      *DownlinkRANConfigurationTransfer
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *DownlinkRANConfigurationTransferDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_SONConfigurationTransferDL:
		var tmp SONConfigurationTransfer
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read SONConfigurationTransferDL", err)
			return
		}
		msg.SONConfigurationTransferDL = &tmp
	case ProtocolIEID_ENDCSONConfigurationTransferDL:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read ENDCSONConfigurationTransferDL", err)
			return
		}
		msg.ENDCSONConfigurationTransferDL = tmp.Value
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
