package ies

import "github.com/lvdund/ngap/aper"

type CompletedCellsInTAINRItem struct {
	NRCGI *NRCGI `True,`
	// IEExtensions CompletedCellsInTAINRItemExtIEs `False,OPTIONAL`
}

func (ie *CompletedCellsInTAINRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.NRCGI != nil {
		if err = ie.NRCGI.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CompletedCellsInTAINRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.NRCGI = new(NRCGI)
	if err = ie.NRCGI.Decode(r); err != nil {
		return
	}
	return
}
