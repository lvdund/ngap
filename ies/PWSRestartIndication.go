package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PWSRestartIndication struct {
	CellIDListForRestart          CellIDListForRestart
	GlobalRANNodeID               GlobalRANNodeID
	TAIListForRestart             []TAI
	EmergencyAreaIDListForRestart *[]EmergencyAreaID
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
		for _, i := range *msg.EmergencyAreaIDListForRestart {
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
func (msg *PWSRestartIndication) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PWSRestartIndication) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_CellIDListForRestart:
		var tmp CellIDListForRestart
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CellIDListForRestart = tmp
	case ProtocolIEID_GlobalRANNodeID:
		var tmp GlobalRANNodeID
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.GlobalRANNodeID = tmp
	case ProtocolIEID_TAIListForRestart:
		tmp := Sequence[*TAI]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforRestart},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
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
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.EmergencyAreaIDListForRestart = &[]EmergencyAreaID{}
		for _, i := range tmp.Value {
			*msg.EmergencyAreaIDListForRestart = append(*msg.EmergencyAreaIDListForRestart, *i)
		}
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
