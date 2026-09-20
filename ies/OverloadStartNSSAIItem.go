package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type OverloadStartNSSAIItem struct {
	SliceOverloadList                   []SliceOverloadItem `lb:1,ub:maxnoofSliceItems,madatory`
	SliceOverloadResponse               *OverloadResponse   `optional`
	SliceTrafficLoadReductionIndication *int64              `lb:1,ub:99,optional`
	// IEExtensions *OverloadStartNSSAIItemExtIEs `optional`
}

func (ie *OverloadStartNSSAIItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SliceOverloadResponse != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SliceTrafficLoadReductionIndication != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if len(ie.SliceOverloadList) > 0 {
		tmp := Sequence[*SliceOverloadItem]{
			Value: []*SliceOverloadItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext:   false,
		}
		for _, i := range ie.SliceOverloadList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode SliceOverloadList: %w", err)
			return
		}
	} else {
		if err != nil {
			err = fmt.Errorf("SliceOverloadList is nil: %w", err)
		}
		return
	}
	if ie.SliceOverloadResponse != nil {
		if err = ie.SliceOverloadResponse.Encode(w); err != nil {
			err = fmt.Errorf("Encode SliceOverloadResponse: %w", err)
			return
		}
	}
	if ie.SliceTrafficLoadReductionIndication != nil {
		tmp_SliceTrafficLoadReductionIndication := NewINTEGER(*ie.SliceTrafficLoadReductionIndication, aper.Constraint{Lb: 1, Ub: 99}, false)
		if err = tmp_SliceTrafficLoadReductionIndication.Encode(w); err != nil {
			err = fmt.Errorf("Encode SliceTrafficLoadReductionIndication: %w", err)
			return
		}
	}
	return
}
func (ie *OverloadStartNSSAIItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	tmp_SliceOverloadList := Sequence[*SliceOverloadItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
		ext: false,
	}
	fn := func() *SliceOverloadItem { return new(SliceOverloadItem) }
	if err = tmp_SliceOverloadList.Decode(r, fn); err != nil {
		err = fmt.Errorf("Read SliceOverloadList: %w", err)
		return
	}
	ie.SliceOverloadList = []SliceOverloadItem{}
	for _, i := range tmp_SliceOverloadList.Value {
		ie.SliceOverloadList = append(ie.SliceOverloadList, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(OverloadResponse)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read SliceOverloadResponse: %w", err)
			return
		}
		ie.SliceOverloadResponse = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_SliceTrafficLoadReductionIndication := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 99},
			ext: false,
		}
		if err = tmp_SliceTrafficLoadReductionIndication.Decode(r); err != nil {
			err = fmt.Errorf("Read SliceTrafficLoadReductionIndication: %w", err)
			return
		}
		ie.SliceTrafficLoadReductionIndication = (*int64)(&tmp_SliceTrafficLoadReductionIndication.Value)
	}
	return
}
