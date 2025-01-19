package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceModifyItemModReq struct {
	PDUSessionID                            int64  `lb:0,ub:255,madatory`
	NASPDU                                  []byte `lb:0,ub:0,optional`
	PDUSessionResourceModifyRequestTransfer []byte `lb:0,ub:0,madatory`
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
	w.WriteBits(optionals, 2)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		err = utils.WrapError("Encode PDUSessionID", err)
		return
	}
	if ie.NASPDU != nil {
		tmp_NASPDU := NewOCTETSTRING(ie.NASPDU, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_NASPDU.Encode(w); err != nil {
			err = utils.WrapError("Encode NASPDU", err)
			return
		}
	}
	tmp_PDUSessionResourceModifyRequestTransfer := NewOCTETSTRING(ie.PDUSessionResourceModifyRequestTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_PDUSessionResourceModifyRequestTransfer.Encode(w); err != nil {
		err = utils.WrapError("Encode PDUSessionResourceModifyRequestTransfer", err)
		return
	}
	return
}
func (ie *PDUSessionResourceModifyItemModReq) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_PDUSessionID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_PDUSessionID.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionID", err)
		return
	}
	ie.PDUSessionID = int64(tmp_PDUSessionID.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_NASPDU := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_NASPDU.Decode(r); err != nil {
			err = utils.WrapError("Read NASPDU", err)
			return
		}
		ie.NASPDU = tmp_NASPDU.Value
	}
	tmp_PDUSessionResourceModifyRequestTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_PDUSessionResourceModifyRequestTransfer.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionResourceModifyRequestTransfer", err)
		return
	}
	ie.PDUSessionResourceModifyRequestTransfer = tmp_PDUSessionResourceModifyRequestTransfer.Value
	return
}
