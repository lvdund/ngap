package ies

import "github.com/lvdund/ngap/aper"

type XnExtTLAs struct {
	Value []*XnExtTLAItem `False,1,16`
}

func (ie *XnExtTLAs) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*XnExtTLAItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *XnExtTLAs) Decode(r *aper.AperReader) (err error) {
	var newItems []*XnExtTLAItem
	newItems, err = aper.ReadSequenceOfEx[*XnExtTLAItem](func() *XnExtTLAItem { return new(XnExtTLAItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*XnExtTLAItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
