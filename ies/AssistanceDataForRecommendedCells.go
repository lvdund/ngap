package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AssistanceDataForRecommendedCells struct {
	RecommendedCellsForPaging RecommendedCellsForPaging `madatory`
	// IEExtensions *AssistanceDataForRecommendedCellsExtIEs `optional`
}

func (ie *AssistanceDataForRecommendedCells) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.RecommendedCellsForPaging.Encode(w); err != nil {
		err = utils.WrapError("Encode RecommendedCellsForPaging", err)
		return
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
	if err = ie.RecommendedCellsForPaging.Decode(r); err != nil {
		err = utils.WrapError("Read RecommendedCellsForPaging", err)
		return
	}
	return
}
