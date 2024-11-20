package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionUsageReport struct {
	PDUSessionTimedReportList *VolumeTimedReportList `False,`
	// IEExtensions PDUSessionUsageReportExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionUsageReport) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PDUSessionTimedReportList != nil {
		if err = ie.PDUSessionTimedReportList.Encode(w); err != nil {
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
	ie.PDUSessionTimedReportList = new(VolumeTimedReportList)
	if err = ie.PDUSessionTimedReportList.Decode(r); err != nil {
		return
	}
	return
}
