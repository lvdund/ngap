package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupItemSUReq struct {
	PDUSessionID                           *PDUSessionID     `False,`
	PDUSessionNASPDU                       *NASPDU           `False,OPTIONAL`
	SNSSAI                                 *SNSSAI           `True,`
	PDUSessionResourceSetupRequestTransfer *aper.OctetString `False,`
	// IEExtensions PDUSessionResourceSetupItemSUReqExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceSetupItemSUReq) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PDUSessionNASPDU != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.PDUSessionID != nil {
		if err = ie.PDUSessionID.Encode(w); err != nil {
			return
		}
	}
	if ie.PDUSessionNASPDU != nil {
		if err = ie.PDUSessionNASPDU.Encode(w); err != nil {
			return
		}
	}
	if ie.SNSSAI != nil {
		if err = ie.SNSSAI.Encode(w); err != nil {
			return
		}
	}
	if ie.PDUSessionResourceSetupRequestTransfer != nil {
		if err = w.WriteOctetString(*ie.PDUSessionResourceSetupRequestTransfer, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceSetupItemSUReq) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.PDUSessionID = new(PDUSessionID)
	ie.PDUSessionNASPDU = new(NASPDU)
	ie.SNSSAI = new(SNSSAI)
	var o []byte
	if err = ie.PDUSessionID.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.PDUSessionNASPDU.Decode(r); err != nil {
			return
		}
	}
	if err = ie.SNSSAI.Decode(r); err != nil {
		return
	}
	if o, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return
	} else {
		ie.PDUSessionResourceSetupRequestTransfer = (*aper.OctetString)(&o)
	}
	return
}
