package ies

import "github.com/lvdund/ngap/aper"

type LastVisitedCellInformation struct {
	Choice     uint64
	NGRANCell  *LastVisitedNGRANCellInformation  `False,,,`
	EUTRANCell *LastVisitedEUTRANCellInformation `False,,,`
	UTRANCell  *LastVisitedUTRANCellInformation  `False,,,`
	GERANCell  *LastVisitedGERANCellInformation  `False,,,`
	// ChoiceExtensions *LastVisitedCellInformationExtIEs `False,,,`
}

func (ie *LastVisitedCellInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.NGRANCell.Encode(w)
	case 2:
		err = ie.EUTRANCell.Encode(w)
	case 3:
		err = ie.UTRANCell.Encode(w)
	case 4:
		err = ie.GERANCell.Encode(w)
	}
	return
}
func (ie *LastVisitedCellInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp LastVisitedNGRANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NGRANCell = &tmp
	case 2:
		var tmp LastVisitedEUTRANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRANCell = &tmp
	case 3:
		var tmp LastVisitedUTRANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UTRANCell = &tmp
	case 4:
		var tmp LastVisitedGERANCellInformation
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GERANCell = &tmp
	}
	return
}
