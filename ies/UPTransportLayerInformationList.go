package ies

import "github.com/lvdund/ngap/aper"

type UPTransportLayerInformationList struct {
	Value []*UPTransportLayerInformationItem `False,1,3`
}

func (ie *UPTransportLayerInformationList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*UPTransportLayerInformationItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return
	}
	return
}
func (ie *UPTransportLayerInformationList) Decode(r *aper.AperReader) (err error) {
	var newItems []*UPTransportLayerInformationItem
	newItems, err = aper.ReadSequenceOfEx[*UPTransportLayerInformationItem](func() *UPTransportLayerInformationItem { return new(UPTransportLayerInformationItem) }, r, &aper.Constraint{Lb: 1, Ub: 3}, false)
	ie.Value = []*UPTransportLayerInformationItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
