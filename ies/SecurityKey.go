package ies

import "github.com/lvdund/ngap/aper"

type SecurityKey struct {
	Value aper.BitString `False,256,256`
}

func (ie *SecurityKey) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 256, Ub: 256}, false)
	return
}
func (ie *SecurityKey) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
