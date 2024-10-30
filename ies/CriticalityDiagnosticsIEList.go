package ies

import "github.com/lvdund/ngap/aper"

type CriticalityDiagnosticsIEList struct {
	Value []*CriticalityDiagnosticsIEItem `False,1,256`
}

func (ie *CriticalityDiagnosticsIEList) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*CriticalityDiagnosticsIEItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return
	}
	return
}
func (ie *CriticalityDiagnosticsIEList) Decode(r *aper.AperReader) (err error) {
	var newItems []*CriticalityDiagnosticsIEItem
	newItems, err = aper.ReadSequenceOfEx[*CriticalityDiagnosticsIEItem](func() *CriticalityDiagnosticsIEItem { return new(CriticalityDiagnosticsIEItem) }, r, &aper.Constraint{Lb: 1, Ub: 256}, false)
	ie.Value = []*CriticalityDiagnosticsIEItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
