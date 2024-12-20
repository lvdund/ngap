package ies

import "github.com/lvdund/ngap/aper"

type SecurityContext struct {
	NextHopChainingCount *NextHopChainingCount `False,`
	NextHopNH            *SecurityKey          `False,`
	// IEExtensions SecurityContextExtIEs `False,OPTIONAL`
}

func (ie *SecurityContext) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.NextHopChainingCount != nil {
		if err = ie.NextHopChainingCount.Encode(w); err != nil {
			return
		}
	}
	if ie.NextHopNH != nil {
		if err = ie.NextHopNH.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *SecurityContext) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.NextHopChainingCount = new(NextHopChainingCount)
	ie.NextHopNH = new(SecurityKey)
	if err = ie.NextHopChainingCount.Decode(r); err != nil {
		return
	}
	if err = ie.NextHopNH.Decode(r); err != nil {
		return
	}
	return
}
