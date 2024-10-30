package ies

import "github.com/lvdund/ngap/aper"

type AMFSetID struct {
	Value aper.BitString `False,10,10`
}

func (ie *AMFSetID) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 10, Ub: 10}, false)
	return
}
func (ie *AMFSetID) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
