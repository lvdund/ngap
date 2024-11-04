package ies

import "github.com/lvdund/ngap/aper"

const (
	BroadcastCancelledAreaListPresentNothing uint64 = iota /* No components present */
	BroadcastCancelledAreaListPresentCellIDCancelledEUTRA
	BroadcastCancelledAreaListPresentTAICancelledEUTRA
	BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledEUTRA
	BroadcastCancelledAreaListPresentCellIDCancelledNR
	BroadcastCancelledAreaListPresentTAICancelledNR
	BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledNR
	BroadcastCancelledAreaListPresentChoiceExtensions
)


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
	if err = w.WriteChoice(ie.Choice, 7, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCancelledAreaListPresentCellIDCancelledEUTRA:
		err = ie.CellIDCancelledEUTRA.Encode(w)
	case BroadcastCancelledAreaListPresentTAICancelledEUTRA:
		err = ie.TAICancelledEUTRA.Encode(w)
	case BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledEUTRA:
		err = ie.EmergencyAreaIDCancelledEUTRA.Encode(w)
	case BroadcastCancelledAreaListPresentCellIDCancelledNR:
		err = ie.CellIDCancelledNR.Encode(w)
	case BroadcastCancelledAreaListPresentTAICancelledNR:
		err = ie.TAICancelledNR.Encode(w)
	case BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledNR:
		err = ie.EmergencyAreaIDCancelledNR.Encode(w)
	}
	return
}
func (ie *BroadcastCancelledAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCancelledAreaListPresentCellIDCancelledEUTRA:
		var tmp CellIDCancelledEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDCancelledEUTRA = &tmp
	case BroadcastCancelledAreaListPresentTAICancelledEUTRA:
		var tmp TAICancelledEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAICancelledEUTRA = &tmp
	case BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledEUTRA:
		var tmp EmergencyAreaIDCancelledEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDCancelledEUTRA = &tmp
	case BroadcastCancelledAreaListPresentCellIDCancelledNR:
		var tmp CellIDCancelledNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDCancelledNR = &tmp
	case BroadcastCancelledAreaListPresentTAICancelledNR:
		var tmp TAICancelledNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAICancelledNR = &tmp
	case BroadcastCancelledAreaListPresentEmergencyAreaIDCancelledNR:
		var tmp EmergencyAreaIDCancelledNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDCancelledNR = &tmp
	}
	return
}
