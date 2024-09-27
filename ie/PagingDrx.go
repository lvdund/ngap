package ie

import "ngap/aper"

type PagingDrx struct {
	PagingDrx []byte `bitstring:"sizeLB:0,sizeUB:150"`
}

func (ie *PagingDrx) Decode(r *aper.AperReader) error {

	return nil
}

func (ie *PagingDrx) Encode(r *aper.AperWriter) (err error) {
	bitLength := 8 * len(ie.PagingDrx)
	if err = r.WriteBitString(ie.PagingDrx, uint(bitLength), &aper.Constraint{Lb: 0, Ub: 150}, false); err != nil {
		return err
	}
	return nil
}
