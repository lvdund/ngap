package ies

import "github.com/lvdund/ngap/aper"

type TAIBroadcastEUTRAItem struct {
	TAI                      TAI
	CompletedCellsInTAIEUTRA []CompletedCellsInTAIEUTRAItem
	// IEExtensions  *TAIBroadcastEUTRAItemExtIEs
}

func (ie *TAIBroadcastEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TAI.Encode(w); err != nil {
		return
	}
	if len(ie.CompletedCellsInTAIEUTRA) > 0 {
		tmp := Sequence[*CompletedCellsInTAIEUTRAItem]{
			Value: []*CompletedCellsInTAIEUTRAItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
			ext:   false,
		}
		for _, i := range ie.CompletedCellsInTAIEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TAIBroadcastEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TAI.Decode(r); err != nil {
		return
	}
	tmp_CompletedCellsInTAIEUTRA := Sequence[*CompletedCellsInTAIEUTRAItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
		ext: false,
	}
	if err = tmp_CompletedCellsInTAIEUTRA.Decode(r); err != nil {
		return
	}
	ie.CompletedCellsInTAIEUTRA = []CompletedCellsInTAIEUTRAItem{}
	for _, i := range tmp_CompletedCellsInTAIEUTRA.Value {
		ie.CompletedCellsInTAIEUTRA = append(ie.CompletedCellsInTAIEUTRA, *i)
	}
	return
}
