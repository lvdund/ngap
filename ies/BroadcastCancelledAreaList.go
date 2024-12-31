package ies

import "github.com/lvdund/ngap/aper"

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
	CellIDCancelledEUTRA          *[]CellIDCancelledEUTRAItem
	TAICancelledEUTRA             *[]TAICancelledEUTRAItem
	EmergencyAreaIDCancelledEUTRA *[]EmergencyAreaIDCancelledEUTRAItem
	CellIDCancelledNR             *[]CellIDCancelledNRItem
	TAICancelledNR                *[]TAICancelledNRItem
	EmergencyAreaIDCancelledNR    *[]EmergencyAreaIDCancelledNRItem
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
		for _, i := range *ie.CellIDCancelledEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentTaicancelledeutra:
		tmp := Sequence[*TAICancelledEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range *ie.TAICancelledEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentEmergencyareaidcancelledeutra:
		tmp := Sequence[*EmergencyAreaIDCancelledEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range *ie.EmergencyAreaIDCancelledEUTRA {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentCellidcancellednr:
		tmp := Sequence[*CellIDCancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		for _, i := range *ie.CellIDCancelledNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentTaicancellednr:
		tmp := Sequence[*TAICancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		for _, i := range *ie.TAICancelledNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	case BroadcastCancelledAreaListPresentEmergencyareaidcancellednr:
		tmp := Sequence[*EmergencyAreaIDCancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		for _, i := range *ie.EmergencyAreaIDCancelledNR {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *BroadcastCancelledAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCancelledAreaListPresentCellidcancelledeutra:
		tmp := Sequence[*CellIDCancelledEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDCancelledEUTRA = &[]CellIDCancelledEUTRAItem{}
		for _, i := range tmp.Value {
			*ie.CellIDCancelledEUTRA = append(*ie.CellIDCancelledEUTRA, *i)
		}
	case BroadcastCancelledAreaListPresentTaicancelledeutra:
		tmp := Sequence[*TAICancelledEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAICancelledEUTRA = &[]TAICancelledEUTRAItem{}
		for _, i := range tmp.Value {
			*ie.TAICancelledEUTRA = append(*ie.TAICancelledEUTRA, *i)
		}
	case BroadcastCancelledAreaListPresentEmergencyareaidcancelledeutra:
		tmp := Sequence[*EmergencyAreaIDCancelledEUTRAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDCancelledEUTRA = &[]EmergencyAreaIDCancelledEUTRAItem{}
		for _, i := range tmp.Value {
			*ie.EmergencyAreaIDCancelledEUTRA = append(*ie.EmergencyAreaIDCancelledEUTRA, *i)
		}
	case BroadcastCancelledAreaListPresentCellidcancellednr:
		tmp := Sequence[*CellIDCancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellIDforWarning},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDCancelledNR = &[]CellIDCancelledNRItem{}
		for _, i := range tmp.Value {
			*ie.CellIDCancelledNR = append(*ie.CellIDCancelledNR, *i)
		}
	case BroadcastCancelledAreaListPresentTaicancellednr:
		tmp := Sequence[*TAICancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforWarning},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAICancelledNR = &[]TAICancelledNRItem{}
		for _, i := range tmp.Value {
			*ie.TAICancelledNR = append(*ie.TAICancelledNR, *i)
		}
	case BroadcastCancelledAreaListPresentEmergencyareaidcancellednr:
		tmp := Sequence[*EmergencyAreaIDCancelledNRItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEmergencyAreaID},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDCancelledNR = &[]EmergencyAreaIDCancelledNRItem{}
		for _, i := range tmp.Value {
			*ie.EmergencyAreaIDCancelledNR = append(*ie.EmergencyAreaIDCancelledNR, *i)
		}
	}
	return
}
