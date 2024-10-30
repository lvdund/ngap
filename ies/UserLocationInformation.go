package ies

import "github.com/lvdund/ngap/aper"

type UserLocationInformation struct {
	Choice                       uint64
	UserLocationInformationEUTRA *UserLocationInformationEUTRA `True,,,`
	UserLocationInformationNR    *UserLocationInformationNR    `True,,,`
	UserLocationInformationN3IWF *UserLocationInformationN3IWF `True,,,`
	// ChoiceExtensions *UserLocationInformationExtIEs `False,,,`
}

func (ie *UserLocationInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.UserLocationInformationEUTRA.Encode(w)
	case 2:
		err = ie.UserLocationInformationNR.Encode(w)
	case 3:
		err = ie.UserLocationInformationN3IWF.Encode(w)
	}
	return
}
func (ie *UserLocationInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp UserLocationInformationEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UserLocationInformationEUTRA = &tmp
	case 2:
		var tmp UserLocationInformationNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UserLocationInformationNR = &tmp
	case 3:
		var tmp UserLocationInformationN3IWF
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UserLocationInformationN3IWF = &tmp
	}
	return
}
