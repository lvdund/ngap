package ies

import "github.com/lvdund/ngap/aper"

type TAICancelledEUTRAItem struct {
	TAI                      TAI
	CancelledCellsInTAIEUTRA []CancelledCellsInTAIEUTRAItem
	// IEExtensions *TAICancelledEUTRAItemExtIEs `optional`
}

func (ie *TAICancelledEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TAI.Encode(w); err != nil {
		return
	}
	if len(ie.CancelledCellsInTAIEUTRA) > 0 {
		tmp := Sequence[*CancelledCellsInTAIEUTRAItem]{
			Value: []*CancelledCellsInTAIEUTRAItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
			ext:   false,
		}
		for _, i := range ie.CancelledCellsInTAIEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TAICancelledEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TAI.Decode(r); err != nil {
		return
	}
	tmp_CancelledCellsInTAIEUTRA := Sequence[*CancelledCellsInTAIEUTRAItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
		ext: false,
	}
	fn := func() *CancelledCellsInTAIEUTRAItem { return new(CancelledCellsInTAIEUTRAItem) }
	if err = tmp_CancelledCellsInTAIEUTRA.Decode(r, fn); err != nil {
		return
	}
	ie.CancelledCellsInTAIEUTRA = []CancelledCellsInTAIEUTRAItem{}
	for _, i := range tmp_CancelledCellsInTAIEUTRA.Value {
		ie.CancelledCellsInTAIEUTRA = append(ie.CancelledCellsInTAIEUTRA, *i)
	}
	return
}
