package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	BroadcastCompletedAreaListPresentNothing uint64 = iota
	BroadcastCompletedAreaListPresentCellidbroadcasteutra
	BroadcastCompletedAreaListPresentTaibroadcasteutra
	BroadcastCompletedAreaListPresentEmergencyareaidbroadcasteutra
	BroadcastCompletedAreaListPresentCellidbroadcastnr
	BroadcastCompletedAreaListPresentTaibroadcastnr
	BroadcastCompletedAreaListPresentEmergencyareaidbroadcastnr
	BroadcastCompletedAreaListPresentChoiceExtensions
)

type BroadcastCompletedAreaList struct {
	Choice                        uint64
	CellIDBroadcastEUTRA          []CellIDBroadcastEUTRAItem
	TAIBroadcastEUTRA             []TAIBroadcastEUTRAItem
	EmergencyAreaIDBroadcastEUTRA []EmergencyAreaIDBroadcastEUTRAItem
	CellIDBroadcastNR             []CellIDBroadcastNRItem
	TAIBroadcastNR                []TAIBroadcastNRItem
	EmergencyAreaIDBroadcastNR    []EmergencyAreaIDBroadcastNRItem
	// ChoiceExtensions *BroadcastCompletedAreaListExtIEs
}

func (ie *BroadcastCompletedAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCompletedAreaListPresentCellidbroadcasteutra:
		tmp := Sequence[*CellIDBroadcastEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		for _, i := range ie.CellIDBroadcastEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentTaibroadcasteutra:
		tmp := Sequence[*TAIBroadcastEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range ie.TAIBroadcastEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcasteutra:
		tmp := Sequence[*EmergencyAreaIDBroadcastEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range ie.EmergencyAreaIDBroadcastEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentCellidbroadcastnr:
		tmp := Sequence[*CellIDBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		for _, i := range ie.CellIDBroadcastNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentTaibroadcastnr:
		tmp := Sequence[*TAIBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range ie.TAIBroadcastNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcastnr:
		tmp := Sequence[*EmergencyAreaIDBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range ie.EmergencyAreaIDBroadcastNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *BroadcastCompletedAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCompletedAreaListPresentCellidbroadcasteutra:
		tmp := NewSequence[*CellIDBroadcastEUTRAItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning}, false)
		fn := func() *CellIDBroadcastEUTRAItem {
			return new(CellIDBroadcastEUTRAItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read CellIDBroadcastEUTRA", err)
			return
		}
		for _, i := range tmp.Value {
			ie.CellIDBroadcastEUTRA = append(ie.CellIDBroadcastEUTRA, *i)
		}
	case BroadcastCompletedAreaListPresentTaibroadcasteutra:
		tmp := NewSequence[*TAIBroadcastEUTRAItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning}, false)
		fn := func() *TAIBroadcastEUTRAItem {
			return new(TAIBroadcastEUTRAItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read TAIBroadcastEUTRA", err)
			return
		}
		for _, i := range tmp.Value {
			ie.TAIBroadcastEUTRA = append(ie.TAIBroadcastEUTRA, *i)
		}
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcasteutra:
		tmp := NewSequence[*EmergencyAreaIDBroadcastEUTRAItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID}, false)
		fn := func() *EmergencyAreaIDBroadcastEUTRAItem {
			return new(EmergencyAreaIDBroadcastEUTRAItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EmergencyAreaIDBroadcastEUTRA", err)
			return
		}
		for _, i := range tmp.Value {
			ie.EmergencyAreaIDBroadcastEUTRA = append(ie.EmergencyAreaIDBroadcastEUTRA, *i)
		}
	case BroadcastCompletedAreaListPresentCellidbroadcastnr:
		tmp := NewSequence[*CellIDBroadcastNRItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning}, false)
		fn := func() *CellIDBroadcastNRItem {
			return new(CellIDBroadcastNRItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read CellIDBroadcastNR", err)
			return
		}
		for _, i := range tmp.Value {
			ie.CellIDBroadcastNR = append(ie.CellIDBroadcastNR, *i)
		}
	case BroadcastCompletedAreaListPresentTaibroadcastnr:
		tmp := NewSequence[*TAIBroadcastNRItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning}, false)
		fn := func() *TAIBroadcastNRItem {
			return new(TAIBroadcastNRItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read TAIBroadcastNR", err)
			return
		}
		for _, i := range tmp.Value {
			ie.TAIBroadcastNR = append(ie.TAIBroadcastNR, *i)
		}
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcastnr:
		tmp := NewSequence[*EmergencyAreaIDBroadcastNRItem](nil, aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID}, false)
		fn := func() *EmergencyAreaIDBroadcastNRItem {
			return new(EmergencyAreaIDBroadcastNRItem)
		}
		if err = tmp.Decode(r, fn); err != nil {
			err = utils.WrapError("Read EmergencyAreaIDBroadcastNR", err)
			return
		}
		for _, i := range tmp.Value {
			ie.EmergencyAreaIDBroadcastNR = append(ie.EmergencyAreaIDBroadcastNR, *i)
		}
	}
	return
}
