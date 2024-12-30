package ies

import "github.com/lvdund/ngap/aper"

type UEPresenceInAreaOfInterestItem struct {
	LocationReportingReferenceID int64
	UEPresence                   UEPresence
	// IEExtensions  *UEPresenceInAreaOfInterestItemExtIEs
}

func (ie *UEPresenceInAreaOfInterestItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_LocationReportingReferenceID := NewINTEGER(ie.LocationReportingReferenceID, aper.Constraint{Lb: 1, Ub: 64}, false)
	if err = tmp_LocationReportingReferenceID.Encode(w); err != nil {
		return
	}
	if err = ie.UEPresence.Encode(w); err != nil {
		return
	}
	return
}
func (ie *UEPresenceInAreaOfInterestItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_LocationReportingReferenceID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 64},
		ext: false,
	}
	if err = tmp_LocationReportingReferenceID.Decode(r); err != nil {
		return
	}
	ie.LocationReportingReferenceID = int64(tmp_LocationReportingReferenceID.Value)
	if err = ie.UEPresence.Decode(r); err != nil {
		return
	}
	return
}
