package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceSetupRequest struct {
	AMFUENGAPID                      int64
	RANUENGAPID                      int64
	RANPagingPriority                *int64
	NASPDU                           *[]byte
	PDUSessionResourceSetupListSUReq []PDUSessionResourceSetupItemSUReq
	UEAggregateMaximumBitRate        *UEAggregateMaximumBitRate
}

func (msg *PDUSessionResourceSetupRequest) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceSetup, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceSetupRequest) toIes() (ies []NgapMessageIE) {
	ies = []NgapMessageIE{}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_AMFUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext:   false,
			Value: aper.Integer(msg.AMFUENGAPID),
		}})
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_RANUENGAPID},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value: &INTEGER{
			c:     aper.Constraint{Lb: 0, Ub: 4294967295},
			ext:   false,
			Value: aper.Integer(msg.RANUENGAPID),
		}})
	if msg.RANPagingPriority != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_RANPagingPriority},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &INTEGER{
				c:     aper.Constraint{Lb: 1, Ub: 256},
				ext:   false,
				Value: aper.Integer(*msg.RANPagingPriority),
			}})
	}
	if msg.NASPDU != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_NASPDU},
			Criticality: Criticality{Value: Criticality_PresentReject},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: *msg.NASPDU,
			}})
	}
	tmp_PDUSessionResourceSetupListSUReq := Sequence[*PDUSessionResourceSetupItemSUReq]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceSetupListSUReq {
		tmp_PDUSessionResourceSetupListSUReq.Value = append(tmp_PDUSessionResourceSetupListSUReq.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceSetupListSUReq},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_PDUSessionResourceSetupListSUReq,
	})
	if msg.UEAggregateMaximumBitRate != nil {
		ies = append(ies, NgapMessageIE{
			Id:          ProtocolIEID{Value: ProtocolIEID_UEAggregateMaximumBitRate},
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value:       msg.UEAggregateMaximumBitRate,
		})
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
	case ProtocolIEID_RANPagingPriority:
		tmp := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.RANPagingPriority = int64(tmp.Value)
	case ProtocolIEID_NASPDU:
		tmp := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		*msg.NASPDU = tmp.Value
	case ProtocolIEID_PDUSessionResourceSetupListSUReq:
		tmp := Sequence[*PDUSessionResourceSetupItemSUReq]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceSetupListSUReq = []PDUSessionResourceSetupItemSUReq{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceSetupListSUReq = append(msg.PDUSessionResourceSetupListSUReq, *i)
		}
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
