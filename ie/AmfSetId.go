package ie

import (
	"ngap/aper"
)

type AmfSetId struct {
	AmfSetId aper.BitString `bitstring:"sizeLB:10,sizeUB:10"`
}

func (a *AmfSetId) Decode(r aper.AperReader) error {

	return nil
}

func (a *AmfSetId) Encode(r aper.AperWriter) (err error) {
	if err = r.WriteBitString(a.AmfSetId, &aper.Constrain{Lb: 10, Ub: 10}, true); err != nil {
		return err
	}
	return nil
}
