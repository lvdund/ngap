package ie

import "ngap/aper"

type PagingDrx struct {
	PagingDrx []byte `bitstring:"sizeLB:0,sizeUB:150"`
}

func (ie *PagingDrx) Decode(r aper.AperReader) error {

	return nil
}

func (ie *PagingDrx) Encode(r aper.AperWriter) (err error) {

	return nil
}
