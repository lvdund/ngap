package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type Dynamic5QIDescriptor struct {
	PriorityLevelQos       int64           `lb:1,ub:127,madatory,valExt`
	PacketDelayBudget      int64           `lb:0,ub:1023,madatory,valExt`
	PacketErrorRate        PacketErrorRate `madatory`
	FiveQI                 *int64          `lb:0,ub:255,optional,valExt`
	DelayCritical          *DelayCritical  `optional`
	AveragingWindow        *int64          `lb:0,ub:4095,optional,valExt`
	MaximumDataBurstVolume *int64          `lb:0,ub:4095,optional,valExt`
	// IEExtensions *Dynamic5QIDescriptorExtIEs `optional`
}

func (ie *Dynamic5QIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
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
	tmp_PriorityLevelQos := NewINTEGER(ie.PriorityLevelQos, aper.Constraint{Lb: 1, Ub: 127}, true)
	if err = tmp_PriorityLevelQos.Encode(w); err != nil {
		err = fmt.Errorf("Encode PriorityLevelQos: %w", err)
		return
	}
	tmp_PacketDelayBudget := NewINTEGER(ie.PacketDelayBudget, aper.Constraint{Lb: 0, Ub: 1023}, true)
	if err = tmp_PacketDelayBudget.Encode(w); err != nil {
		err = fmt.Errorf("Encode PacketDelayBudget: %w", err)
		return
	}
	if err = ie.PacketErrorRate.Encode(w); err != nil {
		err = fmt.Errorf("Encode PacketErrorRate: %w", err)
		return
	}
	if ie.FiveQI != nil {
		tmp_FiveQI := NewINTEGER(*ie.FiveQI, aper.Constraint{Lb: 0, Ub: 255}, true)
		if err = tmp_FiveQI.Encode(w); err != nil {
			err = fmt.Errorf("Encode FiveQI: %w", err)
			return
		}
	}
	if ie.DelayCritical != nil {
		if err = ie.DelayCritical.Encode(w); err != nil {
			err = fmt.Errorf("Encode DelayCritical: %w", err)
			return
		}
	}
	if ie.AveragingWindow != nil {
		tmp_AveragingWindow := NewINTEGER(*ie.AveragingWindow, aper.Constraint{Lb: 0, Ub: 4095}, true)
		if err = tmp_AveragingWindow.Encode(w); err != nil {
			err = fmt.Errorf("Encode AveragingWindow: %w", err)
			return
		}
	}
	if ie.MaximumDataBurstVolume != nil {
		tmp_MaximumDataBurstVolume := NewINTEGER(*ie.MaximumDataBurstVolume, aper.Constraint{Lb: 0, Ub: 4095}, true)
		if err = tmp_MaximumDataBurstVolume.Encode(w); err != nil {
			err = fmt.Errorf("Encode MaximumDataBurstVolume: %w", err)
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
	tmp_PriorityLevelQos := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 127},
		ext: true,
	}
	if err = tmp_PriorityLevelQos.Decode(r); err != nil {
		err = fmt.Errorf("Read PriorityLevelQos: %w", err)
		return
	}
	ie.PriorityLevelQos = int64(tmp_PriorityLevelQos.Value)
	tmp_PacketDelayBudget := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 1023},
		ext: true,
	}
	if err = tmp_PacketDelayBudget.Decode(r); err != nil {
		err = fmt.Errorf("Read PacketDelayBudget: %w", err)
		return
	}
	ie.PacketDelayBudget = int64(tmp_PacketDelayBudget.Value)
	if err = ie.PacketErrorRate.Decode(r); err != nil {
		err = fmt.Errorf("Read PacketErrorRate: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_FiveQI := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 255},
			ext: true,
		}
		if err = tmp_FiveQI.Decode(r); err != nil {
			err = fmt.Errorf("Read FiveQI: %w", err)
			return
		}
		ie.FiveQI = (*int64)(&tmp_FiveQI.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(DelayCritical)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read DelayCritical: %w", err)
			return
		}
		ie.DelayCritical = tmp
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_AveragingWindow := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: true,
		}
		if err = tmp_AveragingWindow.Decode(r); err != nil {
			err = fmt.Errorf("Read AveragingWindow: %w", err)
			return
		}
		ie.AveragingWindow = (*int64)(&tmp_AveragingWindow.Value)
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_MaximumDataBurstVolume := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: true,
		}
		if err = tmp_MaximumDataBurstVolume.Decode(r); err != nil {
			err = fmt.Errorf("Read MaximumDataBurstVolume: %w", err)
			return
		}
		ie.MaximumDataBurstVolume = (*int64)(&tmp_MaximumDataBurstVolume.Value)
	}
	return
}
