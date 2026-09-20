package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          []byte                         `lb:3,ub:3,madatory`
	CancelledCellsInEAIEUTRA []CancelledCellsInEAIEUTRAItem `lb:1,ub:maxnoofCellinEAI,madatory`
	// IEExtensions *EmergencyAreaIDCancelledEUTRAItemExtIEs `optional`
}

func (ie *EmergencyAreaIDCancelledEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EmergencyAreaID := NewOCTETSTRING(ie.EmergencyAreaID, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_EmergencyAreaID.Encode(w); err != nil {
		err = fmt.Errorf("Encode EmergencyAreaID: %w", err)
		return
	}
	if len(ie.CancelledCellsInEAIEUTRA) > 0 {
		tmp := Sequence[*CancelledCellsInEAIEUTRAItem]{
			Value: []*CancelledCellsInEAIEUTRAItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
			ext:   false,
		}
		for _, i := range ie.CancelledCellsInEAIEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode CancelledCellsInEAIEUTRA: %w", err)
			return
		}
	} else {
		if err != nil {
			err = fmt.Errorf("CancelledCellsInEAIEUTRA is nil: %w", err)
		}
		return
	}
	return
}
func (ie *EmergencyAreaIDCancelledEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_EmergencyAreaID := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_EmergencyAreaID.Decode(r); err != nil {
		err = fmt.Errorf("Read EmergencyAreaID: %w", err)
		return
	}
	ie.EmergencyAreaID = tmp_EmergencyAreaID.Value
	tmp_CancelledCellsInEAIEUTRA := Sequence[*CancelledCellsInEAIEUTRAItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
		ext: false,
	}
	fn := func() *CancelledCellsInEAIEUTRAItem { return new(CancelledCellsInEAIEUTRAItem) }
	if err = tmp_CancelledCellsInEAIEUTRA.Decode(r, fn); err != nil {
		err = fmt.Errorf("Read CancelledCellsInEAIEUTRA: %w", err)
		return
	}
	ie.CancelledCellsInEAIEUTRA = []CancelledCellsInEAIEUTRAItem{}
	for _, i := range tmp_CancelledCellsInEAIEUTRA.Value {
		ie.CancelledCellsInEAIEUTRA = append(ie.CancelledCellsInEAIEUTRA, *i)
	}
	return
}
