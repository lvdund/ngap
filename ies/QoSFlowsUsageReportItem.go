package ies

import "github.com/lvdund/ngap/aper"

type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       int64
	RATType                 int64
	QoSFlowsTimedReportList []VolumeTimedReportItem
	// IEExtensions *QoSFlowsUsageReportItemExtIEs `optional`
}

func (ie *QoSFlowsUsageReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_QosFlowIdentifier := NewINTEGER(ie.QosFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_QosFlowIdentifier.Encode(w); err != nil {
		return
	}
	tmp_RATType := NewENUMERATED(ie.RATType, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_RATType.Encode(w); err != nil {
		return
	}
	if len(ie.QoSFlowsTimedReportList) > 0 {
		tmp := Sequence[*VolumeTimedReportItem]{
			Value: []*VolumeTimedReportItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofTimePeriods},
			ext:   false,
		}
		for _, i := range ie.QoSFlowsTimedReportList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QoSFlowsUsageReportItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_QosFlowIdentifier := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 63},
		ext: false,
	}
	if err = tmp_QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	tmp_RATType := ENUMERATED{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_RATType.Decode(r); err != nil {
		return
	}
	ie.RATType = int64(tmp_RATType.Value)
	tmp_QoSFlowsTimedReportList := Sequence[*VolumeTimedReportItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofTimePeriods},
		ext: false,
	}
	fn := func() *VolumeTimedReportItem { return new(VolumeTimedReportItem) }
	if err = tmp_QoSFlowsTimedReportList.Decode(r, fn); err != nil {
		return
	}
	ie.QoSFlowsTimedReportList = []VolumeTimedReportItem{}
	for _, i := range tmp_QoSFlowsTimedReportList.Value {
		ie.QoSFlowsTimedReportList = append(ie.QoSFlowsTimedReportList, *i)
	}
	return
}
