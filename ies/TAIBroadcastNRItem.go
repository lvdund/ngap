package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type TAIBroadcastNRItem struct {
	TAI                   TAI                         `madatory`
	CompletedCellsInTAINR []CompletedCellsInTAINRItem `lb:1,ub:maxnoofCellinTAI,madatory`
	// IEExtensions *TAIBroadcastNRItemExtIEs `optional`
}

func (ie *TAIBroadcastNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.TAI.Encode(w); err != nil {
		err = fmt.Errorf("Encode TAI: %w", err)
		return
	}
	if len(ie.CompletedCellsInTAINR) > 0 {
		tmp := Sequence[*CompletedCellsInTAINRItem]{
			Value: []*CompletedCellsInTAINRItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
			ext:   false,
		}
		for _, i := range ie.CompletedCellsInTAINR {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode CompletedCellsInTAINR: %w", err)
			return
		}
	} else {
		if err != nil {
			err = fmt.Errorf("CompletedCellsInTAINR is nil: %w", err)
		}
		return
	}
	return
}
func (ie *TAIBroadcastNRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.TAI.Decode(r); err != nil {
		err = fmt.Errorf("Read TAI: %w", err)
		return
	}
	tmp_CompletedCellsInTAINR := Sequence[*CompletedCellsInTAINRItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinTAI},
		ext: false,
	}
	fn := func() *CompletedCellsInTAINRItem { return new(CompletedCellsInTAINRItem) }
	if err = tmp_CompletedCellsInTAINR.Decode(r, fn); err != nil {
		err = fmt.Errorf("Read CompletedCellsInTAINR: %w", err)
		return
	}
	ie.CompletedCellsInTAINR = []CompletedCellsInTAINRItem{}
	for _, i := range tmp_CompletedCellsInTAINR.Value {
		ie.CompletedCellsInTAINR = append(ie.CompletedCellsInTAINR, *i)
	}
	return
}
