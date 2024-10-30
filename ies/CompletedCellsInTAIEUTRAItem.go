package ies

import "github.com/lvdund/ngap/aper"

type CompletedCellsInTAIEUTRAItem struct {
	EUTRACGI *EUTRACGI `True,`
	// IEExtensions CompletedCellsInTAIEUTRAItemExtIEs `False,OPTIONAL`
}

func (ie *CompletedCellsInTAIEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.EUTRACGI != nil {
		if err = ie.EUTRACGI.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CompletedCellsInTAIEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.EUTRACGI = new(EUTRACGI)
	if err = ie.EUTRACGI.Decode(r); err != nil {
		return
	}
	return
}
