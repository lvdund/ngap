package ies

import "github.com/lvdund/ngap/aper"

type CNTypeRestrictionsForEquivalentItem struct {
	PlmnIdentity *PLMNIdentity `False,`
	// IEExtensions CNTypeRestrictionsForEquivalentItemExtIEs `False,OPTIONAL`
}

func (ie *CNTypeRestrictionsForEquivalentItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PlmnIdentity != nil {
		if err = ie.PlmnIdentity.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CNTypeRestrictionsForEquivalentItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PlmnIdentity = new(PLMNIdentity)
	if err = ie.PlmnIdentity.Decode(r); err != nil {
		return
	}
	return
}
