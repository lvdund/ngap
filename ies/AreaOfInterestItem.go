package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type AreaOfInterestItem struct {
	AreaOfInterest               AreaOfInterest `madatory`
	LocationReportingReferenceID int64          `lb:1,ub:64,madatory,valExt`
	// IEExtensions *AreaOfInterestItemExtIEs `optional`
}

func (ie *AreaOfInterestItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.AreaOfInterest.Encode(w); err != nil {
		err = fmt.Errorf("Encode AreaOfInterest: %w", err)
		return
	}
	tmp_LocationReportingReferenceID := NewINTEGER(ie.LocationReportingReferenceID, aper.Constraint{Lb: 1, Ub: 64}, true)
	if err = tmp_LocationReportingReferenceID.Encode(w); err != nil {
		err = fmt.Errorf("Encode LocationReportingReferenceID: %w", err)
		return
	}
	return
}
func (ie *AreaOfInterestItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.AreaOfInterest.Decode(r); err != nil {
		err = fmt.Errorf("Read AreaOfInterest: %w", err)
		return
	}
	tmp_LocationReportingReferenceID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 64},
		ext: true,
	}
	if err = tmp_LocationReportingReferenceID.Decode(r); err != nil {
		err = fmt.Errorf("Read LocationReportingReferenceID: %w", err)
		return
	}
	ie.LocationReportingReferenceID = int64(tmp_LocationReportingReferenceID.Value)
	return
}
