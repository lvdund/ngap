package ies

import "github.com/lvdund/ngap/aper"

type UEHistoryInformation struct {
	Value []*LastVisitedCellItem `False,1,16`
}

func (ie *UEHistoryInformation) Encode(w *aper.AperWriter) (err error) {
	if err = aper.WriteSequenceOf[*LastVisitedCellItem](ie.Value, w, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return
	}
	return
}
func (ie *UEHistoryInformation) Decode(r *aper.AperReader) (err error) {
	var newItems []*LastVisitedCellItem
	newItems, err = aper.ReadSequenceOfEx[*LastVisitedCellItem](func() *LastVisitedCellItem { return new(LastVisitedCellItem) }, r, &aper.Constraint{Lb: 1, Ub: 16}, false)
	ie.Value = []*LastVisitedCellItem{}
	ie.Value = append(ie.Value, newItems...)
	return
}
