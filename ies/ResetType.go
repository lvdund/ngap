package ies

import "github.com/lvdund/ngap/aper"

type ResetType struct {
	Choice            uint64
	NGInterface       *ResetAll                            `False,,,`
	PartOfNGInterface *UEassociatedLogicalNGconnectionList `False,,,`
	// ChoiceExtensions *ResetTypeExtIEs `False,,,`
}

func (ie *ResetType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.NGInterface.Encode(w)
	case 2:
		err = ie.PartOfNGInterface.Encode(w)
	}
	return
}
func (ie *ResetType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp ResetAll
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NGInterface = &tmp
	case 2:
		var tmp UEassociatedLogicalNGconnectionList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PartOfNGInterface = &tmp
	}
	return
}
