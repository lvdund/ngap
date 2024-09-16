package ie

import "ngap/aper"

type NbIotDefaultPagingDrx struct {
	NbIotDefaultPagingDrx aper.BitString `bitstring:"sizeLB:0,sizeUB:150"`
}

func (ie *NbIotDefaultPagingDrx) Decode(r aper.AperReader) error {
	return nil
}

func (ie *NbIotDefaultPagingDrx) Encode(r aper.AperWriter) (err error) {
	if err = r.WriteBitString(ie.NbIotDefaultPagingDrx.Bytes, uint(ie.NbIotDefaultPagingDrx.NumBits),&aper.Constrain{Lb: 0, Ub: 150}, false); err != nil {
		return err
	}
	return nil
}
