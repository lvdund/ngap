package ies

import "github.com/lvdund/ngap/aper"

type RATRestrictionsItem struct {
	PLMNIdentity              *PLMNIdentity              `False,`
	RATRestrictionInformation *RATRestrictionInformation `False,`
	// IEExtensions RATRestrictionsItemExtIEs `False,OPTIONAL`
}

func (ie *RATRestrictionsItem) Encode(w *aper.AperWriter) (err error) {
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
	if ie.RATRestrictionInformation != nil {
		if err = ie.RATRestrictionInformation.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *RATRestrictionsItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.RATRestrictionInformation = new(RATRestrictionInformation)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.RATRestrictionInformation.Decode(r); err != nil {
		return
	}
	return
}
