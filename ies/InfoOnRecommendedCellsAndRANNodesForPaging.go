package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type InfoOnRecommendedCellsAndRANNodesForPaging struct {
	RecommendedCellsForPaging  RecommendedCellsForPaging
	RecommendRANNodesForPaging RecommendedRANNodesForPaging
	// IEExtensions *InfoOnRecommendedCellsAndRANNodesForPagingExtIEs `optional`
}

func (ie *InfoOnRecommendedCellsAndRANNodesForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.RecommendedCellsForPaging.Encode(w); err != nil {
		err = utils.WrapError("Read RecommendedCellsForPaging", err)
		return
	}
	if err = ie.RecommendRANNodesForPaging.Encode(w); err != nil {
		err = utils.WrapError("Read RecommendRANNodesForPaging", err)
		return
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
	if err = ie.RecommendedCellsForPaging.Decode(r); err != nil {
		err = utils.WrapError("Read RecommendedCellsForPaging", err)
		return
	}
	if err = ie.RecommendRANNodesForPaging.Decode(r); err != nil {
		err = utils.WrapError("Read RecommendRANNodesForPaging", err)
		return
	}
	return
}
