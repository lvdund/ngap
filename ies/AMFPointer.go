package ies

import "github.com/lvdund/ngap/aper"

type AMFPointer struct {
	Value aper.BitString `False,6,6`
}

func (ie *AMFPointer) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 6, Ub: 6}, false)
	return
}
func (ie *AMFPointer) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
