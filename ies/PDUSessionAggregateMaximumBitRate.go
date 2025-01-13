package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionAggregateMaximumBitRate struct {
	PDUSessionAggregateMaximumBitRateDL int64 `lb:0,ub:4000000000000,madatory,valExt`
	PDUSessionAggregateMaximumBitRateUL int64 `lb:0,ub:4000000000000,madatory,valExt`
	// IEExtensions *PDUSessionAggregateMaximumBitRateExtIEs `optional`
}

func (ie *PDUSessionAggregateMaximumBitRate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionAggregateMaximumBitRateDL := NewINTEGER(ie.PDUSessionAggregateMaximumBitRateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_PDUSessionAggregateMaximumBitRateDL.Encode(w); err != nil {
		err = utils.WrapError("Encode PDUSessionAggregateMaximumBitRateDL", err)
		return
	}
	tmp_PDUSessionAggregateMaximumBitRateUL := NewINTEGER(ie.PDUSessionAggregateMaximumBitRateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, true)
	if err = tmp_PDUSessionAggregateMaximumBitRateUL.Encode(w); err != nil {
		err = utils.WrapError("Encode PDUSessionAggregateMaximumBitRateUL", err)
		return
	}
	return
}
func (ie *PDUSessionAggregateMaximumBitRate) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PDUSessionAggregateMaximumBitRateDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_PDUSessionAggregateMaximumBitRateDL.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionAggregateMaximumBitRateDL", err)
		return
	}
	ie.PDUSessionAggregateMaximumBitRateDL = int64(tmp_PDUSessionAggregateMaximumBitRateDL.Value)
	tmp_PDUSessionAggregateMaximumBitRateUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: true,
	}
	if err = tmp_PDUSessionAggregateMaximumBitRateUL.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionAggregateMaximumBitRateUL", err)
		return
	}
	ie.PDUSessionAggregateMaximumBitRateUL = int64(tmp_PDUSessionAggregateMaximumBitRateUL.Value)
	return
}
