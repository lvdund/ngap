package ies

import (
	"github.com/lvdund/ngap/aper"
)

const (
	GlobalRANNodeIDPresentNothing uint64 = iota /* No components present */
	GlobalRANNodeIDPresentGlobalGNBID
	GlobalRANNodeIDPresentGlobalNgENBID
	GlobalRANNodeIDPresentGlobalN3IWFID
	GlobalRANNodeIDPresentChoiceExtensions
)

type GlobalRANNodeID struct {
	Choice        uint64
	GlobalGNBID   *GlobalGNBID   `True,,,`
	GlobalNgENBID *GlobalNgENBID `True,,,`
	GlobalN3IWFID *GlobalN3IWFID `True,,,`
	// ChoiceExtensions *GlobalRANNodeIDExtIEs `False,,,`
}

func (ie *GlobalRANNodeID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case GlobalRANNodeIDPresentGlobalGNBID:
		err = ie.GlobalGNBID.Encode(w)
	case GlobalRANNodeIDPresentGlobalNgENBID:
		err = ie.GlobalNgENBID.Encode(w)
	case GlobalRANNodeIDPresentGlobalN3IWFID:
		err = ie.GlobalN3IWFID.Encode(w)
	}
	return
}
func (ie *GlobalRANNodeID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case GlobalRANNodeIDPresentGlobalGNBID:
		var tmp GlobalGNBID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalGNBID = &tmp
	case GlobalRANNodeIDPresentGlobalNgENBID:
		var tmp GlobalNgENBID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalNgENBID = &tmp
	case GlobalRANNodeIDPresentGlobalN3IWFID:
		var tmp GlobalN3IWFID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalN3IWFID = &tmp
	}
	return
}
