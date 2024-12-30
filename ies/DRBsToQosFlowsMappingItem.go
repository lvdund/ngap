package ies

import "github.com/lvdund/ngap/aper"

type DRBsToQosFlowsMappingItem struct {
	DRBID                 int64
	AssociatedQosFlowList []AssociatedQosFlowItem
	// IEExtensions  *DRBsToQosFlowsMappingItemExtIEs
}

func (ie *DRBsToQosFlowsMappingItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_DRBID.Encode(w); err != nil {
		return
	}
	if len(ie.AssociatedQosFlowList) > 0 {
		tmp := Sequence[*AssociatedQosFlowItem]{
			Value: []*AssociatedQosFlowItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.AssociatedQosFlowList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *DRBsToQosFlowsMappingItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_DRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: false,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
	tmp_AssociatedQosFlowList := Sequence[*AssociatedQosFlowItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	if err = tmp_AssociatedQosFlowList.Decode(r); err != nil {
		return
	}
	ie.AssociatedQosFlowList = []AssociatedQosFlowItem{}
	for _, i := range tmp_AssociatedQosFlowList.Value {
		ie.AssociatedQosFlowList = append(ie.AssociatedQosFlowList, *i)
	}
	return
}
