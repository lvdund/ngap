package ies

import "github.com/lvdund/ngap/aper"

type AllowedNSSAIItem struct {
	SNSSAI *SNSSAI `True,`
	// IEExtensions AllowedNSSAIItemExtIEs `False,OPTIONAL`
}

func (ie *AllowedNSSAIItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.SNSSAI != nil {
		if err = ie.SNSSAI.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *AllowedNSSAIItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.SNSSAI = new(SNSSAI)
	if err = ie.SNSSAI.Decode(r); err != nil {
		return
	}
	return
}
