package ies

import "github.com/lvdund/ngap/aper"

type DataForwardingResponseDRBList struct {
	Value []*DataForwardingResponseDRBItem `False,1,32`
}

func (ie *DataForwardingResponseDRBList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*DataForwardingResponseDRBItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return
	}
	return
}
func (ie *DataForwardingResponseDRBList) Decode(r *aper.AperReader) (err error) {
	var newItems []*DataForwardingResponseDRBItem
	newItems, err = aper.ReadSequenceOfEx[*DataForwardingResponseDRBItem](func() *DataForwardingResponseDRBItem { return new(DataForwardingResponseDRBItem) }, r, &aper.Constraint{Lb: 1, Ub: 32}, false)
	ie.Value = []*DataForwardingResponseDRBItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
