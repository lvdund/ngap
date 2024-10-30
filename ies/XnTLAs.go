package ies

import "github.com/lvdund/ngap/aper"

type XnTLAs struct {
	Value []*TransportLayerAddress `False,1,2`
}

func (ie *XnTLAs) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*TransportLayerAddress](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return
	}
	return
}
func (ie *XnTLAs) Decode(r *aper.AperReader) (err error) {
	var newItems []*TransportLayerAddress
	newItems, err = aper.ReadSequenceOfEx[*TransportLayerAddress](func() *TransportLayerAddress { return new(TransportLayerAddress) }, r, &aper.Constraint{Lb: 1, Ub: 2}, false)
	ie.Value = []*TransportLayerAddress{}
	ie.Value = append(ie.Value, newItems...)
	return
}
