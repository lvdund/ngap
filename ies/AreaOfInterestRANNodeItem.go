package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterestRANNodeItem struct {
	GlobalRANNodeID *GlobalRANNodeID `False,`
	// IEExtensions AreaOfInterestRANNodeItemExtIEs `False,OPTIONAL`
}

func (ie *AreaOfInterestRANNodeItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.GlobalRANNodeID != nil {
		if err = ie.GlobalRANNodeID.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *AreaOfInterestRANNodeItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.GlobalRANNodeID = new(GlobalRANNodeID)
	if err = ie.GlobalRANNodeID.Decode(r); err != nil {
		return
	}
	return
}
