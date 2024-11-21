package ies

import "github.com/lvdund/ngap/aper"

type PLMNSupportItem struct {
	PLMNIdentity     *PLMNIdentity     `False,`
	SliceSupportList *SliceSupportList `False,`
	// IEExtensions PLMNSupportItemExtIEs `False,OPTIONAL`
}

func (ie *PLMNSupportItem) Encode(w *aper.AperWriter) (err error) {
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
	if ie.SliceSupportList != nil {
		if err = ie.SliceSupportList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PLMNSupportItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.SliceSupportList = new(SliceSupportList)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.SliceSupportList.Decode(r); err != nil {
		return
	}
	return
}
