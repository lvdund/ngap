package ies

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceModifyIndication struct {
	AMFUENGAPID                        int64
	RANUENGAPID                        int64
	PDUSessionResourceModifyListModInd []PDUSessionResourceModifyItemModInd
}

func (msg *PDUSessionResourceModifyIndication) Encode(w io.Writer) (err error) {
	return encodeMessage(w, NgapPduInitiatingMessage, ProcedureCode_PDUSessionResourceModifyIndication, Criticality_PresentReject, msg.toIes())
}
func (msg *PDUSessionResourceModifyIndication) toIes() (ies []NgapMessageIE) {
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
	tmp_PDUSessionResourceModifyListModInd := Sequence[*PDUSessionResourceModifyItemModInd]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
		ext: false,
	}
	for _, i := range msg.PDUSessionResourceModifyListModInd {
		tmp_PDUSessionResourceModifyListModInd.Value = append(tmp_PDUSessionResourceModifyListModInd.Value, &i)
	}
	ies = append(ies, NgapMessageIE{
		Id:          ProtocolIEID{Value: ProtocolIEID_PDUSessionResourceModifyListModInd},
		Criticality: Criticality{Value: Criticality_PresentReject},
		Value:       &tmp_PDUSessionResourceModifyListModInd,
	})
	return
}
func (msg *PDUSessionResourceModifyIndication) Decode(wire []byte) (err error, diagList []CriticalityDiagnostics) {
	r := aper.NewReader(bytes.NewReader(wire))
	r.ReadBool()
	var ies []NgapMessageIE
	if ies, err = aper.ReadSequenceOf[NgapMessageIE](msg.decodeIE, r, &aper.Constraint{Lb: 0, Ub: int64(aper.POW_16 - 1)}, false); err != nil {
		return
	}
	_ = ies
	return
}
func (msg *PDUSessionResourceModifyIndication) decodeIE(r *aper.AperReader) (msgIe *NgapMessageIE, err error) {
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
	case ProtocolIEID_PDUSessionResourceModifyListModInd:
		tmp := Sequence[*PDUSessionResourceModifyItemModInd]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp.Decode(ieR); err != nil {
			return
		}
		msg.PDUSessionResourceModifyListModInd = []PDUSessionResourceModifyItemModInd{}
		for _, i := range tmp.Value {
			msg.PDUSessionResourceModifyListModInd = append(msg.PDUSessionResourceModifyListModInd, *i)
		}
	default:
		err = fmt.Errorf("temporary error")
		return
	}
	return
}
