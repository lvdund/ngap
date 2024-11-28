package ies

import "github.com/lvdund/ngap/aper"

const (
	ResetTypePresentNothing uint64 = iota /* No components present */
	ResetTypePresentNGInterface
	ResetTypePresentPartOfNGInterface
	ResetTypePresentChoiceExtensions
)

type ResetType struct {
	Choice            uint64
	NGInterface       *ResetAll                            `False,,,`
	PartOfNGInterface *UEassociatedLogicalNGconnectionList `False,,,`
	// ChoiceExtensions *ResetTypeExtIEs `False,,,`
}

func (ie *ResetType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResetTypePresentNGInterface:
		err = ie.NGInterface.Encode(w)
	case ResetTypePresentPartOfNGInterface:
		err = ie.PartOfNGInterface.Encode(w)
	}
	return
}
func (ie *ResetType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResetTypePresentNGInterface:
		var tmp ResetAll
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NGInterface = &tmp
	case ResetTypePresentPartOfNGInterface:
		var tmp UEassociatedLogicalNGconnectionList
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PartOfNGInterface = &tmp
	}
	return
}
