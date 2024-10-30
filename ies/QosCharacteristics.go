package ies

import "github.com/lvdund/ngap/aper"

type QosCharacteristics struct {
	Choice        uint64
	NonDynamic5QI *NonDynamic5QIDescriptor `True,,,`
	Dynamic5QI    *Dynamic5QIDescriptor    `True,,,`
	// ChoiceExtensions *QosCharacteristicsExtIEs `False,,,`
}

func (ie *QosCharacteristics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.NonDynamic5QI.Encode(w)
	case 2:
		err = ie.Dynamic5QI.Encode(w)
	}
	return
}
func (ie *QosCharacteristics) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp NonDynamic5QIDescriptor
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NonDynamic5QI = &tmp
	case 2:
		var tmp Dynamic5QIDescriptor
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Dynamic5QI = &tmp
	}
	return
}
