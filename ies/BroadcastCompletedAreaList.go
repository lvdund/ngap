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
	CellIDBroadcastEUTRA          *CellIDBroadcastEUTRAItem
	TAIBroadcastEUTRA             *TAIBroadcastEUTRAItem
	EmergencyAreaIDBroadcastEUTRA *EmergencyAreaIDBroadcastEUTRAItem
	CellIDBroadcastNR             *CellIDBroadcastNRItem
	TAIBroadcastNR                *TAIBroadcastNRItem
	EmergencyAreaIDBroadcastNR    *EmergencyAreaIDBroadcastNRItem
	// ChoiceExtensions *BroadcastCompletedAreaListExtIEs
}

func (ie *BroadcastCompletedAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCompletedAreaListPresentCellidbroadcasteutra:
		err = ie.CellIDBroadcastEUTRA.Encode(w)
	case BroadcastCompletedAreaListPresentTaibroadcasteutra:
		err = ie.TAIBroadcastEUTRA.Encode(w)
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcasteutra:
		err = ie.EmergencyAreaIDBroadcastEUTRA.Encode(w)
	case BroadcastCompletedAreaListPresentCellidbroadcastnr:
		err = ie.CellIDBroadcastNR.Encode(w)
	case BroadcastCompletedAreaListPresentTaibroadcastnr:
		err = ie.TAIBroadcastNR.Encode(w)
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcastnr:
		err = ie.EmergencyAreaIDBroadcastNR.Encode(w)
	}
	return
}
func (ie *BroadcastCompletedAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCompletedAreaListPresentCellidbroadcasteutra:
		var tmp CellIDBroadcastEUTRAItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDBroadcastEUTRA = &tmp
	case BroadcastCompletedAreaListPresentTaibroadcasteutra:
		var tmp TAIBroadcastEUTRAItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIBroadcastEUTRA = &tmp
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcasteutra:
		var tmp EmergencyAreaIDBroadcastEUTRAItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDBroadcastEUTRA = &tmp
	case BroadcastCompletedAreaListPresentCellidbroadcastnr:
		var tmp CellIDBroadcastNRItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDBroadcastNR = &tmp
	case BroadcastCompletedAreaListPresentTaibroadcastnr:
		var tmp TAIBroadcastNRItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIBroadcastNR = &tmp
	case BroadcastCompletedAreaListPresentEmergencyareaidbroadcastnr:
		var tmp EmergencyAreaIDBroadcastNRItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDBroadcastNR = &tmp
	}
	return
}
