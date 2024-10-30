package ies

import "github.com/lvdund/ngap/aper"

type TAICancelledEUTRAItem struct {
	TAI                      *TAI                      `True,`
	CancelledCellsInTAIEUTRA *CancelledCellsInTAIEUTRA `False,`
	// IEExtensions TAICancelledEUTRAItemExtIEs `False,OPTIONAL`
}

func (ie *TAICancelledEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.TAI != nil {
		if err = ie.TAI.Encode(w); err != nil {
			return
		}
	}
	if ie.CancelledCellsInTAIEUTRA != nil {
		if err = ie.CancelledCellsInTAIEUTRA.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TAICancelledEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.TAI = new(TAI)
	ie.CancelledCellsInTAIEUTRA = new(CancelledCellsInTAIEUTRA)
	if err = ie.TAI.Decode(r); err != nil {
		return
	}
	if err = ie.CancelledCellsInTAIEUTRA.Decode(r); err != nil {
		return
	}
	return
}
