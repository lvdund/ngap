package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceItemHORqd struct {
	PDUSessionID             int64
	HandoverRequiredTransfer []byte
	// IEExtensions  *PDUSessionResourceItemHORqdExtIEs
}

func (ie *PDUSessionResourceItemHORqd) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, true)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		return
	}
	tmp_HandoverRequiredTransfer := NewOCTETSTRING(ie.HandoverRequiredTransfer, aper.Constraint{Lb: 0, Ub: 0}, true)
	if err = tmp_HandoverRequiredTransfer.Encode(w); err != nil {
		return
	}
	return
}
func (ie *PDUSessionResourceItemHORqd) Decode(r *aper.AperReader) (err error) {
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
	tmp_HandoverRequiredTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_HandoverRequiredTransfer.Decode(r); err != nil {
		return
	}
	ie.HandoverRequiredTransfer = tmp_HandoverRequiredTransfer.Value
	return
}
