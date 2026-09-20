package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type RecommendedCellItem struct {
	NGRANCGI         NGRANCGI `madatory`
	TimeStayedInCell *int64   `lb:0,ub:4095,optional`
	// IEExtensions *RecommendedCellItemExtIEs `optional`
}

func (ie *RecommendedCellItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TimeStayedInCell != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.NGRANCGI.Encode(w); err != nil {
		err = fmt.Errorf("Encode NGRANCGI: %w", err)
		return
	}
	if ie.TimeStayedInCell != nil {
		tmp_TimeStayedInCell := NewINTEGER(*ie.TimeStayedInCell, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp_TimeStayedInCell.Encode(w); err != nil {
			err = fmt.Errorf("Encode TimeStayedInCell: %w", err)
			return
		}
	}
	return
}
func (ie *RecommendedCellItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.NGRANCGI.Decode(r); err != nil {
		err = fmt.Errorf("Read NGRANCGI: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_TimeStayedInCell := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: false,
		}
		if err = tmp_TimeStayedInCell.Decode(r); err != nil {
			err = fmt.Errorf("Read TimeStayedInCell: %w", err)
			return
		}
		ie.TimeStayedInCell = (*int64)(&tmp_TimeStayedInCell.Value)
	}
	return
}
