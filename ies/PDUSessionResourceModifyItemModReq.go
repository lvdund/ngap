package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyItemModReq struct {
	PDUSessionID                            int64
	NASPDU                                  []byte
	PDUSessionResourceModifyRequestTransfer []byte
	// IEExtensions *PDUSessionResourceModifyItemModReqExtIEs `optional`
}

func (ie *PDUSessionResourceModifyItemModReq) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NASPDU != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.PDUSessionResourceModifyRequestTransfer != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		return
	}
	if ie.NASPDU != nil {
		tmp_NASPDU := NewOCTETSTRING(ie.NASPDU, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_NASPDU.Encode(w); err != nil {
			return
		}
	}
	if ie.PDUSessionResourceModifyRequestTransfer != nil {
		tmp_PDUSessionResourceModifyRequestTransfer := NewOCTETSTRING(ie.PDUSessionResourceModifyRequestTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_PDUSessionResourceModifyRequestTransfer.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceModifyItemModReq) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
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
	if aper.IsBitSet(optionals, 1) {
		tmp_NASPDU := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_NASPDU.Decode(r); err != nil {
			return
		}
		ie.NASPDU = tmp_NASPDU.Value
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_PDUSessionResourceModifyRequestTransfer := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_PDUSessionResourceModifyRequestTransfer.Decode(r); err != nil {
			return
		}
		ie.PDUSessionResourceModifyRequestTransfer = tmp_PDUSessionResourceModifyRequestTransfer.Value
	}
	return
}
