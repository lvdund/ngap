package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionAggregateMaximumBitRate struct {
	PDUSessionAggregateMaximumBitRateDL *BitRate `False,`
	PDUSessionAggregateMaximumBitRateUL *BitRate `False,`
	// IEExtensions PDUSessionAggregateMaximumBitRateExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionAggregateMaximumBitRate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PDUSessionAggregateMaximumBitRateDL != nil {
		if err = ie.PDUSessionAggregateMaximumBitRateDL.Encode(w); err != nil {
			return
		}
	}
	fmt.Printf("log PDUSessionAggregateMaximumBitRate:\t%b\n", aper.GetWriter(*w))
	if ie.PDUSessionAggregateMaximumBitRateUL != nil {
		if err = ie.PDUSessionAggregateMaximumBitRateUL.Encode(w); err != nil {
			return
		}
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
	ie.PDUSessionAggregateMaximumBitRateDL = new(BitRate)
	ie.PDUSessionAggregateMaximumBitRateUL = new(BitRate)
	if err = ie.PDUSessionAggregateMaximumBitRateDL.Decode(r); err != nil {
		return
	}
	if err = ie.PDUSessionAggregateMaximumBitRateUL.Decode(r); err != nil {
		return
	}
	return
}
