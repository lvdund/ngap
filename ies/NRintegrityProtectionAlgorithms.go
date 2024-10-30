package ies

import "github.com/lvdund/ngap/aper"

type NRintegrityProtectionAlgorithms struct {
	Value aper.BitString `True,16,16`
}

func (ie *NRintegrityProtectionAlgorithms) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, true)
	return
}
func (ie *NRintegrityProtectionAlgorithms) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, true); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
