package ies

import "github.com/lvdund/ngap/aper"

type CancelledCellsInTAINRItem struct {
	NRCGI              NRCGI
	NumberOfBroadcasts int64
	// IEExtensions *CancelledCellsInTAINRItemExtIEs `optional`
}

func (ie *CancelledCellsInTAINRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NRCGI.Encode(w); err != nil {
		return
	}
	tmp_NumberOfBroadcasts := NewINTEGER(ie.NumberOfBroadcasts, aper.Constraint{Lb: 0, Ub: 65535}, false)
	if err = tmp_NumberOfBroadcasts.Encode(w); err != nil {
		return
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
	if err = ie.NRCGI.Decode(r); err != nil {
		return
	}
	tmp_NumberOfBroadcasts := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 65535},
		ext: false,
	}
	if err = tmp_NumberOfBroadcasts.Decode(r); err != nil {
		return
	}
	ie.NumberOfBroadcasts = int64(tmp_NumberOfBroadcasts.Value)
	return
}
