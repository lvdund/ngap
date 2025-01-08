package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type TAICancelledNRItem struct {
	TAI                   TAI
	CancelledCellsInTAINR []CancelledCellsInTAINRItem
	// IEExtensions *TAICancelledNRItemExtIEs `optional`
}

func (ie *TAICancelledNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TAI.Encode(w); err != nil {
		err = utils.WrapError("Read TAI", err)
		return
	}
	if len(ie.CancelledCellsInTAINR) > 0 {
		tmp := Sequence[*CancelledCellsInTAINRItem]{
			Value: []*CancelledCellsInTAINRItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
			ext:   false,
		}
		for _, i := range ie.CancelledCellsInTAINR {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Read CancelledCellsInTAINR", err)
			return
		}
	}
	return
}
func (ie *TAICancelledNRItem) Decode(r *aper.AperReader) (err error) {
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
	tmp_CancelledCellsInTAINR := Sequence[*CancelledCellsInTAINRItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
		ext: false,
	}
	fn := func() *CancelledCellsInTAINRItem { return new(CancelledCellsInTAINRItem) }
	if err = tmp_CancelledCellsInTAINR.Decode(r, fn); err != nil {
		err = utils.WrapError("Read CancelledCellsInTAINR", err)
		return
	}
	ie.CancelledCellsInTAINR = []CancelledCellsInTAINRItem{}
	for _, i := range tmp_CancelledCellsInTAINR.Value {
		ie.CancelledCellsInTAINR = append(ie.CancelledCellsInTAINR, *i)
	}
	return
}
