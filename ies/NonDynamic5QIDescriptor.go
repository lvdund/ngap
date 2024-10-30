package ies

import "github.com/lvdund/ngap/aper"

type NonDynamic5QIDescriptor struct {
	FiveQI                 *FiveQI                 `False,`
	PriorityLevelQos       *PriorityLevelQos       `False,OPTIONAL`
	AveragingWindow        *AveragingWindow        `False,OPTIONAL`
	MaximumDataBurstVolume *MaximumDataBurstVolume `False,OPTIONAL`
	// IEExtensions NonDynamic5QIDescriptorExtIEs `False,OPTIONAL`
}

func (ie *NonDynamic5QIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PriorityLevelQos != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.AveragingWindow != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.MaximumDataBurstVolume != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.FiveQI != nil {
		if err = ie.FiveQI.Encode(w); err != nil {
			return
		}
	}
	if ie.PriorityLevelQos != nil {
		if err = ie.PriorityLevelQos.Encode(w); err != nil {
			return
		}
	}
	if ie.AveragingWindow != nil {
		if err = ie.AveragingWindow.Encode(w); err != nil {
			return
		}
	}
	if ie.MaximumDataBurstVolume != nil {
		if err = ie.MaximumDataBurstVolume.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *NonDynamic5QIDescriptor) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	ie.FiveQI = new(FiveQI)
	ie.PriorityLevelQos = new(PriorityLevelQos)
	ie.AveragingWindow = new(AveragingWindow)
	ie.MaximumDataBurstVolume = new(MaximumDataBurstVolume)
	if err = ie.FiveQI.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.PriorityLevelQos.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.AveragingWindow.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.MaximumDataBurstVolume.Decode(r); err != nil {
			return
		}
	}
	return
}
