package ies

import "github.com/lvdund/ngap/aper"

type TAIBroadcastEUTRAItem struct {
	TAI                      *TAI                      `True,`
	CompletedCellsInTAIEUTRA *CompletedCellsInTAIEUTRA `False,`
	// IEExtensions TAIBroadcastEUTRAItemExtIEs `False,OPTIONAL`
}

func (ie *TAIBroadcastEUTRAItem) Encode(w *aper.AperWriter) (err error) {
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
	if ie.CompletedCellsInTAIEUTRA != nil {
		if err = ie.CompletedCellsInTAIEUTRA.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TAIBroadcastEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.TAI = new(TAI)
	ie.CompletedCellsInTAIEUTRA = new(CompletedCellsInTAIEUTRA)
	if err = ie.TAI.Decode(r); err != nil {
		return
	}
	if err = ie.CompletedCellsInTAIEUTRA.Decode(r); err != nil {
		return
	}
	return
}
