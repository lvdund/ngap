package ies

import "github.com/lvdund/ngap/aper"

type WarningAreaList struct {
	Choice                 uint64
	EUTRACGIListForWarning *EUTRACGIListForWarning `False,,,`
	NRCGIListForWarning    *NRCGIListForWarning    `False,,,`
	TAIListForWarning      *TAIListForWarning      `False,,,`
	EmergencyAreaIDList    *EmergencyAreaIDList    `False,,,`
	// ChoiceExtensions *WarningAreaListExtIEs `False,,,`
}

func (ie *WarningAreaList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.EUTRACGIListForWarning.Encode(w)
	case 2:
		err = ie.NRCGIListForWarning.Encode(w)
	case 3:
		err = ie.TAIListForWarning.Encode(w)
	case 4:
		err = ie.EmergencyAreaIDList.Encode(w)
	}
	return
}
func (ie *WarningAreaList) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp EUTRACGIListForWarning
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EUTRACGIListForWarning = &tmp
	case 2:
		var tmp NRCGIListForWarning
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NRCGIListForWarning = &tmp
	case 3:
		var tmp TAIListForWarning
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.TAIListForWarning = &tmp
	case 4:
		var tmp EmergencyAreaIDList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EmergencyAreaIDList = &tmp
	}
	return
}
