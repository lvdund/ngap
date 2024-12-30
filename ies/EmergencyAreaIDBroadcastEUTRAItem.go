package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDBroadcastEUTRAItem struct {
	EmergencyAreaID          []byte
	CompletedCellsInEAIEUTRA []CompletedCellsInEAIEUTRAItem
	// IEExtensions  *EmergencyAreaIDBroadcastEUTRAItemExtIEs
}

func (ie *EmergencyAreaIDBroadcastEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EmergencyAreaID := NewOCTETSTRING(ie.EmergencyAreaID, aper.Constraint{Lb: 3, Ub: 3}, true)
	if err = tmp_EmergencyAreaID.Encode(w); err != nil {
		return
	}
	if len(ie.CompletedCellsInEAIEUTRA) > 0 {
		tmp := Sequence[*CompletedCellsInEAIEUTRAItem]{
			Value: []*CompletedCellsInEAIEUTRAItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
			ext:   false,
		}
		for _, i := range ie.CompletedCellsInEAIEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EmergencyAreaIDBroadcastEUTRAItem) Decode(r *aper.AperReader) (err error) {
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
	tmp_CompletedCellsInEAIEUTRA := Sequence[*CompletedCellsInEAIEUTRAItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
		ext: false,
	}
	if err = tmp_CompletedCellsInEAIEUTRA.Decode(r); err != nil {
		return
	}
	ie.CompletedCellsInEAIEUTRA = []CompletedCellsInEAIEUTRAItem{}
	for _, i := range tmp_CompletedCellsInEAIEUTRA.Value {
		ie.CompletedCellsInEAIEUTRA = append(ie.CompletedCellsInEAIEUTRA, *i)
	}
	return
}
