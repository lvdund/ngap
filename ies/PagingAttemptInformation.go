package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PagingAttemptInformation struct {
	PagingAttemptCount             int64                `lb:1,ub:16,madatory,valExt`
	IntendedNumberOfPagingAttempts int64                `lb:1,ub:16,madatory,valExt`
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
	tmp_PagingAttemptCount := NewINTEGER(ie.PagingAttemptCount, aper.Constraint{Lb: 1, Ub: 16}, true)
	if err = tmp_PagingAttemptCount.Encode(w); err != nil {
		err = fmt.Errorf("Encode PagingAttemptCount: %w", err)
		return
	}
	tmp_IntendedNumberOfPagingAttempts := NewINTEGER(ie.IntendedNumberOfPagingAttempts, aper.Constraint{Lb: 1, Ub: 16}, true)
	if err = tmp_IntendedNumberOfPagingAttempts.Encode(w); err != nil {
		err = fmt.Errorf("Encode IntendedNumberOfPagingAttempts: %w", err)
		return
	}
	if ie.NextPagingAreaScope != nil {
		if err = ie.NextPagingAreaScope.Encode(w); err != nil {
			err = fmt.Errorf("Encode NextPagingAreaScope: %w", err)
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
		ext: true,
	}
	if err = tmp_PagingAttemptCount.Decode(r); err != nil {
		err = fmt.Errorf("Read PagingAttemptCount: %w", err)
		return
	}
	ie.PagingAttemptCount = int64(tmp_PagingAttemptCount.Value)
	tmp_IntendedNumberOfPagingAttempts := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 16},
		ext: true,
	}
	if err = tmp_IntendedNumberOfPagingAttempts.Decode(r); err != nil {
		err = fmt.Errorf("Read IntendedNumberOfPagingAttempts: %w", err)
		return
	}
	ie.IntendedNumberOfPagingAttempts = int64(tmp_IntendedNumberOfPagingAttempts.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(NextPagingAreaScope)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read NextPagingAreaScope: %w", err)
			return
		}
		ie.NextPagingAreaScope = tmp
	}
	return
}
