package ies

import "github.com/lvdund/ngap/aper"

type TargetNGRANNodeToSourceNGRANNodeTransparentContainer struct {
	RRCContainer []byte
	// IEExtensions *TargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs `optional`
}

func (ie *TargetNGRANNodeToSourceNGRANNodeTransparentContainer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_RRCContainer := NewOCTETSTRING(ie.RRCContainer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_RRCContainer.Encode(w); err != nil {
		return
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
	tmp_RRCContainer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_RRCContainer.Decode(r); err != nil {
		return
	}
	ie.RRCContainer = tmp_RRCContainer.Value
	return
}
