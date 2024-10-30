package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupItemCxtRes struct {
	PDUSessionID                            *PDUSessionID     `False,`
	PDUSessionResourceSetupResponseTransfer *aper.OctetString `False,`
	// IEExtensions PDUSessionResourceSetupItemCxtResExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceSetupItemCxtRes) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PDUSessionID != nil {
		if err = ie.PDUSessionID.Encode(w); err != nil {
			return
		}
	}
	if ie.PDUSessionResourceSetupResponseTransfer != nil {
		if err = w.WriteOctetString(*ie.PDUSessionResourceSetupResponseTransfer, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceSetupItemCxtRes) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PDUSessionID = new(PDUSessionID)
	var o []byte
	if err = ie.PDUSessionID.Decode(r); err != nil {
		return
	}
	if o, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return
	} else {
		ie.PDUSessionResourceSetupResponseTransfer = (*aper.OctetString)(&o)
	}
	return
}
