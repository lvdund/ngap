package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionUsageReport struct {
	RATType                   RATType                 `madatory`
	PDUSessionTimedReportList []VolumeTimedReportItem `lb:1,ub:maxnoofTimePeriods,madatory`
	// IEExtensions *PDUSessionUsageReportExtIEs `optional`
}

func (ie *PDUSessionUsageReport) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.RATType.Encode(w); err != nil {
		err = utils.WrapError("Encode RATType", err)
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
			err = utils.WrapError("Encode PDUSessionTimedReportList", err)
			return
		}
	} else {
		err = utils.WrapError("PDUSessionTimedReportList is nil", err)
		return
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
	if err = ie.RATType.Decode(r); err != nil {
		err = utils.WrapError("Read RATType", err)
		return
	}
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
