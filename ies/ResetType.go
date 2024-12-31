package ies

import "github.com/lvdund/ngap/aper"

const (
	ResetTypePresentNothing uint64 = iota
	ResetTypePresentNgInterface
	ResetTypePresentPartofngInterface
	ResetTypePresentChoiceExtensions
)

type ResetType struct {
	Choice            uint64
	NGInterface       *ResetAll
	PartOfNGInterface *[]UEassociatedLogicalNGconnectionItem
	// ChoiceExtensions *ResetTypeExtIEs
}

func (ie *ResetType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResetTypePresentNgInterface:
		err = ie.NGInterface.Encode(w)
	case ResetTypePresentPartofngInterface:
		tmp := Sequence[*UEassociatedLogicalNGconnectionItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofNGConnectionsToReset},
			ext: false,
		}
		for _, i := range *ie.PartOfNGInterface {
			tmp.Value = append(tmp.Value, &i)
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *ResetType) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case ResetTypePresentNgInterface:
		var tmp ResetAll
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NGInterface = &tmp
	case ResetTypePresentPartofngInterface:
		tmp := Sequence[*UEassociatedLogicalNGconnectionItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofNGConnectionsToReset},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.PartOfNGInterface = &[]UEassociatedLogicalNGconnectionItem{}
		for _, i := range tmp.Value {
			*ie.PartOfNGInterface = append(*ie.PartOfNGInterface, *i)
		}
	}
	return
}
