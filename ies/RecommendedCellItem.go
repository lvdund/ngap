package ies

import "github.com/lvdund/ngap/aper"

type RecommendedCellItem struct {
	NGRANCGI         NGRANCGI
	TimeStayedInCell *int64 `optional`
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
		return
	}
	if ie.TimeStayedInCell != nil {
		tmp_TimeStayedInCell := NewINTEGER(*ie.TimeStayedInCell, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp_TimeStayedInCell.Encode(w); err != nil {
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
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_TimeStayedInCell := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: false,
		}
		if err = tmp_TimeStayedInCell.Decode(r); err != nil {
			return
		}
		ie.TimeStayedInCell = (*int64)(&tmp_TimeStayedInCell.Value)
	}
	return
}
