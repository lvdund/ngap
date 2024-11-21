package ies

import "github.com/lvdund/ngap/aper"

type OverloadStartNSSAIItem struct {
	SliceOverloadList                   *SliceOverloadList              `False,`
	SliceOverloadResponse               *OverloadResponse               `False,OPTIONAL`
	SliceTrafficLoadReductionIndication *TrafficLoadReductionIndication `False,OPTIONAL`
	// IEExtensions OverloadStartNSSAIItemExtIEs `False,OPTIONAL`
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
	if ie.SliceOverloadList != nil {
		if err = ie.SliceOverloadList.Encode(w); err != nil {
			return
		}
	}
	if ie.SliceOverloadResponse != nil {
		if err = ie.SliceOverloadResponse.Encode(w); err != nil {
			return
		}
	}
	if ie.SliceTrafficLoadReductionIndication != nil {
		if err = ie.SliceTrafficLoadReductionIndication.Encode(w); err != nil {
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
	ie.SliceOverloadList = new(SliceOverloadList)
	ie.SliceOverloadResponse = new(OverloadResponse)
	ie.SliceTrafficLoadReductionIndication = new(TrafficLoadReductionIndication)
	if err = ie.SliceOverloadList.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.SliceOverloadResponse.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.SliceTrafficLoadReductionIndication.Decode(r); err != nil {
			return
		}
	}
	return
}
