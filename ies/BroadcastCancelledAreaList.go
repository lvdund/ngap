package ies

import "github.com/lvdund/ngap/aper"

type BroadcastCancelledAreaList struct {
	Choice                        uint64
	CellIDCancelledEUTRA          *CellIDCancelledEUTRA          `False,,,`
	TAICancelledEUTRA             *TAICancelledEUTRA             `False,,,`
	EmergencyAreaIDCancelledEUTRA *EmergencyAreaIDCancelledEUTRA `False,,,`
	CellIDCancelledNR             *CellIDCancelledNR             `False,,,`
	TAICancelledNR                *TAICancelledNR                `False,,,`
	EmergencyAreaIDCancelledNR    *EmergencyAreaIDCancelledNR    `False,,,`
	// ChoiceExtensions *BroadcastCancelledAreaListExtIEs `False,,,`
}

func (ie *BroadcastCancelledAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.CellIDCancelledEUTRA.Encode(w)
	case 2:
		err = ie.TAICancelledEUTRA.Encode(w)
	case 3:
		err = ie.EmergencyAreaIDCancelledEUTRA.Encode(w)
	case 4:
		err = ie.CellIDCancelledNR.Encode(w)
	case 5:
		err = ie.TAICancelledNR.Encode(w)
	case 6:
		err = ie.EmergencyAreaIDCancelledNR.Encode(w)
	}
	return
}
func (ie *BroadcastCancelledAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp CellIDCancelledEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDCancelledEUTRA = &tmp
	case 2:
		var tmp TAICancelledEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAICancelledEUTRA = &tmp
	case 3:
		var tmp EmergencyAreaIDCancelledEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDCancelledEUTRA = &tmp
	case 4:
		var tmp CellIDCancelledNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDCancelledNR = &tmp
	case 5:
		var tmp TAICancelledNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAICancelledNR = &tmp
	case 6:
		var tmp EmergencyAreaIDCancelledNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDCancelledNR = &tmp
	}
	return
}
