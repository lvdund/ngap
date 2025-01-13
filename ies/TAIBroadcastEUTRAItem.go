package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TAIBroadcastEUTRAItem struct {
	TAI                      TAI                            `madatory`
	CompletedCellsInTAIEUTRA []CompletedCellsInTAIEUTRAItem `lb:1,ub:maxnoofCellinTAI,madatory`
	// IEExtensions *TAIBroadcastEUTRAItemExtIEs `optional`
}

func (ie *TAIBroadcastEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TAI.Encode(w); err != nil {
		err = utils.WrapError("Encode TAI", err)
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
			err = utils.WrapError("Encode CompletedCellsInTAIEUTRA", err)
			return
		}
	} else {
		err = utils.WrapError("CompletedCellsInTAIEUTRA is nil", err)
		return
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
		err = utils.WrapError("Read TAI", err)
		return
	}
	tmp_CompletedCellsInTAIEUTRA := Sequence[*CompletedCellsInTAIEUTRAItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
		ext: false,
	}
	fn := func() *CompletedCellsInTAIEUTRAItem { return new(CompletedCellsInTAIEUTRAItem) }
	if err = tmp_CompletedCellsInTAIEUTRA.Decode(r, fn); err != nil {
		err = utils.WrapError("Read CompletedCellsInTAIEUTRA", err)
		return
	}
	ie.CompletedCellsInTAIEUTRA = []CompletedCellsInTAIEUTRAItem{}
	for _, i := range tmp_CompletedCellsInTAIEUTRA.Value {
		ie.CompletedCellsInTAIEUTRA = append(ie.CompletedCellsInTAIEUTRA, *i)
	}
	return
}
