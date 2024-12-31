package ies

import "github.com/lvdund/ngap/aper"

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
	CellIDBroadcastEUTRA          *[]CellIDBroadcastEUTRAItem
	TAIBroadcastEUTRA             *[]TAIBroadcastEUTRAItem
	EmergencyAreaIDBroadcastEUTRA *[]EmergencyAreaIDBroadcastEUTRAItem
	CellIDBroadcastNR             *[]CellIDBroadcastNRItem
	TAIBroadcastNR                *[]TAIBroadcastNRItem
	EmergencyAreaIDBroadcastNR    *[]EmergencyAreaIDBroadcastNRItem
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
		for _, i := range *ie.CellIDBroadcastEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentTaibroadcasteutra:
		tmp := Sequence[*TAIBroadcastEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range *ie.TAIBroadcastEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcasteutra:
		tmp := Sequence[*EmergencyAreaIDBroadcastEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range *ie.EmergencyAreaIDBroadcastEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentCellidbroadcastnr:
		tmp := Sequence[*CellIDBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		for _, i := range *ie.CellIDBroadcastNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentTaibroadcastnr:
		tmp := Sequence[*TAIBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range *ie.TAIBroadcastNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcastnr:
		tmp := Sequence[*EmergencyAreaIDBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range *ie.EmergencyAreaIDBroadcastNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *BroadcastCompletedAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCompletedAreaListPresentCellidbroadcasteutra:
		tmp := Sequence[*CellIDBroadcastEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDBroadcastEUTRA = &[]CellIDBroadcastEUTRAItem{}
		for _, i := range tmp.Value {
			*ie.CellIDBroadcastEUTRA = append(*ie.CellIDBroadcastEUTRA, *i)
		}
	case BroadcastCompletedAreaListPresentTaibroadcasteutra:
		tmp := Sequence[*TAIBroadcastEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIBroadcastEUTRA = &[]TAIBroadcastEUTRAItem{}
		for _, i := range tmp.Value {
			*ie.TAIBroadcastEUTRA = append(*ie.TAIBroadcastEUTRA, *i)
		}
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcasteutra:
		tmp := Sequence[*EmergencyAreaIDBroadcastEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDBroadcastEUTRA = &[]EmergencyAreaIDBroadcastEUTRAItem{}
		for _, i := range tmp.Value {
			*ie.EmergencyAreaIDBroadcastEUTRA = append(*ie.EmergencyAreaIDBroadcastEUTRA, *i)
		}
	case BroadcastCompletedAreaListPresentCellidbroadcastnr:
		tmp := Sequence[*CellIDBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDBroadcastNR = &[]CellIDBroadcastNRItem{}
		for _, i := range tmp.Value {
			*ie.CellIDBroadcastNR = append(*ie.CellIDBroadcastNR, *i)
		}
	case BroadcastCompletedAreaListPresentTaibroadcastnr:
		tmp := Sequence[*TAIBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIBroadcastNR = &[]TAIBroadcastNRItem{}
		for _, i := range tmp.Value {
			*ie.TAIBroadcastNR = append(*ie.TAIBroadcastNR, *i)
		}
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcastnr:
		tmp := Sequence[*EmergencyAreaIDBroadcastNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDBroadcastNR = &[]EmergencyAreaIDBroadcastNRItem{}
		for _, i := range tmp.Value {
			*ie.EmergencyAreaIDBroadcastNR = append(*ie.EmergencyAreaIDBroadcastNR, *i)
		}
	}
	return
}
