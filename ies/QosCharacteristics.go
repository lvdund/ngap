package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	QosCharacteristicsPresentNothing uint64 = iota
	QosCharacteristicsPresentNondynamic5Qi
	QosCharacteristicsPresentDynamic5Qi
	QosCharacteristicsPresentChoiceExtensions
)

type QosCharacteristics struct {
	Choice        uint64
	NonDynamic5QI *NonDynamic5QIDescriptor
	Dynamic5QI    *Dynamic5QIDescriptor
	// ChoiceExtensions *QosCharacteristicsExtIEs
}

func (ie *QosCharacteristics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case QosCharacteristicsPresentNondynamic5Qi:
		err = ie.NonDynamic5QI.Encode(w)
	case QosCharacteristicsPresentDynamic5Qi:
		err = ie.Dynamic5QI.Encode(w)
	}
	return
}
func (ie *QosCharacteristics) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case QosCharacteristicsPresentNondynamic5Qi:
		var tmp NonDynamic5QIDescriptor
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read NonDynamic5QI", err)
			return
		}
		ie.NonDynamic5QI = &tmp
	case QosCharacteristicsPresentDynamic5Qi:
		var tmp Dynamic5QIDescriptor
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read Dynamic5QI", err)
			return
		}
		ie.Dynamic5QI = &tmp
	}
	return
}
