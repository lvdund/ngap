package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceInformationItem struct {
	PDUSessionID              *PDUSessionID              `False,`
	QosFlowInformationList    *QosFlowInformationList    `False,`
	DRBsToQosFlowsMappingList *DRBsToQosFlowsMappingList `False,OPTIONAL`
	// IEExtensions PDUSessionResourceInformationItemExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DRBsToQosFlowsMappingList != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.PDUSessionID != nil {
		if err = ie.PDUSessionID.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowInformationList != nil {
		if err = ie.QosFlowInformationList.Encode(w); err != nil {
			return
		}
	}
	if ie.DRBsToQosFlowsMappingList != nil {
		if err = ie.DRBsToQosFlowsMappingList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.PDUSessionID = new(PDUSessionID)
	ie.QosFlowInformationList = new(QosFlowInformationList)
	ie.DRBsToQosFlowsMappingList = new(DRBsToQosFlowsMappingList)
	if err = ie.PDUSessionID.Decode(r); err != nil {
		return
	}
	if err = ie.QosFlowInformationList.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DRBsToQosFlowsMappingList.Decode(r); err != nil {
			return
		}
	}
	return
}
