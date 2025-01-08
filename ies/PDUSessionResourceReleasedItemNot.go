package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceReleasedItemNot struct {
	PDUSessionID                             int64
	PDUSessionResourceNotifyReleasedTransfer []byte
	// IEExtensions *PDUSessionResourceReleasedItemNotExtIEs `optional`
}

func (ie *PDUSessionResourceReleasedItemNot) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		return
	}
	tmp_PDUSessionResourceNotifyReleasedTransfer := NewOCTETSTRING(ie.PDUSessionResourceNotifyReleasedTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_PDUSessionResourceNotifyReleasedTransfer.Encode(w); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceReleasedItemNot) Decode(r *aper.AperReader) (err error) {
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
	tmp_PDUSessionResourceNotifyReleasedTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_PDUSessionResourceNotifyReleasedTransfer.Decode(r); err != nil {
		return
	}
	ie.PDUSessionResourceNotifyReleasedTransfer = tmp_PDUSessionResourceNotifyReleasedTransfer.Value
	return
}
