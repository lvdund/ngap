package ies

import "github.com/lvdund/ngap/aper"

type GlobalRANNodeID struct {
	Choice        uint64
	GlobalGNBID   *GlobalGNBID   `True,,,`
	GlobalNgENBID *GlobalNgENBID `True,,,`
	GlobalN3IWFID *GlobalN3IWFID `True,,,`
	// ChoiceExtensions *GlobalRANNodeIDExtIEs `False,,,`
}

func (ie *GlobalRANNodeID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.GlobalGNBID.Encode(w)
	case 2:
		err = ie.GlobalNgENBID.Encode(w)
	case 3:
		err = ie.GlobalN3IWFID.Encode(w)
	}
	return
}
func (ie *GlobalRANNodeID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp GlobalGNBID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalGNBID = &tmp
	case 2:
		var tmp GlobalNgENBID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalNgENBID = &tmp
	case 3:
		var tmp GlobalN3IWFID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalN3IWFID = &tmp
	}
	return
}
