package ies

import "github.com/lvdund/ngap/aper"

type QosFlowLevelQosParameters struct {
	QosCharacteristics             *QosCharacteristics             `False,`
	AllocationAndRetentionPriority *AllocationAndRetentionPriority `True,`
	GBRQosInformation              *GBRQosInformation              `True,OPTIONAL`
	ReflectiveQosAttribute         *ReflectiveQosAttribute         `False,OPTIONAL`
	AdditionalQosFlowInformation   *AdditionalQosFlowInformation   `False,OPTIONAL`
	// IEExtensions QosFlowLevelQosParametersExtIEs `False,OPTIONAL`
}

func (ie *QosFlowLevelQosParameters) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.GBRQosInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ReflectiveQosAttribute != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.AdditionalQosFlowInformation != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.QosCharacteristics != nil {
		if err = ie.QosCharacteristics.Encode(w); err != nil {
			return
		}
	}
	if ie.AllocationAndRetentionPriority != nil {
		if err = ie.AllocationAndRetentionPriority.Encode(w); err != nil {
			return
		}
	}
	if ie.GBRQosInformation != nil {
		if err = ie.GBRQosInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.ReflectiveQosAttribute != nil {
		if err = ie.ReflectiveQosAttribute.Encode(w); err != nil {
			return
		}
	}
	if ie.AdditionalQosFlowInformation != nil {
		if err = ie.AdditionalQosFlowInformation.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QosFlowLevelQosParameters) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	ie.QosCharacteristics = new(QosCharacteristics)
	ie.AllocationAndRetentionPriority = new(AllocationAndRetentionPriority)
	ie.GBRQosInformation = new(GBRQosInformation)
	ie.ReflectiveQosAttribute = new(ReflectiveQosAttribute)
	ie.AdditionalQosFlowInformation = new(AdditionalQosFlowInformation)
	if err = ie.QosCharacteristics.Decode(r); err != nil {
		return
	}
	if err = ie.AllocationAndRetentionPriority.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.GBRQosInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ReflectiveQosAttribute.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.AdditionalQosFlowInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
