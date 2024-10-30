package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceItemCxtRelReq struct {
	PDUSessionID *PDUSessionID `False,`
	// IEExtensions PDUSessionResourceItemCxtRelReqExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceItemCxtRelReq) Encode(w *aper.AperWriter) (err error) {
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
	return
}
func (ie *PDUSessionResourceItemCxtRelReq) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PDUSessionID = new(PDUSessionID)
	if err = ie.PDUSessionID.Decode(r); err != nil {
		return
	}
	return
}
