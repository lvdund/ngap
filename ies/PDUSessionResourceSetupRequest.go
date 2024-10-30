package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceSetupRequest struct {
	AMFUENGAPID                      *AMFUENGAPID                      `,reject,mandatory`
	RANUENGAPID                      *RANUENGAPID                      `,reject,mandatory`
	RANPagingPriority                *RANPagingPriority                `,ignore,optional`
	NASPDU                           *NASPDU                           `,reject,optional`
	PDUSessionResourceSetupListSUReq *PDUSessionResourceSetupListSUReq `,reject,mandatory`
	UEAggregateMaximumBitRate        *UEAggregateMaximumBitRate        `,ignore,optional`
}

func (msg *PDUSessionResourceSetupRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceSetupRequest) toIes() (ies []NgapMessageIE) {
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
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.NASPDU})
	}
	if msg.PDUSessionResourceSetupListSUReq != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListSUReq},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value:       msg.PDUSessionResourceSetupListSUReq})
	}
	if msg.UEAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UEAggregateMaximumBitRate})
	}
	return
}
func (msg *PDUSessionResourceSetupRequest) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PDUSessionResourceSetupRequest) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceSetupListSUReq:
		var tmp PDUSessionResourceSetupListSUReq
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSetupListSUReq = &tmp
	case ProtocolIEID_UEAggregateMaximumBitRate:
		var tmp UEAggregateMaximumBitRate
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.UEAggregateMaximumBitRate = &tmp
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
