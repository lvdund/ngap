package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	BroadcastCancelledAreaListPresentNothing uint64 = iota
	BroadcastCancelledAreaListPresentCellidcancelledeutra
	BroadcastCancelledAreaListPresentTaicancelledeutra
	BroadcastCancelledAreaListPresentEmergencyareaidcancelledeutra
	BroadcastCancelledAreaListPresentCellidcancellednr
	BroadcastCancelledAreaListPresentTaicancellednr
	BroadcastCancelledAreaListPresentEmergencyareaidcancellednr
	BroadcastCancelledAreaListPresentChoiceExtensions
)

type BroadcastCancelledAreaList struct {
	Choice                        uint64
	CellIDCancelledEUTRA          []CellIDCancelledEUTRAItem
	TAICancelledEUTRA             []TAICancelledEUTRAItem
	EmergencyAreaIDCancelledEUTRA []EmergencyAreaIDCancelledEUTRAItem
	CellIDCancelledNR             []CellIDCancelledNRItem
	TAICancelledNR                []TAICancelledNRItem
	EmergencyAreaIDCancelledNR    []EmergencyAreaIDCancelledNRItem
	// ChoiceExtensions *BroadcastCancelledAreaListExtIEs
}

func (ie *BroadcastCancelledAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCancelledAreaListPresentCellidcancelledeutra:
		tmp := Sequence[*CellIDCancelledEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		for _, i := range ie.CellIDCancelledEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentTaicancelledeutra:
		tmp := Sequence[*TAICancelledEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range ie.TAICancelledEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentEmergencyareaidcancelledeutra:
		tmp := Sequence[*EmergencyAreaIDCancelledEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range ie.EmergencyAreaIDCancelledEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentCellidcancellednr:
		tmp := Sequence[*CellIDCancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		for _, i := range ie.CellIDCancelledNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentTaicancellednr:
		tmp := Sequence[*TAICancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range ie.TAICancelledNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentEmergencyareaidcancellednr:
		tmp := Sequence[*EmergencyAreaIDCancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range ie.EmergencyAreaIDCancelledNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *BroadcastCancelledAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCancelledAreaListPresentCellidcancelledeutra:
		tmp := NewSequence[*CellIDCancelledEUTRAItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning}, false)
		fn := func() *CellIDCancelledEUTRAItem {
			return new(CellIDCancelledEUTRAItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read CellIDCancelledEUTRA", err)
			return
		}
		for _, i := range tmp.Value {
			ie.CellIDCancelledEUTRA = append(ie.CellIDCancelledEUTRA, *i)
		}
	case BroadcastCancelledAreaListPresentTaicancelledeutra:
		tmp := NewSequence[*TAICancelledEUTRAItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning}, false)
		fn := func() *TAICancelledEUTRAItem {
			return new(TAICancelledEUTRAItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read TAICancelledEUTRA", err)
			return
		}
		for _, i := range tmp.Value {
			ie.TAICancelledEUTRA = append(ie.TAICancelledEUTRA, *i)
		}
	case BroadcastCancelledAreaListPresentEmergencyareaidcancelledeutra:
		tmp := NewSequence[*EmergencyAreaIDCancelledEUTRAItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID}, false)
		fn := func() *EmergencyAreaIDCancelledEUTRAItem {
			return new(EmergencyAreaIDCancelledEUTRAItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EmergencyAreaIDCancelledEUTRA", err)
			return
		}
		for _, i := range tmp.Value {
			ie.EmergencyAreaIDCancelledEUTRA = append(ie.EmergencyAreaIDCancelledEUTRA, *i)
		}
	case BroadcastCancelledAreaListPresentCellidcancellednr:
		tmp := NewSequence[*CellIDCancelledNRItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning}, false)
		fn := func() *CellIDCancelledNRItem {
			return new(CellIDCancelledNRItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read CellIDCancelledNR", err)
			return
		}
		for _, i := range tmp.Value {
			ie.CellIDCancelledNR = append(ie.CellIDCancelledNR, *i)
		}
	case BroadcastCancelledAreaListPresentTaicancellednr:
		tmp := NewSequence[*TAICancelledNRItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning}, false)
		fn := func() *TAICancelledNRItem {
			return new(TAICancelledNRItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read TAICancelledNR", err)
			return
		}
		for _, i := range tmp.Value {
			ie.TAICancelledNR = append(ie.TAICancelledNR, *i)
		}
	case BroadcastCancelledAreaListPresentEmergencyareaidcancellednr:
		tmp := NewSequence[*EmergencyAreaIDCancelledNRItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID}, false)
		fn := func() *EmergencyAreaIDCancelledNRItem {
			return new(EmergencyAreaIDCancelledNRItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EmergencyAreaIDCancelledNR", err)
			return
		}
		for _, i := range tmp.Value {
			ie.EmergencyAreaIDCancelledNR = append(ie.EmergencyAreaIDCancelledNR, *i)
		}
	}
	return
}
