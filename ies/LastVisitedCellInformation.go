package ies

import "github.com/lvdund/ngap/aper"

const (
	LastVisitedCellInformationPresentNothing uint64 = iota
	LastVisitedCellInformationPresentNgrancell
	LastVisitedCellInformationPresentEutrancell
	LastVisitedCellInformationPresentUtrancell
	LastVisitedCellInformationPresentGerancell
	LastVisitedCellInformationPresentChoiceExtensions
)

type LastVisitedCellInformation struct {
	Choice     uint64
	NGRANCell  *LastVisitedNGRANCellInformation
	EUTRANCell []byte
	UTRANCell  []byte
	GERANCell  []byte
	// ChoiceExtensions *LastVisitedCellInformationExtIEs
}

func (ie *LastVisitedCellInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return
	}
	switch ie.Choice {
	case LastVisitedCellInformationPresentNgrancell:
		err = ie.NGRANCell.Encode(w)
	case LastVisitedCellInformationPresentEutrancell:
		tmp := NewOCTETSTRING(ie.EUTRANCell, aper.Constraint{Lb: 0, Ub: 0}, false)
		err = tmp.Encode(w)
	case LastVisitedCellInformationPresentUtrancell:
		tmp := NewOCTETSTRING(ie.UTRANCell, aper.Constraint{Lb: 0, Ub: 0}, false)
		err = tmp.Encode(w)
	case LastVisitedCellInformationPresentGerancell:
		tmp := NewOCTETSTRING(ie.GERANCell, aper.Constraint{Lb: 0, Ub: 0}, false)
		err = tmp.Encode(w)
	}
	return
}
func (ie *LastVisitedCellInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return
	}
	switch ie.Choice {
	case LastVisitedCellInformationPresentNgrancell:
		var tmp LastVisitedNGRANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NGRANCell = &tmp
	case LastVisitedCellInformationPresentEutrancell:
		tmp := NewOCTETSTRING(nil, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRANCell = tmp.Value
	case LastVisitedCellInformationPresentUtrancell:
		tmp := NewOCTETSTRING(nil, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UTRANCell = tmp.Value
	case LastVisitedCellInformationPresentGerancell:
		tmp := NewOCTETSTRING(nil, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GERANCell = tmp.Value
	}
	return
}
