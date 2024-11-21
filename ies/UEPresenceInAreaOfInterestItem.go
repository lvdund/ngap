package ies

import "github.com/lvdund/ngap/aper"

type UEPresenceInAreaOfInterestItem struct {
	LocationReportingReferenceID *LocationReportingReferenceID `False,`
	UEPresence                   *UEPresence                   `False,`
	// IEExtensions UEPresenceInAreaOfInterestItemExtIEs `False,OPTIONAL`
}

func (ie *UEPresenceInAreaOfInterestItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.LocationReportingReferenceID != nil {
		if err = ie.LocationReportingReferenceID.Encode(w); err != nil {
			return
		}
	}
	if ie.UEPresence != nil {
		if err = ie.UEPresence.Encode(w); err != nil {
			return
		}
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
	ie.LocationReportingReferenceID = new(LocationReportingReferenceID)
	ie.UEPresence = new(UEPresence)
	if err = ie.LocationReportingReferenceID.Decode(r); err != nil {
		return
	}
	if err = ie.UEPresence.Decode(r); err != nil {
		return
	}
	return
}
