package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       int64                   `lb:0,ub:63,madatory,valExt`
	RATType                 RATType                 `madatory`
	QoSFlowsTimedReportList []VolumeTimedReportItem `lb:1,ub:maxnoofTimePeriods,madatory`
	// IEExtensions *QoSFlowsUsageReportItemExtIEs `optional`
}

func (ie *QoSFlowsUsageReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_QosFlowIdentifier := NewINTEGER(ie.QosFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, true)
	if err = tmp_QosFlowIdentifier.Encode(w); err != nil {
		err = fmt.Errorf("Encode QosFlowIdentifier: %w", err)
		return
	}
	if err = ie.RATType.Encode(w); err != nil {
		err = fmt.Errorf("Encode RATType: %w", err)
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
			err = fmt.Errorf("Encode QoSFlowsTimedReportList: %w", err)
			return
		}
	} else {
		if err != nil {
			err = fmt.Errorf("QoSFlowsTimedReportList is nil: %w", err)
		}
		return
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
		ext: true,
	}
	if err = tmp_QosFlowIdentifier.Decode(r); err != nil {
		err = fmt.Errorf("Read QosFlowIdentifier: %w", err)
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	if err = ie.RATType.Decode(r); err != nil {
		err = fmt.Errorf("Read RATType: %w", err)
		return
	}
	tmp_QoSFlowsTimedReportList := Sequence[*VolumeTimedReportItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofTimePeriods},
		ext: false,
	}
	fn := func() *VolumeTimedReportItem { return new(VolumeTimedReportItem) }
	if err = tmp_QoSFlowsTimedReportList.Decode(r, fn); err != nil {
		err = fmt.Errorf("Read QoSFlowsTimedReportList: %w", err)
		return
	}
	ie.QoSFlowsTimedReportList = []VolumeTimedReportItem{}
	for _, i := range tmp_QoSFlowsTimedReportList.Value {
		ie.QoSFlowsTimedReportList = append(ie.QoSFlowsTimedReportList, *i)
	}
	return
}
