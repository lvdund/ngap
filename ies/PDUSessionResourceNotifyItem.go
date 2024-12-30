package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceNotifyItem struct {
	PDUSessionID                     int64
	PDUSessionResourceNotifyTransfer []byte
	// IEExtensions  *PDUSessionResourceNotifyItemExtIEs
}

func (ie *PDUSessionResourceNotifyItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, true)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		return
	}
	tmp_PDUSessionResourceNotifyTransfer := NewOCTETSTRING(ie.PDUSessionResourceNotifyTransfer, aper.Constraint{Lb: 0, Ub: 0}, true)
	if err = tmp_PDUSessionResourceNotifyTransfer.Encode(w); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceNotifyItem) Decode(r *aper.AperReader) (err error) {
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
	tmp_PDUSessionResourceNotifyTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_PDUSessionResourceNotifyTransfer.Decode(r); err != nil {
		return
	}
	ie.PDUSessionResourceNotifyTransfer = tmp_PDUSessionResourceNotifyTransfer.Value
	return
}
