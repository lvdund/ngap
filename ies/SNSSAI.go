package ies

import "github.com/lvdund/ngap/aper"

type SNSSAI struct {
	SST *SST `False,`
	SD  *SD  `False,OPTIONAL`
	// IEExtensions SNSSAIExtIEs `False,OPTIONAL`
}

func (ie *SNSSAI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SD != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.SST != nil {
		if err = ie.SST.Encode(w); err != nil {
			return
		}
	}
	if ie.SD != nil {
		if err = ie.SD.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *SNSSAI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.SST = new(SST)
	ie.SD = new(SD)
	if err = ie.SST.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.SD.Decode(r); err != nil {
			return
		}
	}
	return
}
