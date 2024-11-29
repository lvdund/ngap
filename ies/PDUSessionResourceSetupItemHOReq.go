package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupItemHOReq struct {
	PDUSessionID            *PDUSessionID     `False,`
	SNSSAI                  *SNSSAI           `True,`
	HandoverRequestTransfer *aper.OctetString `False,`
	// IEExtensions PDUSessionResourceSetupItemHOReqExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceSetupItemHOReq) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PDUSessionID != nil {
		if err = ie.PDUSessionID.Encode(w); err != nil {
			return
		}
	}
	if ie.SNSSAI != nil {
		if err = ie.SNSSAI.Encode(w); err != nil {
			return
		}
	}
	if ie.HandoverRequestTransfer != nil {
		if err = w.WriteOctetString(*ie.HandoverRequestTransfer, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceSetupItemHOReq) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PDUSessionID = new(PDUSessionID)
	ie.SNSSAI = new(SNSSAI)
	var o []byte
	if err = ie.PDUSessionID.Decode(r); err != nil {
		return
	}
	if err = ie.SNSSAI.Decode(r); err != nil {
		return
	}
	if o, err = r.ReadOctetString(nil, false); err != nil {
		return
	} else {
		ie.HandoverRequestTransfer = (*aper.OctetString)(&o)
	}
	return
}
