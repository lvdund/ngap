package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AssistanceDataForPaging struct {
	AssistanceDataForRecommendedCells *AssistanceDataForRecommendedCells `optional`
	PagingAttemptInformation          *PagingAttemptInformation          `optional`
	// IEExtensions *AssistanceDataForPagingExtIEs `optional`
}

func (ie *AssistanceDataForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AssistanceDataForRecommendedCells != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.PagingAttemptInformation != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.AssistanceDataForRecommendedCells != nil {
		if err = ie.AssistanceDataForRecommendedCells.Encode(w); err != nil {
			err = utils.WrapError("Read AssistanceDataForRecommendedCells", err)
			return
		}
	}
	if ie.PagingAttemptInformation != nil {
		if err = ie.PagingAttemptInformation.Encode(w); err != nil {
			err = utils.WrapError("Read PagingAttemptInformation", err)
			return
		}
	}
	return
}
func (ie *AssistanceDataForPaging) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.AssistanceDataForRecommendedCells.Decode(r); err != nil {
			err = utils.WrapError("Read AssistanceDataForRecommendedCells", err)
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.PagingAttemptInformation.Decode(r); err != nil {
			err = utils.WrapError("Read PagingAttemptInformation", err)
			return
		}
	}
	return
}
