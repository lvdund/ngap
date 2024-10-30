package ies

import "github.com/lvdund/ngap/aper"

type TargetNGRANNodeToSourceNGRANNodeTransparentContainer struct {
	RRCContainer *RRCContainer `False,`
	// IEExtensions TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs `False,OPTIONAL`
}

func (ie *TargetNGRANNodeToSourceNGRANNodeTransparentContainer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.RRCContainer != nil {
		if err = ie.RRCContainer.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TargetNGRANNodeToSourceNGRANNodeTransparentContainer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.RRCContainer = new(RRCContainer)
	if err = ie.RRCContainer.Decode(r); err != nil {
		return
	}
	return
}
