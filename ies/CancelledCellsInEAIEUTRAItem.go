package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CancelledCellsInEAIEUTRAItem struct {
	EUTRACGI           EUTRACGI
	NumberOfBroadcasts int64
	// IEExtensions *CancelledCellsInEAIEUTRAItemExtIEs `optional`
}

func (ie *CancelledCellsInEAIEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.EUTRACGI.Encode(w); err != nil {
		err = utils.WrapError("Read EUTRACGI", err)
		return
	}
	tmp_NumberOfBroadcasts := NewINTEGER(ie.NumberOfBroadcasts, aper.Constraint{Lb: 0, Ub: 65535}, false)
	if err = tmp_NumberOfBroadcasts.Encode(w); err != nil {
		err = utils.WrapError("Read NumberOfBroadcasts", err)
		return
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
	if err = ie.EUTRACGI.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRACGI", err)
		return
	}
	tmp_NumberOfBroadcasts := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 65535},
		ext: false,
	}
	if err = tmp_NumberOfBroadcasts.Decode(r); err != nil {
		err = utils.WrapError("Read NumberOfBroadcasts", err)
		return
	}
	ie.NumberOfBroadcasts = int64(tmp_NumberOfBroadcasts.Value)
	return
}
