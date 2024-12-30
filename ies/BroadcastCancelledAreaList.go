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
	CellIDCancelledEUTRA          *CellIDCancelledEUTRAItem
	TAICancelledEUTRA             *TAICancelledEUTRAItem
	EmergencyAreaIDCancelledEUTRA *EmergencyAreaIDCancelledEUTRAItem
	CellIDCancelledNR             *CellIDCancelledNRItem
	TAICancelledNR                *TAICancelledNRItem
	EmergencyAreaIDCancelledNR    *EmergencyAreaIDCancelledNRItem
	// ChoiceExtensions *BroadcastCancelledAreaListExtIEs
}

func (ie *BroadcastCancelledAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCancelledAreaListPresentCellidcancelledeutra:
		err = ie.CellIDCancelledEUTRA.Encode(w)
	case BroadcastCancelledAreaListPresentTaicancelledeutra:
		err = ie.TAICancelledEUTRA.Encode(w)
	case BroadcastCancelledAreaListPresentEmergencyareaidcancelledeutra:
		err = ie.EmergencyAreaIDCancelledEUTRA.Encode(w)
	case BroadcastCancelledAreaListPresentCellidcancellednr:
		err = ie.CellIDCancelledNR.Encode(w)
	case BroadcastCancelledAreaListPresentTaicancellednr:
		err = ie.TAICancelledNR.Encode(w)
	case BroadcastCancelledAreaListPresentEmergencyareaidcancellednr:
		err = ie.EmergencyAreaIDCancelledNR.Encode(w)
	}
	return
}
func (ie *BroadcastCancelledAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return
	}
	switch ie.Choice {
	case BroadcastCancelledAreaListPresentCellidcancelledeutra:
		var tmp CellIDCancelledEUTRAItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDCancelledEUTRA = &tmp
	case BroadcastCancelledAreaListPresentTaicancelledeutra:
		var tmp TAICancelledEUTRAItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAICancelledEUTRA = &tmp
	case BroadcastCancelledAreaListPresentEmergencyareaidcancelledeutra:
		var tmp EmergencyAreaIDCancelledEUTRAItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDCancelledEUTRA = &tmp
	case BroadcastCancelledAreaListPresentCellidcancellednr:
		var tmp CellIDCancelledNRItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.CellIDCancelledNR = &tmp
	case BroadcastCancelledAreaListPresentTaicancellednr:
		var tmp TAICancelledNRItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAICancelledNR = &tmp
	case BroadcastCancelledAreaListPresentEmergencyareaidcancellednr:
		var tmp EmergencyAreaIDCancelledNRItem
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDCancelledNR = &tmp
	}
	return
}
