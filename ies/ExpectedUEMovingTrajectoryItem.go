package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ExpectedUEMovingTrajectoryItem struct {
	NGRANCGI         NGRANCGI `madatory`
	TimeStayedInCell *int64   `lb:0,ub:4095,optional`
	// IEExtensions *ExpectedUEMovingTrajectoryItemExtIEs `optional`
}

func (ie *ExpectedUEMovingTrajectoryItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TimeStayedInCell != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.NGRANCGI.Encode(w); err != nil {
		err = utils.WrapError("Encode NGRANCGI", err)
		return
	}
	if ie.TimeStayedInCell != nil {
		tmp_TimeStayedInCell := NewINTEGER(*ie.TimeStayedInCell, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp_TimeStayedInCell.Encode(w); err != nil {
			err = utils.WrapError("Encode TimeStayedInCell", err)
			return
		}
	}
	return
}
func (ie *ExpectedUEMovingTrajectoryItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.NGRANCGI.Decode(r); err != nil {
		err = utils.WrapError("Read NGRANCGI", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_TimeStayedInCell := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: false,
		}
		if err = tmp_TimeStayedInCell.Decode(r); err != nil {
			err = utils.WrapError("Read TimeStayedInCell", err)
			return
		}
		ie.TimeStayedInCell = (*int64)(&tmp_TimeStayedInCell.Value)
	}
	return
}
