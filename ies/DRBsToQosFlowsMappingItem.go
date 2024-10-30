package ies

import "github.com/lvdund/ngap/aper"

type DRBsToQosFlowsMappingItem struct {
	DRBID                 *DRBID                 `False,`
	AssociatedQosFlowList *AssociatedQosFlowList `False,`
	// IEExtensions DRBsToQosFlowsMappingItemExtIEs `False,OPTIONAL`
}

func (ie *DRBsToQosFlowsMappingItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.DRBID != nil {
		if err = ie.DRBID.Encode(w); err != nil {
			return
		}
	}
	if ie.AssociatedQosFlowList != nil {
		if err = ie.AssociatedQosFlowList.Encode(w); err != nil {
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
	ie.DRBID = new(DRBID)
	ie.AssociatedQosFlowList = new(AssociatedQosFlowList)
	if err = ie.DRBID.Decode(r); err != nil {
		return
	}
	if err = ie.AssociatedQosFlowList.Decode(r); err != nil {
		return
	}
	return
}
