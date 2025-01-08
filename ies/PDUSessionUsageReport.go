package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionUsageReport struct {
	RATType                   int64
	PDUSessionTimedReportList []VolumeTimedReportItem
	// IEExtensions *PDUSessionUsageReportExtIEs `optional`
}

func (ie *PDUSessionUsageReport) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_RATType := NewENUMERATED(ie.RATType, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_RATType.Encode(w); err != nil {
		err = utils.WrapError("Read RATType", err)
		return
	}
	if len(ie.PDUSessionTimedReportList) > 0 {
		tmp := Sequence[*VolumeTimedReportItem]{
			Value: []*VolumeTimedReportItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofTimePeriods},
			ext:   false,
		}
		for _, i := range ie.PDUSessionTimedReportList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Read PDUSessionTimedReportList", err)
			return
		}
	}
	return
}
func (ie *PDUSessionUsageReport) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_RATType := ENUMERATED{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_RATType.Decode(r); err != nil {
		err = utils.WrapError("Read RATType", err)
		return
	}
	ie.RATType = int64(tmp_RATType.Value)
	tmp_PDUSessionTimedReportList := Sequence[*VolumeTimedReportItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofTimePeriods},
		ext: false,
	}
	fn := func() *VolumeTimedReportItem { return new(VolumeTimedReportItem) }
	if err = tmp_PDUSessionTimedReportList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read PDUSessionTimedReportList", err)
		return
	}
	ie.PDUSessionTimedReportList = []VolumeTimedReportItem{}
	for _, i := range tmp_PDUSessionTimedReportList.Value {
		ie.PDUSessionTimedReportList = append(ie.PDUSessionTimedReportList, *i)
	}
	return
}
