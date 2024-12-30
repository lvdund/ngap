package ies

import "github.com/lvdund/ngap/aper"

const (
	WarningAreaListPresentNothing uint64 = iota
	WarningAreaListPresentEutraCgilistforwarning
	WarningAreaListPresentNrCgilistforwarning
	WarningAreaListPresentTailistforwarning
	WarningAreaListPresentEmergencyareaidlist
	WarningAreaListPresentChoiceExtensions
)

type WarningAreaList struct {
	Choice                 uint64
	EUTRACGIListForWarning *EUTRACGI
	NRCGIListForWarning    *NRCGI
	TAIListForWarning      *TAI
	// EmergencyAreaIDList    *EmergencyAreaID
	// ChoiceExtensions *WarningAreaListExtIEs
}

func (ie *WarningAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return
	}
	switch ie.Choice {
	case WarningAreaListPresentEutraCgilistforwarning:
		err = ie.EUTRACGIListForWarning.Encode(w)
	case WarningAreaListPresentNrCgilistforwarning:
		err = ie.NRCGIListForWarning.Encode(w)
	case WarningAreaListPresentTailistforwarning:
		err = ie.TAIListForWarning.Encode(w)
	case WarningAreaListPresentEmergencyareaidlist:
		// err = ie.EmergencyAreaIDList.Encode(w)
	}
	return
}
func (ie *WarningAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return
	}
	switch ie.Choice {
	case WarningAreaListPresentEutraCgilistforwarning:
		var tmp EUTRACGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIListForWarning = &tmp
	case WarningAreaListPresentNrCgilistforwarning:
		var tmp NRCGI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIListForWarning = &tmp
	case WarningAreaListPresentTailistforwarning:
		var tmp TAI
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIListForWarning = &tmp
	case WarningAreaListPresentEmergencyareaidlist:
		// var tmp EmergencyAreaID
		// if err = tmp.Decode(r); err != nil {
		// 	return
		// }
		// ie.EmergencyAreaIDList = &tmp
	}
	return
}
