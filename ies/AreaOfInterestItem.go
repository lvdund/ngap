package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type AreaOfInterestItem struct {
	AreaOfInterest               AreaOfInterest
	LocationReportingReferenceID int64
	// IEExtensions *AreaOfInterestItemExtIEs `optional`
}

func (ie *AreaOfInterestItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.AreaOfInterest.Encode(w); err != nil {
		err = utils.WrapError("Read AreaOfInterest", err)
		return
	}
	tmp_LocationReportingReferenceID := NewINTEGER(ie.LocationReportingReferenceID, aper.Constraint{Lb: 1, Ub: 64}, false)
	if err = tmp_LocationReportingReferenceID.Encode(w); err != nil {
		err = utils.WrapError("Read LocationReportingReferenceID", err)
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
		err = utils.WrapError("Read AreaOfInterest", err)
		return
	}
	tmp_LocationReportingReferenceID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 64},
		ext: false,
	}
	if err = tmp_LocationReportingReferenceID.Decode(r); err != nil {
		err = utils.WrapError("Read LocationReportingReferenceID", err)
		return
	}
	ie.LocationReportingReferenceID = int64(tmp_LocationReportingReferenceID.Value)
	return
}
