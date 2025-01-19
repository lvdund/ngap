package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceSetupItemCxtReq struct {
	PDUSessionID                           int64  `lb:0,ub:255,madatory`
	NASPDU                                 []byte `lb:0,ub:0,optional`
	SNSSAI                                 SNSSAI `madatory`
	PDUSessionResourceSetupRequestTransfer []byte `lb:0,ub:0,madatory`
	// IEExtensions *PDUSessionResourceSetupItemCxtReqExtIEs `optional`
}

func (ie *PDUSessionResourceSetupItemCxtReq) Encode(w *aper.AperWriter) (err error) {
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
	if err = ie.SNSSAI.Encode(w); err != nil {
		err = utils.WrapError("Encode SNSSAI", err)
		return
	}
	tmp_PDUSessionResourceSetupRequestTransfer := NewOCTETSTRING(ie.PDUSessionResourceSetupRequestTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_PDUSessionResourceSetupRequestTransfer.Encode(w); err != nil {
		err = utils.WrapError("Encode PDUSessionResourceSetupRequestTransfer", err)
		return
	}
	return
}
func (ie *PDUSessionResourceSetupItemCxtReq) Decode(r *aper.AperReader) (err error) {
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
	if err = ie.SNSSAI.Decode(r); err != nil {
		err = utils.WrapError("Read SNSSAI", err)
		return
	}
	tmp_PDUSessionResourceSetupRequestTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_PDUSessionResourceSetupRequestTransfer.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionResourceSetupRequestTransfer", err)
		return
	}
	ie.PDUSessionResourceSetupRequestTransfer = tmp_PDUSessionResourceSetupRequestTransfer.Value
	return
}
