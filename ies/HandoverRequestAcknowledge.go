package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type HandoverRequestAcknowledge struct {
	AMFUENGAPID                              *AMFUENGAPID                              `,ignore,mandatory`
	RANUENGAPID                              *RANUENGAPID                              `,ignore,mandatory`
	PDUSessionResourceAdmittedList           *PDUSessionResourceAdmittedList           `,ignore,mandatory`
	PDUSessionResourceFailedToSetupListHOAck *PDUSessionResourceFailedToSetupListHOAck `,ignore,optional`
	TargetToSourceTransparentContainer       *TargetToSourceTransparentContainer       `,reject,mandatory`
	CriticalityDiagnostics                   *CriticalityDiagnostics                   `,ignore,optional`
}

func (msg *HandoverRequestAcknowledge) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduSuccessfulOutcome, ProcedureCode_HandoverResourceAllocation, Criticality_PresentReject, msg.toIes())
}
func (msg *HandoverRequestAcknowledge) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RANUENGAPID})
	}
	if msg.PDUSessionResourceAdmittedList != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceAdmittedList},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceAdmittedList})
	}
	if msg.PDUSessionResourceFailedToSetupListHOAck != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceFailedToSetupListHOAck},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.PDUSessionResourceFailedToSetupListHOAck})
	}
	if msg.TargetToSourceTransparentContainer != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_TargetToSourceTransparentContainer},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.TargetToSourceTransparentContainer})
	}
	if msg.CriticalityDiagnostics != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_CriticalityDiagnostics},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.CriticalityDiagnostics})
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
	case ProtocolIEID_PDUSessionResourceAdmittedList:
		var tmp PDUSessionResourceAdmittedList
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceAdmittedList = &tmp
	case ProtocolIEID_PDUSessionResourceFailedToSetupListHOAck:
		var tmp PDUSessionResourceFailedToSetupListHOAck
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceFailedToSetupListHOAck = &tmp
	case ProtocolIEID_TargetToSourceTransparentContainer:
		var tmp TargetToSourceTransparentContainer
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.TargetToSourceTransparentContainer = &tmp
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
