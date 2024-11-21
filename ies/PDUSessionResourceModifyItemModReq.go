package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceModifyItemModReq struct {
	PDUSessionID                            *PDUSessionID     `False,`
	NASPDU                                  *NASPDU           `False,OPTIONAL`
	PDUSessionResourceModifyRequestTransfer *aper.OctetString `False,`
	// IEExtensions PDUSessionResourceModifyItemModReqExtIEs `False,OPTIONAL`
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
	if ie.PDUSessionID != nil {
		if err = ie.PDUSessionID.Encode(w); err != nil {
			return
		}
	}
	if ie.NASPDU != nil {
		if err = ie.NASPDU.Encode(w); err != nil {
			return
		}
	}
	if ie.PDUSessionResourceModifyRequestTransfer != nil {
		if err = w.WriteOctetString(*ie.PDUSessionResourceModifyRequestTransfer, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
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
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.PDUSessionID = new(PDUSessionID)
	ie.NASPDU = new(NASPDU)
	var o []byte
	if err = ie.PDUSessionID.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.NASPDU.Decode(r); err != nil {
			return
		}
	}
	if o, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return
	} else {
		ie.PDUSessionResourceModifyRequestTransfer = (*aper.OctetString)(&o)
	}
	return
}
