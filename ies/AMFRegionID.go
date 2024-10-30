package ies

import "github.com/lvdund/ngap/aper"

type AMFRegionID struct {
	Value aper.BitString `False,8,8`
}

func (ie *AMFRegionID) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false)
	return
}
func (ie *AMFRegionID) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
