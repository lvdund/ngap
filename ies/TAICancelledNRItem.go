package ies

import "github.com/lvdund/ngap/aper"

type TAICancelledNRItem struct {
	TAI                   *TAI                   `True,`
	CancelledCellsInTAINR *CancelledCellsInTAINR `False,`
	// IEExtensions TAICancelledNRItemExtIEs `False,OPTIONAL`
}

func (ie *TAICancelledNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.TAI != nil {
		if err = ie.TAI.Encode(w); err != nil {
			return
		}
	}
	if ie.CancelledCellsInTAINR != nil {
		if err = ie.CancelledCellsInTAINR.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TAICancelledNRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.TAI = new(TAI)
	ie.CancelledCellsInTAINR = new(CancelledCellsInTAINR)
	if err = ie.TAI.Decode(r); err != nil {
		return
	}
	if err = ie.CancelledCellsInTAINR.Decode(r); err != nil {
		return
	}
	return
}
