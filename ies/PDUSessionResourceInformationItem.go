package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceInformationItem struct {
	PDUSessionID              int64                       `lb:0,ub:255,madatory`
	QosFlowInformationList    []QosFlowInformationItem    `lb:1,ub:maxnoofQosFlows,madatory`
	DRBsToQosFlowsMappingList []DRBsToQosFlowsMappingItem `lb:1,ub:maxnoofDRBs,optional`
	// IEExtensions *PDUSessionResourceInformationItemExtIEs `optional`
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
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		err = utils.WrapError("Encode PDUSessionID", err)
		return
	}
	if len(ie.QosFlowInformationList) > 0 {
		tmp := Sequence[*QosFlowInformationItem]{
			Value: []*QosFlowInformationItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowInformationList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode QosFlowInformationList", err)
			return
		}
	} else {
		err = utils.WrapError("QosFlowInformationList is nil", err)
		return
	}
	if len(ie.DRBsToQosFlowsMappingList) > 0 {
		tmp := Sequence[*DRBsToQosFlowsMappingItem]{
			Value: []*DRBsToQosFlowsMappingItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext:   false,
		}
		for _, i := range ie.DRBsToQosFlowsMappingList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode DRBsToQosFlowsMappingList", err)
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
	tmp_PDUSessionID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_PDUSessionID.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionID", err)
		return
	}
	ie.PDUSessionID = int64(tmp_PDUSessionID.Value)
	tmp_QosFlowInformationList := Sequence[*QosFlowInformationItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	fn := func() *QosFlowInformationItem { return new(QosFlowInformationItem) }
	if err = tmp_QosFlowInformationList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read QosFlowInformationList", err)
		return
	}
	ie.QosFlowInformationList = []QosFlowInformationItem{}
	for _, i := range tmp_QosFlowInformationList.Value {
		ie.QosFlowInformationList = append(ie.QosFlowInformationList, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_DRBsToQosFlowsMappingList := Sequence[*DRBsToQosFlowsMappingItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DRBsToQosFlowsMappingItem { return new(DRBsToQosFlowsMappingItem) }
		if err = tmp_DRBsToQosFlowsMappingList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read DRBsToQosFlowsMappingList", err)
			return
		}
		ie.DRBsToQosFlowsMappingList = []DRBsToQosFlowsMappingItem{}
		for _, i := range tmp_DRBsToQosFlowsMappingList.Value {
			ie.DRBsToQosFlowsMappingList = append(ie.DRBsToQosFlowsMappingList, *i)
		}
	}
	return
}
