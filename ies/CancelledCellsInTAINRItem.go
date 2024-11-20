package ies

import "github.com/lvdund/ngap/aper"

type CancelledCellsInTAINRItem struct {
	NRCGI              *NRCGI              `True,`
	NumberOfBroadcasts *NumberOfBroadcasts `False,`
	// IEExtensions CancelledCellsInTAINRItemExtIEs `False,OPTIONAL`
}

func (ie *CancelledCellsInTAINRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.NRCGI != nil {
		if err = ie.NRCGI.Encode(w); err != nil {
			return
		}
	}
	if ie.NumberOfBroadcasts != nil {
		if err = ie.NumberOfBroadcasts.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CancelledCellsInTAINRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.NRCGI = new(NRCGI)
	ie.NumberOfBroadcasts = new(NumberOfBroadcasts)
	if err = ie.NRCGI.Decode(r); err != nil {
		return
	}
	if err = ie.NumberOfBroadcasts.Decode(r); err != nil {
		return
	}
	return
}
