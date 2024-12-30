package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDCancelledNRItem struct {
	EmergencyAreaID       []byte
	CancelledCellsInEAINR []CancelledCellsInEAINRItem
	// IEExtensions  *EmergencyAreaIDCancelledNRItemExtIEs
}

func (ie *EmergencyAreaIDCancelledNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EmergencyAreaID := NewOCTETSTRING(ie.EmergencyAreaID, aper.Constraint{Lb: 3, Ub: 3}, true)
	if err = tmp_EmergencyAreaID.Encode(w); err != nil {
		return
	}
	if len(ie.CancelledCellsInEAINR) > 0 {
		tmp := Sequence[*CancelledCellsInEAINRItem]{
			Value: []*CancelledCellsInEAINRItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
			ext:   false,
		}
		for _, i := range ie.CancelledCellsInEAINR {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EmergencyAreaIDCancelledNRItem) Decode(r *aper.AperReader) (err error) {
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
	tmp_CancelledCellsInEAINR := Sequence[*CancelledCellsInEAINRItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
		ext: false,
	}
	if err = tmp_CancelledCellsInEAINR.Decode(r); err != nil {
		return
	}
	ie.CancelledCellsInEAINR = []CancelledCellsInEAINRItem{}
	for _, i := range tmp_CancelledCellsInEAINR.Value {
		ie.CancelledCellsInEAINR = append(ie.CancelledCellsInEAINR, *i)
	}
	return
}
