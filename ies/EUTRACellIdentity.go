package ies

import "github.com/lvdund/ngap/aper"

type EUTRACellIdentity struct {
	Value aper.BitString `False,28,28`
}

func (ie *EUTRACellIdentity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 28, Ub: 28}, false)
	return
}
func (ie *EUTRACellIdentity) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
