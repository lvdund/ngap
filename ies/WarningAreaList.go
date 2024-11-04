package ies

import "github.com/lvdund/ngap/aper"

const (
	WarningAreaListPresentNothing uint64 = iota /* No components present */
	WarningAreaListPresentEUTRACGIListForWarning
	WarningAreaListPresentNRCGIListForWarning
	WarningAreaListPresentTAIListForWarning
	WarningAreaListPresentEmergencyAreaIDList
	WarningAreaListPresentChoiceExtensions
)

type WarningAreaList struct {
	Choice                 uint64
	EUTRACGIListForWarning *EUTRACGIListForWarning `False,,,`
	NRCGIListForWarning    *NRCGIListForWarning    `False,,,`
	TAIListForWarning      *TAIListForWarning      `False,,,`
	EmergencyAreaIDList    *EmergencyAreaIDList    `False,,,`
	// ChoiceExtensions *WarningAreaListExtIEs `False,,,`
}

func (ie *WarningAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return
	}
	switch ie.Choice {
	case WarningAreaListPresentEUTRACGIListForWarning:
		err = ie.EUTRACGIListForWarning.Encode(w)
	case WarningAreaListPresentNRCGIListForWarning:
		err = ie.NRCGIListForWarning.Encode(w)
	case WarningAreaListPresentTAIListForWarning:
		err = ie.TAIListForWarning.Encode(w)
	case WarningAreaListPresentEmergencyAreaIDList:
		err = ie.EmergencyAreaIDList.Encode(w)
	}
	return
}
func (ie *WarningAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return
	}
	switch ie.Choice {
	case WarningAreaListPresentEUTRACGIListForWarning:
		var tmp EUTRACGIListForWarning
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIListForWarning = &tmp
	case WarningAreaListPresentNRCGIListForWarning:
		var tmp NRCGIListForWarning
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIListForWarning = &tmp
	case WarningAreaListPresentTAIListForWarning:
		var tmp TAIListForWarning
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIListForWarning = &tmp
	case WarningAreaListPresentEmergencyAreaIDList:
		var tmp EmergencyAreaIDList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDList = &tmp
	}
	return
}
