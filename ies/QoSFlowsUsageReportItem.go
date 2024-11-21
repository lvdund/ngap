package ies

import "github.com/lvdund/ngap/aper"

type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       *QosFlowIdentifier     `False,`
	QoSFlowsTimedReportList *VolumeTimedReportList `False,`
	// IEExtensions QoSFlowsUsageReportItemExtIEs `False,OPTIONAL`
}

func (ie *QoSFlowsUsageReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.QosFlowIdentifier != nil {
		if err = ie.QosFlowIdentifier.Encode(w); err != nil {
			return
		}
	}
	if ie.QoSFlowsTimedReportList != nil {
		if err = ie.QoSFlowsTimedReportList.Encode(w); err != nil {
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
	ie.QosFlowIdentifier = new(QosFlowIdentifier)
	ie.QoSFlowsTimedReportList = new(VolumeTimedReportList)
	if err = ie.QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	if err = ie.QoSFlowsTimedReportList.Decode(r); err != nil {
		return
	}
	return
}
