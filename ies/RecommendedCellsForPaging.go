package ies

import "github.com/lvdund/ngap/aper"

type RecommendedCellsForPaging struct {
	RecommendedCellList *RecommendedCellList `False,`
	// IEExtensions RecommendedCellsForPagingExtIEs `False,OPTIONAL`
}

func (ie *RecommendedCellsForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.RecommendedCellList != nil {
		if err = ie.RecommendedCellList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *RecommendedCellsForPaging) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.RecommendedCellList = new(RecommendedCellList)
	if err = ie.RecommendedCellList.Decode(r); err != nil {
		return
	}
	return
}
