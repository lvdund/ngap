package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PWSRestartIndication struct {
	CellIDListForRestart          CellIDListForRestart
	GlobalRANNodeID               GlobalRANNodeID
	TAIListForRestart             []TAI
	EmergencyAreaIDListForRestart []EmergencyAreaID `optional`
}

func (msg *PWSRestartIndication) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PWSRestartIndication, Criticality_PresentIgnore, msg.toIes())
}
func (msg *PWSRestartIndication) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_CellIDListForRestart},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.CellIDListForRestart,
	})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &msg.GlobalRANNodeID,
	})
	tmp_TAIListForRestart := Sequence[*TAI]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforRestart},
		ext: false,
	}
	for _, i := range msg.TAIListForRestart {
		tmp_TAIListForRestart.Value = append(tmp_TAIListForRestart.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TAIListForRestart},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_TAIListForRestart,
	})
	if msg.EmergencyAreaIDListForRestart != nil {
		tmp_EmergencyAreaIDListForRestart := Sequence[*EmergencyAreaID]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEAIforRestart},
			ext: false,
		}
		for _, i := range msg.EmergencyAreaIDListForRestart {
			tmp_EmergencyAreaIDListForRestart.Value = append(tmp_EmergencyAreaIDListForRestart.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_EmergencyAreaIDListForRestart},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       &tmp_EmergencyAreaIDListForRestart,
		})
	}
	return
}
func (msg *PWSRestartIndication) Decode(wire []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	decoder := PWSRestartIndicationDecoder{
		msg:  msg,
		list: make(map[aper.Integer]*NgapMessageIE),
	}
	if _, err = aper.ReadSequenceOf[NgapMessageIE](decoder.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	if _, ok := decoder.list[ProtocolIEID_CellIDListForRestart]; !ok {
		err = fmt.Errorf("Mandatory field CellIDListForRestart is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_CellIDListForRestart},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_GlobalRANNodeID]; !ok {
		err = fmt.Errorf("Mandatory field GlobalRANNodeID is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_GlobalRANNodeID},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	if _, ok := decoder.list[ProtocolIEID_TAIListForRestart]; !ok {
		err = fmt.Errorf("Mandatory field TAIListForRestart is missing")
		decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{
			IECriticality: Criticality{Value: Criticality_PresentReject},
			IEID:          ProtocolIEID{Value: ProtocolIEID_TAIListForRestart},
			TypeOfError:   TypeOfError{Value: TypeOfErrorMissing},
		})
		return
	}
	return
}

type PWSRestartIndicationDecoder struct {
	msg      *PWSRestartIndication
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*NgapMessageIE
}

func (decoder *PWSRestartIndicationDecoder) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_CellIDListForRestart:
		var tmp CellIDListForRestart
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read CellIDListForRestart", err)
			return
		}
		msg.CellIDListForRestart = tmp
	case ProtocolIEID_GlobalRANNodeID:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(ieR); err != nil {
			err = utils.WrapError("Read GlobalRANNodeID", err)
			return
		}
		msg.GlobalRANNodeID = tmp
	case ProtocolIEID_TAIListForRestart:
		tmp := Sequence[*TAI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforRestart},
			ext: false,
		}
		fn := func() *TAI { return new(TAI) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read TAIListForRestart", err)
			return
		}
		msg.TAIListForRestart = []TAI{}
		for _, i := range tmp.Value {
			msg.TAIListForRestart = append(msg.TAIListForRestart, *i)
		}
	case ProtocolIEID_EmergencyAreaIDListForRestart:
		tmp := Sequence[*EmergencyAreaID]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEAIforRestart},
			ext: false,
		}
		fn := func() *EmergencyAreaID { return new(EmergencyAreaID) }
		if err = tmp.Decode(ieR, fn); err != nil {
			err = utils.WrapError("Read EmergencyAreaIDListForRestart", err)
			return
		}
		msg.EmergencyAreaIDListForRestart = []EmergencyAreaID{}
		for _, i := range tmp.Value {
			msg.EmergencyAreaIDListForRestart = append(msg.EmergencyAreaIDListForRestart, *i)
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
