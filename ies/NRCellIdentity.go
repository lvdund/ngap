package ies

import "github.com/lvdund/ngap/aper"

type NRCellIdentity struct {
	Value aper.BitString `False,36,36`
}

func (ie *NRCellIdentity) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 36, Ub: 36}, false)
	return
}
func (ie *NRCellIdentity) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 36, Ub: 36}, false); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
