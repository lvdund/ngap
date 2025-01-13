package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AMFStatusIndication struct {
	UnavailableGUAMIList []UnavailableGUAMIItem `lb:1,ub:maxnoofServedGUAMIs,mandatory,reject`
}

func (msg *AMFStatusIndication) Encode(w io.Writer) (err error) {
	var ies []NgapMessageIE
	if ies, err = msg.toIes(); err != nil {
		return
	}
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_AMFStatusIndication, Criticality_PresentIgnore, ies)
}
func (msg *AMFStatusIndication) toIes() (ies []NgapMessageIE, err error) {
	ies = []NgapMessageIE{}
	if len(msg.UnavailableGUAMIList) > 0 {
		tmp_UnavailableGUAMIList := Sequence[*UnavailableGUAMIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofServedGUAMIs},
			ext: false,
		}
		for _, i := range msg.UnavailableGUAMIList {
			tmp_UnavailableGUAMIList.Value = append(tmp_UnavailableGUAMIList.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UnavailableGUAMIList},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_UnavailableGUAMIList,
		})
	} else {
		err = utils.WrapError("UnavailableGUAMIList is nil", err)
		return
	}
	return
}
func (msg *AMFStatusIndication) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := AMFStatusIndicationDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_UnavailableGUAMIList]; !ok {
		err = fmt.Errorf("Mandatory field UnavailableGUAMIList is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_UnavailableGUAMIList},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type AMFStatusIndicationDecoder struct {
	msg      *AMFStatusIndication
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *AMFStatusIndicationDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_UnavailableGUAMIList:
		tmp := Sequence[*UnavailableGUAMIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofServedGUAMIs},
			ext: false,
		}
		fn := func() *UnavailableGUAMIItem { return new(UnavailableGUAMIItem) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read UnavailableGUAMIList", err)
			return
		}
		msg.UnavailableGUAMIList = []UnavailableGUAMIItem{}
		for _, i := range tmp.Value {
			msg.UnavailableGUAMIList = append(msg.UnavailableGUAMIList, *i)
		}
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
