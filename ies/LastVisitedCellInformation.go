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
	EUTRANCell *OCTETSTRING
	UTRANCell  *OCTETSTRING
	GERANCell  *OCTETSTRING
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
		err = ie.EUTRANCell.Encode(w)
	case LastVisitedCellInformationPresentUtrancell:
		err = ie.UTRANCell.Encode(w)
	case LastVisitedCellInformationPresentGerancell:
		err = ie.GERANCell.Encode(w)
	}
	return
}
func (ie *LastVisitedCellInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
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
		var tmp OCTETSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRANCell = &tmp
	case LastVisitedCellInformationPresentUtrancell:
		var tmp OCTETSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UTRANCell = &tmp
	case LastVisitedCellInformationPresentGerancell:
		var tmp OCTETSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GERANCell = &tmp
	}
	return
}
