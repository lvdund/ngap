package ies

import "github.com/lvdund/ngap/aper"

type AssistanceDataForRecommendedCells struct {
	RecommendedCellsForPaging *RecommendedCellsForPaging `True,`
	// IEExtensions AssistanceDataForRecommendedCellsExtIEs `False,OPTIONAL`
}

func (ie *AssistanceDataForRecommendedCells) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.RecommendedCellsForPaging != nil {
		if err = ie.RecommendedCellsForPaging.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *AssistanceDataForRecommendedCells) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.RecommendedCellsForPaging = new(RecommendedCellsForPaging)
	if err = ie.RecommendedCellsForPaging.Decode(r); err != nil {
		return
	}
	return
}
