package ies

import "github.com/lvdund/ngap/aper"

type RATRestrictionInformation struct {
	Value aper.BitString `True,8,8`
}

func (ie *RATRestrictionInformation) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, true)
	return
}
func (ie *RATRestrictionInformation) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, true); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
