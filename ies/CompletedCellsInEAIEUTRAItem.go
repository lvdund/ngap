package ies

import "github.com/lvdund/ngap/aper"

type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI *EUTRACGI `True,`
	// IEExtensions CompletedCellsInEAIEUTRAItemExtIEs `False,OPTIONAL`
}

func (ie *CompletedCellsInEAIEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
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
func (ie *CompletedCellsInEAIEUTRAItem) Decode(r *aper.AperReader) (err error) {
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
