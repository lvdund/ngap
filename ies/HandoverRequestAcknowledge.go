package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverRequestAcknowledge struct {
	AMFUENGAPID                              int64
	RANUENGAPID                              int64
	PDUSessionResourceAdmittedList           []PDUSessionResourceAdmittedItem
	PDUSessionResourceFailedToSetupListHOAck *[]PDUSessionResourceFailedToSetupItemHOAck
	TargetToSourceTransparentContainer       []byte
	CriticalityDiagnostics                   *CriticalityDiagnostics
}

func (msg *HandoverRequestAcknowledge) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_HandoverResourceAllocation, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverRequestAcknowledge) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.AMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	tmp_PDUSessionResourceAdmittedList := Sequence[*PDUSessionResourceAdmittedItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceAdmittedList {
		tmp_PDUSessionResourceAdmittedList.Value = append(tmp_PDUSessionResourceAdmittedList.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceAdmittedList},
		Criticality: Criticality{Value: Criticality_PresentIgnore},
		Value:       &tmp_PDUSessionResourceAdmittedList,
	})
	if msg.PDUSessionResourceFailedToSetupListHOAck != nil {
		tmp_PDUSessionResourceFailedToSetupListHOAck := Sequence[*PDUSessionResourceFailedToSetupItemHOAck]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		for _, i := range *msg.PDUSessionResourceFailedToSetupListHOAck {
			tmp_PDUSessionResourceFailedToSetupListHOAck.Value = append(tmp_PDUSessionResourceFailedToSetupListHOAck.Value, &i)
		}
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToSetupListHOAck},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       &tmp_PDUSessionResourceFailedToSetupListHOAck,
		})
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_TargetToSourceTransparentContainer},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &OCTETSTRING{
			c:     aper.Constraint{Lb: 0, Ub: 0},
			ext:   false,
			Value: msg.TargetToSourceTransparentContainer,
		}})
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics,
		})
	}
	return
}
func (msg *HandoverRequestAcknowledge) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *HandoverRequestAcknowledge) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_AMFUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.AMFUENGAPID = int64(tmp.Value)
	case ProtocolIEID_RANUENGAPID:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANUENGAPID = int64(tmp.Value)
	case ProtocolIEID_PDUSessionResourceAdmittedList:
		tmp := Sequence[*PDUSessionResourceAdmittedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceAdmittedList = []PDUSessionResourceAdmittedItem{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceAdmittedList = append(msg.PDUSessionResourceAdmittedList, *i)
		}
	case ProtocolIEID_PDUSessionResourceFailedToSetupListHOAck:
		tmp := Sequence[*PDUSessionResourceFailedToSetupItemHOAck]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceFailedToSetupListHOAck = &[]PDUSessionResourceFailedToSetupItemHOAck{}
		for _, i := range tmp.Value {
			*msg.PDUSessionResourceFailedToSetupListHOAck = append(*msg.PDUSessionResourceFailedToSetupListHOAck, *i)
		}
	case ProtocolIEID_TargetToSourceTransparentContainer:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TargetToSourceTransparentContainer = tmp.Value
	case ProtocolIEID_CriticalityDiagnostics:
		var tmp CriticalityDiagnostics
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.CriticalityDiagnostics = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
