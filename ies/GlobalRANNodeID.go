package ies

import "github.com/lvdund/ngap/aper"

const (
	GlobalRANNodeIDPresentNothing uint64 = iota
	GlobalRANNodeIDPresentGlobalgnbId
	GlobalRANNodeIDPresentGlobalngenbId
	GlobalRANNodeIDPresentGlobaln3IwfId
	GlobalRANNodeIDPresentChoiceExtensions
)

type GlobalRANNodeID struct {
	Choice        uint64
	GlobalGNBID   *GlobalGNBID
	GlobalNgENBID *GlobalNgENBID
	GlobalN3IWFID *GlobalN3IWFID
	// ChoiceExtensions *GlobalRANNodeIDExtIEs
}

func (ie *GlobalRANNodeID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case GlobalRANNodeIDPresentGlobalgnbId:
		err = ie.GlobalGNBID.Encode(w)
	case GlobalRANNodeIDPresentGlobalngenbId:
		err = ie.GlobalNgENBID.Encode(w)
	case GlobalRANNodeIDPresentGlobaln3IwfId:
		err = ie.GlobalN3IWFID.Encode(w)
	}
	return
}
func (ie *GlobalRANNodeID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case GlobalRANNodeIDPresentGlobalgnbId:
		var tmp GlobalGNBID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalGNBID = &tmp
	case GlobalRANNodeIDPresentGlobalngenbId:
		var tmp GlobalNgENBID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalNgENBID = &tmp
	case GlobalRANNodeIDPresentGlobaln3IwfId:
		var tmp GlobalN3IWFID
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GlobalN3IWFID = &tmp
	}
	return
}
