package ies

import "github.com/lvdund/ngap/aper"

type BroadcastPLMNItem struct {
	PLMNIdentity        *PLMNIdentity     `False,`
	TAISliceSupportList *SliceSupportList `False,`
	// IEExtensions BroadcastPLMNItemExtIEs `False,OPTIONAL`
}

func (ie *BroadcastPLMNItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PLMNIdentity != nil {
		if err = ie.PLMNIdentity.Encode(w); err != nil {
			return
		}
	}
	if ie.TAISliceSupportList != nil {
		if err = ie.TAISliceSupportList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *BroadcastPLMNItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.TAISliceSupportList = new(SliceSupportList)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.TAISliceSupportList.Decode(r); err != nil {
		return
	}
	return
}
