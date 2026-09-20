package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type LocationReportingRequestType struct {
	EventType                                 EventType            `madatory`
	ReportArea                                ReportArea           `madatory`
	AreaOfInterestList                        []AreaOfInterestItem `lb:1,ub:maxnoofAoI,optional`
	LocationReportingReferenceIDToBeCancelled *int64               `lb:1,ub:64,optional,valExt`
	// IEExtensions *LocationReportingRequestTypeExtIEs `optional`
}

func (ie *LocationReportingRequestType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AreaOfInterestList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.LocationReportingReferenceIDToBeCancelled != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.EventType.Encode(w); err != nil {
		err = fmt.Errorf("Encode EventType: %w", err)
		return
	}
	if err = ie.ReportArea.Encode(w); err != nil {
		err = fmt.Errorf("Encode ReportArea: %w", err)
		return
	}
	if len(ie.AreaOfInterestList) > 0 {
		tmp := Sequence[*AreaOfInterestItem]{
			Value: []*AreaOfInterestItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofAoI},
			ext:   false,
		}
		for _, i := range ie.AreaOfInterestList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode AreaOfInterestList: %w", err)
			return
		}
	}
	if ie.LocationReportingReferenceIDToBeCancelled != nil {
		tmp_LocationReportingReferenceIDToBeCancelled := NewINTEGER(*ie.LocationReportingReferenceIDToBeCancelled, aper.Constraint{Lb: 1, Ub: 64}, true)
		if err = tmp_LocationReportingReferenceIDToBeCancelled.Encode(w); err != nil {
			err = fmt.Errorf("Encode LocationReportingReferenceIDToBeCancelled: %w", err)
			return
		}
	}
	return
}
func (ie *LocationReportingRequestType) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.EventType.Decode(r); err != nil {
		err = fmt.Errorf("Read EventType: %w", err)
		return
	}
	if err = ie.ReportArea.Decode(r); err != nil {
		err = fmt.Errorf("Read ReportArea: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_AreaOfInterestList := Sequence[*AreaOfInterestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAoI},
			ext: false,
		}
		fn := func() *AreaOfInterestItem { return new(AreaOfInterestItem) }
		if err = tmp_AreaOfInterestList.Decode(r, fn); err != nil {
			err = fmt.Errorf("Read AreaOfInterestList: %w", err)
			return
		}
		ie.AreaOfInterestList = []AreaOfInterestItem{}
		for _, i := range tmp_AreaOfInterestList.Value {
			ie.AreaOfInterestList = append(ie.AreaOfInterestList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_LocationReportingReferenceIDToBeCancelled := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 64},
			ext: true,
		}
		if err = tmp_LocationReportingReferenceIDToBeCancelled.Decode(r); err != nil {
			err = fmt.Errorf("Read LocationReportingReferenceIDToBeCancelled: %w", err)
			return
		}
		ie.LocationReportingReferenceIDToBeCancelled = (*int64)(&tmp_LocationReportingReferenceIDToBeCancelled.Value)
	}
	return
}
