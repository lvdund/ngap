package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EmergencyAreaIDCancelledNRItem struct {
	EmergencyAreaID       []byte                      `lb:3,ub:3,madatory`
	CancelledCellsInEAINR []CancelledCellsInEAINRItem `lb:1,ub:maxnoofCellinEAI,madatory`
	// IEExtensions *EmergencyAreaIDCancelledNRItemExtIEs `optional`
}

func (ie *EmergencyAreaIDCancelledNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EmergencyAreaID := NewOCTETSTRING(ie.EmergencyAreaID, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_EmergencyAreaID.Encode(w); err != nil {
		err = utils.WrapError("Encode EmergencyAreaID", err)
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
			err = utils.WrapError("Encode CancelledCellsInEAINR", err)
			return
		}
	} else {
		err = utils.WrapError("CancelledCellsInEAINR is nil", err)
		return
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
		err = utils.WrapError("Read EmergencyAreaID", err)
		return
	}
	ie.EmergencyAreaID = tmp_EmergencyAreaID.Value
	tmp_CancelledCellsInEAINR := Sequence[*CancelledCellsInEAINRItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinEAI},
		ext: false,
	}
	fn := func() *CancelledCellsInEAINRItem { return new(CancelledCellsInEAINRItem) }
	if err = tmp_CancelledCellsInEAINR.Decode(r, fn); err != nil {
		err = utils.WrapError("Read CancelledCellsInEAINR", err)
		return
	}
	ie.CancelledCellsInEAINR = []CancelledCellsInEAINRItem{}
	for _, i := range tmp_CancelledCellsInEAINR.Value {
		ie.CancelledCellsInEAINR = append(ie.CancelledCellsInEAINR, *i)
	}
	return
}
