package ies

import "github.com/lvdund/ngap/aper"

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
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.CellIDBroadcastEUTRA.Encode(w)
	case 2:
		err = ie.TAIBroadcastEUTRA.Encode(w)
	case 3:
		err = ie.EmergencyAreaIDBroadcastEUTRA.Encode(w)
	case 4:
		err = ie.CellIDBroadcastNR.Encode(w)
	case 5:
		err = ie.TAIBroadcastNR.Encode(w)
	case 6:
		err = ie.EmergencyAreaIDBroadcastNR.Encode(w)
	}
	return
}
func (ie *BroadcastCompletedAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp CellIDBroadcastEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDBroadcastEUTRA = &tmp
	case 2:
		var tmp TAIBroadcastEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIBroadcastEUTRA = &tmp
	case 3:
		var tmp EmergencyAreaIDBroadcastEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDBroadcastEUTRA = &tmp
	case 4:
		var tmp CellIDBroadcastNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDBroadcastNR = &tmp
	case 5:
		var tmp TAIBroadcastNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIBroadcastNR = &tmp
	case 6:
		var tmp EmergencyAreaIDBroadcastNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDBroadcastNR = &tmp
	}
	return
}
