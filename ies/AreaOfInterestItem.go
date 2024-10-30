package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterestItem struct {
	AreaOfInterest               *AreaOfInterest               `True,`
	LocationReportingReferenceID *LocationReportingReferenceID `False,`
	// IEExtensions AreaOfInterestItemExtIEs `False,OPTIONAL`
}

func (ie *AreaOfInterestItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.AreaOfInterest != nil {
		if err = ie.AreaOfInterest.Encode(w); err != nil {
			return
		}
	}
	if ie.LocationReportingReferenceID != nil {
		if err = ie.LocationReportingReferenceID.Encode(w); err != nil {
			return
		}
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
	ie.AreaOfInterest = new(AreaOfInterest)
	ie.LocationReportingReferenceID = new(LocationReportingReferenceID)
	if err = ie.AreaOfInterest.Decode(r); err != nil {
		return
	}
	if err = ie.LocationReportingReferenceID.Decode(r); err != nil {
		return
	}
	return
}
