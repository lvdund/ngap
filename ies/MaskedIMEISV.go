package ies

import "github.com/lvdund/ngap/aper"

type MaskedIMEISV struct {
	Value aper.BitString `False,64,64`
}

func (ie *MaskedIMEISV) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 64, Ub: 64}, false)
	return
}
func (ie *MaskedIMEISV) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
