package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	UserLocationInformationPresentNothing uint64 = iota
	UserLocationInformationPresentUserlocationinformationeutra
	UserLocationInformationPresentUserlocationinformationnr
	UserLocationInformationPresentUserlocationinformationn3Iwf
	UserLocationInformationPresentChoiceExtensions
)

type UserLocationInformation struct {
	Choice                       uint64
	UserLocationInformationEUTRA *UserLocationInformationEUTRA
	UserLocationInformationNR    *UserLocationInformationNR
	UserLocationInformationN3IWF *UserLocationInformationN3IWF
	// ChoiceExtensions *UserLocationInformationExtIEs
}

func (ie *UserLocationInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case UserLocationInformationPresentUserlocationinformationeutra:
		err = ie.UserLocationInformationEUTRA.Encode(w)
	case UserLocationInformationPresentUserlocationinformationnr:
		err = ie.UserLocationInformationNR.Encode(w)
	case UserLocationInformationPresentUserlocationinformationn3Iwf:
		err = ie.UserLocationInformationN3IWF.Encode(w)
	}
	return
}
func (ie *UserLocationInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case UserLocationInformationPresentUserlocationinformationeutra:
		var tmp UserLocationInformationEUTRA
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read UserLocationInformationEUTRA", err)
			return
		}
		ie.UserLocationInformationEUTRA = &tmp
	case UserLocationInformationPresentUserlocationinformationnr:
		var tmp UserLocationInformationNR
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read UserLocationInformationNR", err)
			return
		}
		ie.UserLocationInformationNR = &tmp
	case UserLocationInformationPresentUserlocationinformationn3Iwf:
		var tmp UserLocationInformationN3IWF
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read UserLocationInformationN3IWF", err)
			return
		}
		ie.UserLocationInformationN3IWF = &tmp
	}
	return
}
