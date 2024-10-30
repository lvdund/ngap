package ies

import "github.com/lvdund/ngap/aper"

type TransportLayerAddress struct {
	Value aper.BitString `True,1,160`
}

func (ie *TransportLayerAddress) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, true)
	return
}
func (ie *TransportLayerAddress) Decode(r *aper.AperReader) (err error) {
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 160}, true); err != nil {
		return
	}
	ie.Value.Bytes = v
	ie.Value.NumBits = uint64(n)
	return
}
