package ies

import "github.com/lvdund/ngap/aper"

type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           *EUTRACGI           `True,`
	NumberOfBroadcasts *NumberOfBroadcasts `False,`
	// IEExtensions CancelledCellsInEAIEUTRAItemExtIEs `False,OPTIONAL`
}

func (ie *CancelledCellsInEAIEUTRAItem) Encode(w *aper.AperWriter) (err error) {
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
	if ie.NumberOfBroadcasts != nil {
		if err = ie.NumberOfBroadcasts.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CancelledCellsInEAIEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.EUTRACGI = new(EUTRACGI)
	ie.NumberOfBroadcasts = new(NumberOfBroadcasts)
	if err = ie.EUTRACGI.Decode(r); err != nil {
		return
	}
	if err = ie.NumberOfBroadcasts.Decode(r); err != nil {
		return
	}
	return
}
