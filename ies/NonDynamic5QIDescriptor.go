package ies

import "github.com/lvdund/ngap/aper"

type NonDynamic5QIDescriptor struct {
	FiveQI                 int64
	PriorityLevelQos       *int64 `optional`
	AveragingWindow        *int64 `optional`
	MaximumDataBurstVolume *int64 `optional`
	// IEExtensions *NonDynamic5QIDescriptorExtIEs `optional`
}

func (ie *NonDynamic5QIDescriptor) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
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
	tmp_FiveQI := NewINTEGER(ie.FiveQI, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_FiveQI.Encode(w); err != nil {
		return
	}
	if ie.PriorityLevelQos != nil {
		tmp_PriorityLevelQos := NewINTEGER(*ie.PriorityLevelQos, aper.Constraint{Lb: 1, Ub: 127}, false)
		if err = tmp_PriorityLevelQos.Encode(w); err != nil {
			return
		}
	}
	if ie.AveragingWindow != nil {
		tmp_AveragingWindow := NewINTEGER(*ie.AveragingWindow, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp_AveragingWindow.Encode(w); err != nil {
			return
		}
	}
	if ie.MaximumDataBurstVolume != nil {
		tmp_MaximumDataBurstVolume := NewINTEGER(*ie.MaximumDataBurstVolume, aper.Constraint{Lb: 0, Ub: 4095}, false)
		if err = tmp_MaximumDataBurstVolume.Encode(w); err != nil {
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
	tmp_FiveQI := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_FiveQI.Decode(r); err != nil {
		return
	}
	ie.FiveQI = int64(tmp_FiveQI.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_PriorityLevelQos := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 127},
			ext: false,
		}
		if err = tmp_PriorityLevelQos.Decode(r); err != nil {
			return
		}
		ie.PriorityLevelQos = (*int64)(&tmp_PriorityLevelQos.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_AveragingWindow := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: false,
		}
		if err = tmp_AveragingWindow.Decode(r); err != nil {
			return
		}
		ie.AveragingWindow = (*int64)(&tmp_AveragingWindow.Value)
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_MaximumDataBurstVolume := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4095},
			ext: false,
		}
		if err = tmp_MaximumDataBurstVolume.Decode(r); err != nil {
			return
		}
		ie.MaximumDataBurstVolume = (*int64)(&tmp_MaximumDataBurstVolume.Value)
	}
	return
}
