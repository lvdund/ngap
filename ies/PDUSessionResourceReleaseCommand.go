package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceReleaseCommand struct {
	AMFUENGAPID                           *AMFUENGAPID                           `,reject,mandatory`
	RANUENGAPID                           *RANUENGAPID                           `,reject,mandatory`
	RANPagingPriority                     *RANPagingPriority                     `,ignore,optional`
	NASPDU                                *NASPDU                                `,ignore,optional`
	PDUSessionResourceToReleaseListRelCmd *PDUSessionResourceToReleaseListRelCmd `,reject,mandatory`
}

func (msg *PDUSessionResourceReleaseCommand) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceRelease, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceReleaseCommand) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	if msg.AMFUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.AMFUENGAPID})
	}
	if msg.RANUENGAPID != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.RANUENGAPID})
	}
	if msg.RANPagingPriority != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANPagingPriority},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.RANPagingPriority})
	}
	if msg.NASPDU != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.NASPDU})
	}
	if msg.PDUSessionResourceToReleaseListRelCmd != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToReleaseListRelCmd},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionResourceToReleaseListRelCmd})
	}
	return
}
func (msg *PDUSessionResourceReleaseCommand) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PDUSessionResourceReleaseCommand) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_RANPagingPriority:
		var tmp RANPagingPriority
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.RANPagingPriority = &tmp
	case ProtocolIEID_NASPDU:
		var tmp NASPDU
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.NASPDU = &tmp
	case ProtocolIEID_PDUSessionResourceToReleaseListRelCmd:
		var tmp PDUSessionResourceToReleaseListRelCmd
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceToReleaseListRelCmd = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
