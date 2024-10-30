package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterestCellItem struct {
	NGRANCGI *NGRANCGI `False,`
	// IEExtensions AreaOfInterestCellItemExtIEs `False,OPTIONAL`
}

func (ie *AreaOfInterestCellItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.NGRANCGI != nil {
		if err = ie.NGRANCGI.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *AreaOfInterestCellItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.NGRANCGI = new(NGRANCGI)
	if err = ie.NGRANCGI.Decode(r); err != nil {
		return
	}
	return
}
