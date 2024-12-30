package ies

import "github.com/lvdund/ngap/aper"

type OverloadStartNSSAIItem struct {
	SliceOverloadList                   []SliceOverloadItem
	SliceOverloadResponse               *OverloadResponse
	SliceTrafficLoadReductionIndication *int64
	// IEExtensions  *OverloadStartNSSAIItemExtIEs
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
			return
		}
	}
	if ie.SliceOverloadResponse != nil {
		if err = ie.SliceOverloadResponse.Encode(w); err != nil {
			return
		}
	}
	if ie.SliceTrafficLoadReductionIndication != nil {
		tmp_SliceTrafficLoadReductionIndication := NewINTEGER(*ie.SliceTrafficLoadReductionIndication, aper.Constraint{Lb: 1, Ub: 99}, false)
		if err = tmp_SliceTrafficLoadReductionIndication.Encode(w); err != nil {
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
	if err = tmp_SliceOverloadList.Decode(r); err != nil {
		return
	}
	ie.SliceOverloadList = []SliceOverloadItem{}
	for _, i := range tmp_SliceOverloadList.Value {
		ie.SliceOverloadList = append(ie.SliceOverloadList, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.SliceOverloadResponse.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_SliceTrafficLoadReductionIndication := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 99},
			ext: false,
		}
		if err = tmp_SliceTrafficLoadReductionIndication.Decode(r); err != nil {
			return
		}
		*ie.SliceTrafficLoadReductionIndication = int64(tmp_SliceTrafficLoadReductionIndication.Value)
	}
	return
}
