package ies

import "github.com/lvdund/ngap/aper"

type UEAggregateMaximumBitRate struct {
	UEAggregateMaximumBitRateDL int64
	UEAggregateMaximumBitRateUL int64
	// IEExtensions *UEAggregateMaximumBitRateExtIEs `optional`
}

func (ie *UEAggregateMaximumBitRate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_UEAggregateMaximumBitRateDL := NewINTEGER(ie.UEAggregateMaximumBitRateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, false)
	if err = tmp_UEAggregateMaximumBitRateDL.Encode(w); err != nil {
		return
	}
	tmp_UEAggregateMaximumBitRateUL := NewINTEGER(ie.UEAggregateMaximumBitRateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, false)
	if err = tmp_UEAggregateMaximumBitRateUL.Encode(w); err != nil {
		return
	}
	return
}
func (ie *UEAggregateMaximumBitRate) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_UEAggregateMaximumBitRateDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: false,
	}
	if err = tmp_UEAggregateMaximumBitRateDL.Decode(r); err != nil {
		return
	}
	ie.UEAggregateMaximumBitRateDL = int64(tmp_UEAggregateMaximumBitRateDL.Value)
	tmp_UEAggregateMaximumBitRateUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: false,
	}
	if err = tmp_UEAggregateMaximumBitRateUL.Decode(r); err != nil {
		return
	}
	ie.UEAggregateMaximumBitRateUL = int64(tmp_UEAggregateMaximumBitRateUL.Value)
	return
}
