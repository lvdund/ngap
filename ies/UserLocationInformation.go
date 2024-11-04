package ies

import "github.com/lvdund/ngap/aper"

const (
	UserLocationInformationPresentNothing uint64 = iota /* No components present */
	UserLocationInformationPresentUserLocationInformationEUTRA
	UserLocationInformationPresentUserLocationInformationNR
	UserLocationInformationPresentUserLocationInformationN3IWF
	UserLocationInformationPresentChoiceExtensions
)

type UserLocationInformation struct {
	Choice                       uint64
	UserLocationInformationEUTRA *UserLocationInformationEUTRA `True,,,`
	UserLocationInformationNR    *UserLocationInformationNR    `True,,,`
	UserLocationInformationN3IWF *UserLocationInformationN3IWF `True,,,`
	// ChoiceExtensions *UserLocationInformationExtIEs `False,,,`
}

func (ie *UserLocationInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return
	}
	switch ie.Choice {
	case UserLocationInformationPresentUserLocationInformationEUTRA:
		err = ie.UserLocationInformationEUTRA.Encode(w)
	case UserLocationInformationPresentUserLocationInformationNR:
		err = ie.UserLocationInformationNR.Encode(w)
	case UserLocationInformationPresentUserLocationInformationN3IWF:
		err = ie.UserLocationInformationN3IWF.Encode(w)
	}
	return
}
func (ie *UserLocationInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return
	}
	switch ie.Choice {
	case UserLocationInformationPresentUserLocationInformationEUTRA:
		var tmp UserLocationInformationEUTRA
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UserLocationInformationEUTRA = &tmp
	case UserLocationInformationPresentUserLocationInformationNR:
		var tmp UserLocationInformationNR
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UserLocationInformationNR = &tmp
	case UserLocationInformationPresentUserLocationInformationN3IWF:
		var tmp UserLocationInformationN3IWF
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.UserLocationInformationN3IWF = &tmp
	}
	return
}
