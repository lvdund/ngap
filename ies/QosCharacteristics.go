package ies

import "github.com/lvdund/ngap/aper"

const (
	QosCharacteristicsPresentNothing uint64 = iota /* No components present */
	QosCharacteristicsPresentNonDynamic5QI
	QosCharacteristicsPresentDynamic5QI
	QosCharacteristicsPresentChoiceExtensions
)

type QosCharacteristics struct {
	Choice        uint64
	NonDynamic5QI *NonDynamic5QIDescriptor `True,,,`
	Dynamic5QI    *Dynamic5QIDescriptor    `True,,,`
	// ChoiceExtensions *QosCharacteristicsExtIEs `False,,,`
}

func (ie *QosCharacteristics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return
	}
	switch ie.Choice {
	case QosCharacteristicsPresentNonDynamic5QI:
		err = ie.NonDynamic5QI.Encode(w)
	case QosCharacteristicsPresentDynamic5QI:
		err = ie.Dynamic5QI.Encode(w)
	}
	return
}
func (ie *QosCharacteristics) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return
	}
	switch ie.Choice {
	case QosCharacteristicsPresentNonDynamic5QI:
		var tmp NonDynamic5QIDescriptor
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.NonDynamic5QI = &tmp
	case QosCharacteristicsPresentDynamic5QI:
		var tmp Dynamic5QIDescriptor
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.Dynamic5QI = &tmp
	}
	return
}
