package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EmergencyAreaIDBroadcastNRItem struct {
	EmergencyAreaID       []byte
	CompletedCellsInEAINR []CompletedCellsInEAINRItem
	// IEExtensions *EmergencyAreaIDBroadcastNRItemExtIEs `optional`
}

func (ie *EmergencyAreaIDBroadcastNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EmergencyAreaID := NewOCTETSTRING(ie.EmergencyAreaID, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_EmergencyAreaID.Encode(w); err != nil {
		err = utils.WrapError("Read EmergencyAreaID", err)
		return
	}
	if len(ie.CompletedCellsInEAINR) > 0 {
		tmp := Sequence[*CompletedCellsInEAINRItem]{
			Value: []*CompletedCellsInEAINRItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
			ext:   false,
		}
		for _, i := range ie.CompletedCellsInEAINR {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Read CompletedCellsInEAINR", err)
			return
		}
	}
	return
}
func (ie *EmergencyAreaIDBroadcastNRItem) Decode(r *aper.AperReader) (err error) {
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
		err = utils.WrapError("Read EmergencyAreaID", err)
		return
	}
	ie.EmergencyAreaID = tmp_EmergencyAreaID.Value
	tmp_CompletedCellsInEAINR := Sequence[*CompletedCellsInEAINRItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
		ext: false,
	}
	fn := func() *CompletedCellsInEAINRItem { return new(CompletedCellsInEAINRItem) }
	if err = tmp_CompletedCellsInEAINR.Decode(r, fn); err != nil {
		err = utils.WrapError("Read CompletedCellsInEAINR", err)
		return
	}
	ie.CompletedCellsInEAINR = []CompletedCellsInEAINRItem{}
	for _, i := range tmp_CompletedCellsInEAINR.Value {
		ie.CompletedCellsInEAINR = append(ie.CompletedCellsInEAINR, *i)
	}
	return
}
