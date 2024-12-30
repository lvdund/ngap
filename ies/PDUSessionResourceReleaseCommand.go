package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceReleaseCommand struct {
	AMFUENGAPID                           int64
	RANUENGAPID                           int64
	RANPagingPriority                     *int64
	NASPDU                                *[]byte
	PDUSessionResourceToReleaseListRelCmd []PDUSessionResourceToReleaseItemRelCmd
}

func (msg *PDUSessionResourceReleaseCommand) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceRelease, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceReleaseCommand) toIes() (ies []NgapMessageIE) {
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
			Criticality: Criticality{Value: Criticality_PresentIgnore},
			Value: &OCTETSTRING{
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
				Value: *msg.NASPDU,
			}})
	}
	tmp_PDUSessionResourceToReleaseListRelCmd := Sequence[*PDUSessionResourceToReleaseItemRelCmd]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceToReleaseListRelCmd {
		tmp_PDUSessionResourceToReleaseListRelCmd.Value = append(tmp_PDUSessionResourceToReleaseListRelCmd.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceToReleaseListRelCmd},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_PDUSessionResourceToReleaseListRelCmd,
	})
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
	case ProtocolIEID_PDUSessionResourceToReleaseListRelCmd:
		tmp := Sequence[*PDUSessionResourceToReleaseItemRelCmd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceToReleaseListRelCmd = []PDUSessionResourceToReleaseItemRelCmd{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceToReleaseListRelCmd = append(msg.PDUSessionResourceToReleaseListRelCmd, *i)
		}
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
