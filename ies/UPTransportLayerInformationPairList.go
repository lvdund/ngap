package ies

import "github.com/lvdund/ngap/aper"

type UPTransportLayerInformationPairList struct {
	Value []*UPTransportLayerInformationPairItem `False,1,3`
}

func (ie *UPTransportLayerInformationPairList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*UPTransportLayerInformationPairItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return
	}
	return
}
func (ie *UPTransportLayerInformationPairList) Decode(r *aper.AperReader) (err error) {
	var newItems []*UPTransportLayerInformationPairItem
	newItems, err = aper.ReadSequenceOfEx[*UPTransportLayerInformationPairItem](func() *UPTransportLayerInformationPairItem { return new(UPTransportLayerInformationPairItem) }, r, &aper.Constraint{Lb: 1, Ub: 3}, false)
	ie.Value = []*UPTransportLayerInformationPairItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
