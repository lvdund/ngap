package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToSetupItemCxtRes struct {
	PDUSessionID                                *PDUSessionID     `False,`
	PDUSessionResourceSetupUnsuccessfulTransfer *aper.OctetString `False,`
	// IEExtensions PDUSessionResourceFailedToSetupItemCxtResExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceFailedToSetupItemCxtRes) Encode(w *aper.AperWriter) (err error) {
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
	if ie.PDUSessionResourceSetupUnsuccessfulTransfer != nil {
		if err = w.WriteOctetString(*ie.PDUSessionResourceSetupUnsuccessfulTransfer, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceFailedToSetupItemCxtRes) Decode(r *aper.AperReader) (err error) {
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
		ie.PDUSessionResourceSetupUnsuccessfulTransfer = (*aper.OctetString)(&o)
	}
	return
}
