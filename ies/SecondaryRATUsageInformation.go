package ies

import "github.com/lvdund/ngap/aper"

type SecondaryRATUsageInformation struct {
	PDUSessionUsageReport   *PDUSessionUsageReport    `optional`
	QosFlowsUsageReportList []QoSFlowsUsageReportItem `optional`
	// IEExtension *SecondaryRATUsageInformationExtIEs `optional`
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
		if len(ie.QosFlowsUsageReportList) > 0 {
			tmp := Sequence[*QoSFlowsUsageReportItem]{
				Value: []*QoSFlowsUsageReportItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range ie.QosFlowsUsageReportList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
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
	if aper.IsBitSet(optionals, 1) {
		if err = ie.PDUSessionUsageReport.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_QosFlowsUsageReportList := Sequence[*QoSFlowsUsageReportItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QoSFlowsUsageReportItem { return new(QoSFlowsUsageReportItem) }
		if err = tmp_QosFlowsUsageReportList.Decode(r, fn); err != nil {
			return
		}
		ie.QosFlowsUsageReportList = []QoSFlowsUsageReportItem{}
		for _, i := range tmp_QosFlowsUsageReportList.Value {
			ie.QosFlowsUsageReportList = append(ie.QosFlowsUsageReportList, *i)
		}
	}
	return
}
