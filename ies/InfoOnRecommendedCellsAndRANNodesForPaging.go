package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type InfoOnRecommendedCellsAndRANNodesForPaging struct {
	RecommendedCellsForPaging  RecommendedCellsForPaging    `madatory`
	RecommendRANNodesForPaging RecommendedRANNodesForPaging `madatory`
	// IEExtensions *InfoOnRecommendedCellsAndRANNodesForPagingExtIEs `optional`
}

func (ie *InfoOnRecommendedCellsAndRANNodesForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.RecommendedCellsForPaging.Encode(w); err != nil {
		err = fmt.Errorf("Encode RecommendedCellsForPaging: %w", err)
		return
	}
	if err = ie.RecommendRANNodesForPaging.Encode(w); err != nil {
		err = fmt.Errorf("Encode RecommendRANNodesForPaging: %w", err)
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
		err = fmt.Errorf("Read RecommendedCellsForPaging: %w", err)
		return
	}
	if err = ie.RecommendRANNodesForPaging.Decode(r); err != nil {
		err = fmt.Errorf("Read RecommendRANNodesForPaging: %w", err)
		return
	}
	return
}
