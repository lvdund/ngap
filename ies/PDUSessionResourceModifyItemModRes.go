package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyItemModRes struct {
	PDUSessionID                             int64
	PDUSessionResourceModifyResponseTransfer []byte
	// IEExtensions *PDUSessionResourceModifyItemModResExtIEs `optional`
}

func (ie *PDUSessionResourceModifyItemModRes) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		return
	}
	tmp_PDUSessionResourceModifyResponseTransfer := NewOCTETSTRING(ie.PDUSessionResourceModifyResponseTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_PDUSessionResourceModifyResponseTransfer.Encode(w); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceModifyItemModRes) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PDUSessionID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_PDUSessionID.Decode(r); err != nil {
		return
	}
	ie.PDUSessionID = int64(tmp_PDUSessionID.Value)
	tmp_PDUSessionResourceModifyResponseTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_PDUSessionResourceModifyResponseTransfer.Decode(r); err != nil {
		return
	}
	ie.PDUSessionResourceModifyResponseTransfer = tmp_PDUSessionResourceModifyResponseTransfer.Value
	return
}
