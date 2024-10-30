package ies

import "github.com/lvdund/ngap/aper"

type InfoOnRecommendedCellsAndRANNodesForPaging struct {
	RecommendedCellsForPaging  *RecommendedCellsForPaging    `True,`
	RecommendRANNodesForPaging *RecommendedRANNodesForPaging `True,`
	// IEExtensions InfoOnRecommendedCellsAndRANNodesForPagingExtIEs `False,OPTIONAL`
}

func (ie *InfoOnRecommendedCellsAndRANNodesForPaging) Encode(w *aper.AperWriter) (err error) {
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
	if ie.RecommendRANNodesForPaging != nil {
		if err = ie.RecommendRANNodesForPaging.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *InfoOnRecommendedCellsAndRANNodesForPaging) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.RecommendedCellsForPaging = new(RecommendedCellsForPaging)
	ie.RecommendRANNodesForPaging = new(RecommendedRANNodesForPaging)
	if err = ie.RecommendedCellsForPaging.Decode(r); err != nil {
		return
	}
	if err = ie.RecommendRANNodesForPaging.Decode(r); err != nil {
		return
	}
	return
}
