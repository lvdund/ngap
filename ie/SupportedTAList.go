package ie

import "ngap/aper"

// Need to import "github.com/free5gc/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct SupportedTAList */
/* SupportedTAItem */
type SupportedTAList struct {
	List []*SupportedTaItem `aper:"valueExt,sizeLB:1,sizeUB:256"`
}

func (ie *SupportedTAList) Decode(r *aper.AperReader) error {
	newItems, err := aper.ReadSequenceOfEx[*SupportedTaItem](func() *SupportedTaItem {
		return new(SupportedTaItem)
	}, r, &aper.Constraint{Lb: 1, Ub: 256}, true)
	ie.List = []*SupportedTaItem{}
	ie.List = append(ie.List, newItems...)
	return err
}

func (ie *SupportedTAList) Encode(r *aper.AperWriter) (err error) {
	if err := aper.WriteSequenceOf[*SupportedTaItem](ie.List, r, &aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
		return err
	}
	return nil
}
