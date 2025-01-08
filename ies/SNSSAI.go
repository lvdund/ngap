package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SNSSAI struct {
	SST []byte
	SD  []byte
	// IEExtensions *SNSSAIExtIEs `optional`
}

func (ie *SNSSAI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SD != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_SST := NewOCTETSTRING(ie.SST, aper.Constraint{Lb: 1, Ub: 1}, false)
	if err = tmp_SST.Encode(w); err != nil {
		err = utils.WrapError("Read SST", err)
		return
	}
	if ie.SD != nil {
		tmp_SD := NewOCTETSTRING(ie.SD, aper.Constraint{Lb: 3, Ub: 3}, false)
		if err = tmp_SD.Encode(w); err != nil {
			err = utils.WrapError("Read SD", err)
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
	tmp_SST := OCTETSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 1},
		ext: false,
	}
	if err = tmp_SST.Decode(r); err != nil {
		err = utils.WrapError("Read SST", err)
		return
	}
	ie.SST = tmp_SST.Value
	if aper.IsBitSet(optionals, 1) {
		tmp_SD := OCTETSTRING{
			c:   aper.Constraint{Lb: 3, Ub: 3},
			ext: false,
		}
		if err = tmp_SD.Decode(r); err != nil {
			err = utils.WrapError("Read SD", err)
			return
		}
		ie.SD = tmp_SD.Value
	}
	return
}
