package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
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
		err = utils.WrapError("Encode QosFlowIdentifier", err)
		return
	}
	if err = ie.RATType.Encode(w); err != nil {
		err = utils.WrapError("Encode RATType", err)
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
			err = utils.WrapError("Encode QoSFlowsTimedReportList", err)
			return
		}
	} else {
		err = utils.WrapError("QoSFlowsTimedReportList is nil", err)
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
		err = utils.WrapError("Read QosFlowIdentifier", err)
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	if err = ie.RATType.Decode(r); err != nil {
		err = utils.WrapError("Read RATType", err)
		return
	}
	tmp_QoSFlowsTimedReportList := Sequence[*VolumeTimedReportItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofTimePeriods},
		ext: false,
	}
	fn := func() *VolumeTimedReportItem { return new(VolumeTimedReportItem) }
	if err = tmp_QoSFlowsTimedReportList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read QoSFlowsTimedReportList", err)
		return
	}
	ie.QoSFlowsTimedReportList = []VolumeTimedReportItem{}
	for _, i := range tmp_QoSFlowsTimedReportList.Value {
		ie.QoSFlowsTimedReportList = append(ie.QoSFlowsTimedReportList, *i)
	}
	return
}
