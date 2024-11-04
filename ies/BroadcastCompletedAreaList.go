package ies

import "github.com/lvdund/ngap/aper"

const (
	BroadcastCompletedAreaListPresentNothing uint64 = iota /* No components present */
	BroadcastCompletedAreaListPresentCellIDBroadcastEUTRA
	BroadcastCompletedAreaListPresentTAIBroadcastEUTRA
	BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastEUTRA
	BroadcastCompletedAreaListPresentCellIDBroadcastNR
	BroadcastCompletedAreaListPresentTAIBroadcastNR
	BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastNR
	BroadcastCompletedAreaListPresentChoiceExtensions
)

type BroadcastCompletedAreaList struct {
	Choice                        uint64
	CellIDBroadcastEUTRA          *CellIDBroadcastEUTRA          `False,,,`
	TAIBroadcastEUTRA             *TAIBroadcastEUTRA             `False,,,`
	EmergencyAreaIDBroadcastEUTRA *EmergencyAreaIDBroadcastEUTRA `False,,,`
	CellIDBroadcastNR             *CellIDBroadcastNR             `False,,,`
	TAIBroadcastNR                *TAIBroadcastNR                `False,,,`
	EmergencyAreaIDBroadcastNR    *EmergencyAreaIDBroadcastNR    `False,,,`
	// ChoiceExtensions *BroadcastCompletedAreaListExtIEs `False,,,`
}

func (ie *BroadcastCompletedAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 7, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCompletedAreaListPresentCellIDBroadcastEUTRA:
		err = ie.CellIDBroadcastEUTRA.Encode(w)
	case BroadcastCompletedAreaListPresentTAIBroadcastEUTRA:
		err = ie.TAIBroadcastEUTRA.Encode(w)
	case BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastEUTRA:
		err = ie.EmergencyAreaIDBroadcastEUTRA.Encode(w)
	case BroadcastCompletedAreaListPresentCellIDBroadcastNR:
		err = ie.CellIDBroadcastNR.Encode(w)
	case BroadcastCompletedAreaListPresentTAIBroadcastNR:
		err = ie.TAIBroadcastNR.Encode(w)
	case BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastNR:
		err = ie.EmergencyAreaIDBroadcastNR.Encode(w)
	}
	return
}
func (ie *BroadcastCompletedAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCompletedAreaListPresentCellIDBroadcastEUTRA:
		var tmp CellIDBroadcastEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDBroadcastEUTRA = &tmp
	case BroadcastCompletedAreaListPresentTAIBroadcastEUTRA:
		var tmp TAIBroadcastEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIBroadcastEUTRA = &tmp
	case BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastEUTRA:
		var tmp EmergencyAreaIDBroadcastEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDBroadcastEUTRA = &tmp
	case BroadcastCompletedAreaListPresentCellIDBroadcastNR:
		var tmp CellIDBroadcastNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDBroadcastNR = &tmp
	case BroadcastCompletedAreaListPresentTAIBroadcastNR:
		var tmp TAIBroadcastNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIBroadcastNR = &tmp
	case BroadcastCompletedAreaListPresentEmergencyAreaIDBroadcastNR:
		var tmp EmergencyAreaIDBroadcastNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDBroadcastNR = &tmp
	}
	return
}
