package ies

import "github.com/lvdund/ngap/aper"

type ServiceAreaInformation struct {
	Value []*ServiceAreaInformationItem `False,1,16`
}

func (ie *ServiceAreaInformation) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*ServiceAreaInformationItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *ServiceAreaInformation) Decode(r *aper.AperReader) (err error) {
	var newItems []*ServiceAreaInformationItem
	newItems, err = aper.ReadSequenceOfEx[*ServiceAreaInformationItem](func() *ServiceAreaInformationItem { return new(ServiceAreaInformationItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*ServiceAreaInformationItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
