package ies

import "github.com/lvdund/ngap/aper"

const (
	LastVisitedCellInformationPresentNothing uint64 = iota /* No components present */
	LastVisitedCellInformationPresentNGRANCell
	LastVisitedCellInformationPresentEUTRANCell
	LastVisitedCellInformationPresentUTRANCell
	LastVisitedCellInformationPresentGERANCell
	LastVisitedCellInformationPresentChoiceExtensions
)

type LastVisitedCellInformation struct {
	Choice     uint64
	NGRANCell  *LastVisitedNGRANCellInformation  `False,,,`
	EUTRANCell *LastVisitedEUTRANCellInformation `False,,,`
	UTRANCell  *LastVisitedUTRANCellInformation  `False,,,`
	GERANCell  *LastVisitedGERANCellInformation  `False,,,`
	// ChoiceExtensions *LastVisitedCellInformationExtIEs `False,,,`
}

func (ie *LastVisitedCellInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return
	}
	switch ie.Choice {
	case LastVisitedCellInformationPresentNGRANCell:
		err = ie.NGRANCell.Encode(w)
	case LastVisitedCellInformationPresentEUTRANCell:
		err = ie.EUTRANCell.Encode(w)
	case LastVisitedCellInformationPresentUTRANCell:
		err = ie.UTRANCell.Encode(w)
	case LastVisitedCellInformationPresentGERANCell:
		err = ie.GERANCell.Encode(w)
	}
	return
}
func (ie *LastVisitedCellInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return
	}
	switch ie.Choice {
	case LastVisitedCellInformationPresentNGRANCell:
		var tmp LastVisitedNGRANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NGRANCell = &tmp
	case LastVisitedCellInformationPresentEUTRANCell:
		var tmp LastVisitedEUTRANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRANCell = &tmp
	case LastVisitedCellInformationPresentUTRANCell:
		var tmp LastVisitedUTRANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UTRANCell = &tmp
	case LastVisitedCellInformationPresentGERANCell:
		var tmp LastVisitedGERANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GERANCell = &tmp
	}
	return
}
