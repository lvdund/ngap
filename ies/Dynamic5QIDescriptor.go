package ies

import "github.com/lvdund/ngap/aper"

type Dynamic5QIDescriptor struct {
	PriorityLevelQos       *PriorityLevelQos       `False,`
	PacketDelayBudget      *PacketDelayBudget      `False,`
	PacketErrorRate        *PacketErrorRate        `True,`
	FiveQI                 *FiveQI                 `False,OPTIONAL`
	DelayCritical          *DelayCritical          `False,OPTIONAL`
	AveragingWindow        *AveragingWindow        `False,OPTIONAL`
	MaximumDataBurstVolume *MaximumDataBurstVolume `False,OPTIONAL`
	// IEExtensions Dynamic5QIDescriptorExtIEs `False,OPTIONAL`
}

func (ie *Dynamic5QIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.FiveQI != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.DelayCritical != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.AveragingWindow != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.MaximumDataBurstVolume != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if ie.PriorityLevelQos != nil {
		if err = ie.PriorityLevelQos.Encode(w); err != nil {
			return
		}
	}
	if ie.PacketDelayBudget != nil {
		if err = ie.PacketDelayBudget.Encode(w); err != nil {
			return
		}
	}
	if ie.PacketErrorRate != nil {
		if err = ie.PacketErrorRate.Encode(w); err != nil {
			return
		}
	}
	if ie.FiveQI != nil {
		if err = ie.FiveQI.Encode(w); err != nil {
			return
		}
	}
	if ie.DelayCritical != nil {
		if err = ie.DelayCritical.Encode(w); err != nil {
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
func (ie *Dynamic5QIDescriptor) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	ie.PriorityLevelQos = new(PriorityLevelQos)
	ie.PacketDelayBudget = new(PacketDelayBudget)
	ie.PacketErrorRate = new(PacketErrorRate)
	ie.FiveQI = new(FiveQI)
	ie.DelayCritical = new(DelayCritical)
	ie.AveragingWindow = new(AveragingWindow)
	ie.MaximumDataBurstVolume = new(MaximumDataBurstVolume)
	if err = ie.PriorityLevelQos.Decode(r); err != nil {
		return
	}
	if err = ie.PacketDelayBudget.Decode(r); err != nil {
		return
	}
	if err = ie.PacketErrorRate.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.FiveQI.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.DelayCritical.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.AveragingWindow.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 5) {
		if err = ie.MaximumDataBurstVolume.Decode(r); err != nil {
			return
		}
	}
	return
}
