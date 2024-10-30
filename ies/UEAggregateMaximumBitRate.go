package ies

import "github.com/lvdund/ngap/aper"

type UEAggregateMaximumBitRate struct {
	UEAggregateMaximumBitRateDL *BitRate `False,`
	UEAggregateMaximumBitRateUL *BitRate `False,`
	// IEExtensions UEAggregateMaximumBitRateExtIEs `False,OPTIONAL`
}

func (ie *UEAggregateMaximumBitRate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.UEAggregateMaximumBitRateDL != nil {
		if err = ie.UEAggregateMaximumBitRateDL.Encode(w); err != nil {
			return
		}
	}
	if ie.UEAggregateMaximumBitRateUL != nil {
		if err = ie.UEAggregateMaximumBitRateUL.Encode(w); err != nil {
			return
		}
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
	ie.UEAggregateMaximumBitRateDL = new(BitRate)
	ie.UEAggregateMaximumBitRateUL = new(BitRate)
	if err = ie.UEAggregateMaximumBitRateDL.Decode(r); err != nil {
		return
	}
	if err = ie.UEAggregateMaximumBitRateUL.Decode(r); err != nil {
		return
	}
	return
}
