package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PagingAttemptInformation struct {
	PagingAttemptCount             int64
	IntendedNumberOfPagingAttempts int64
	NextPagingAreaScope            *NextPagingAreaScope `optional`
	// IEExtensions *PagingAttemptInformationExtIEs `optional`
}

func (ie *PagingAttemptInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NextPagingAreaScope != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_PagingAttemptCount := NewINTEGER(ie.PagingAttemptCount, aper.Constraint{Lb: 1, Ub: 16}, false)
	if err = tmp_PagingAttemptCount.Encode(w); err != nil {
		err = utils.WrapError("Read PagingAttemptCount", err)
		return
	}
	tmp_IntendedNumberOfPagingAttempts := NewINTEGER(ie.IntendedNumberOfPagingAttempts, aper.Constraint{Lb: 1, Ub: 16}, false)
	if err = tmp_IntendedNumberOfPagingAttempts.Encode(w); err != nil {
		err = utils.WrapError("Read IntendedNumberOfPagingAttempts", err)
		return
	}
	if ie.NextPagingAreaScope != nil {
		if err = ie.NextPagingAreaScope.Encode(w); err != nil {
			err = utils.WrapError("Read NextPagingAreaScope", err)
			return
		}
	}
	return
}
func (ie *PagingAttemptInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_PagingAttemptCount := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 16},
		ext: false,
	}
	if err = tmp_PagingAttemptCount.Decode(r); err != nil {
		err = utils.WrapError("Read PagingAttemptCount", err)
		return
	}
	ie.PagingAttemptCount = int64(tmp_PagingAttemptCount.Value)
	tmp_IntendedNumberOfPagingAttempts := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 16},
		ext: false,
	}
	if err = tmp_IntendedNumberOfPagingAttempts.Decode(r); err != nil {
		err = utils.WrapError("Read IntendedNumberOfPagingAttempts", err)
		return
	}
	ie.IntendedNumberOfPagingAttempts = int64(tmp_IntendedNumberOfPagingAttempts.Value)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.NextPagingAreaScope.Decode(r); err != nil {
			err = utils.WrapError("Read NextPagingAreaScope", err)
			return
		}
	}
	return
}
