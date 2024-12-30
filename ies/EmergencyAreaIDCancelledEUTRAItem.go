package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          []byte
	CancelledCellsInEAIEUTRA []CancelledCellsInEAIEUTRAItem
	// IEExtensions  *EmergencyAreaIDCancelledEUTRAItemExtIEs
}

func (ie *EmergencyAreaIDCancelledEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EmergencyAreaID := NewOCTETSTRING(ie.EmergencyAreaID, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_EmergencyAreaID.Encode(w); err != nil {
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
			return
		}
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
		return
	}
	ie.EmergencyAreaID = tmp_EmergencyAreaID.Value
	tmp_CancelledCellsInEAIEUTRA := Sequence[*CancelledCellsInEAIEUTRAItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
		ext: false,
	}
	if err = tmp_CancelledCellsInEAIEUTRA.Decode(r); err != nil {
		return
	}
	ie.CancelledCellsInEAIEUTRA = []CancelledCellsInEAIEUTRAItem{}
	for _, i := range tmp_CancelledCellsInEAIEUTRA.Value {
		ie.CancelledCellsInEAIEUTRA = append(ie.CancelledCellsInEAIEUTRA, *i)
	}
	return
}
