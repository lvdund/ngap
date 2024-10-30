package ies

import "github.com/lvdund/ngap/aper"

type TAIBroadcastNRItem struct {
	TAI                   *TAI                   `True,`
	CompletedCellsInTAINR *CompletedCellsInTAINR `False,`
	// IEExtensions TAIBroadcastNRItemExtIEs `False,OPTIONAL`
}

func (ie *TAIBroadcastNRItem) Encode(w *aper.AperWriter) (err error) {
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
	if ie.CompletedCellsInTAINR != nil {
		if err = ie.CompletedCellsInTAINR.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TAIBroadcastNRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.TAI = new(TAI)
	ie.CompletedCellsInTAINR = new(CompletedCellsInTAINR)
	if err = ie.TAI.Decode(r); err != nil {
		return
	}
	if err = ie.CompletedCellsInTAINR.Decode(r); err != nil {
		return
	}
	return
}
