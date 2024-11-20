package ies

import "github.com/lvdund/ngap/aper"

type SecondaryRATUsageInformation struct {
	PDUSessionUsageReport   *PDUSessionUsageReport   `True,OPTIONAL`
	QosFlowsUsageReportList *QoSFlowsUsageReportList `False,OPTIONAL`
	// IEExtension SecondaryRATUsageInformationExtIEs `False,OPTIONAL`
}

func (ie *SecondaryRATUsageInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PDUSessionUsageReport != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.QosFlowsUsageReportList != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.PDUSessionUsageReport != nil {
		if err = ie.PDUSessionUsageReport.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowsUsageReportList != nil {
		if err = ie.QosFlowsUsageReportList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *SecondaryRATUsageInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	ie.PDUSessionUsageReport = new(PDUSessionUsageReport)
	ie.QosFlowsUsageReportList = new(QoSFlowsUsageReportList)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.PDUSessionUsageReport.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.QosFlowsUsageReportList.Decode(r); err != nil {
			return
		}
	}
	return
}
