package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceItemCxtRelCpl struct {
	PDUSessionID *PDUSessionID `False,`
	// IEExtensions PDUSessionResourceItemCxtRelCplExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceItemCxtRelCpl) Encode(w *aper.AperWriter) (err error) {
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
	return
}
func (ie *PDUSessionResourceItemCxtRelCpl) Decode(r *aper.AperReader) (err error) {
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
