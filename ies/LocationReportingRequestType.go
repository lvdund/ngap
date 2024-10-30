package ies

import "github.com/lvdund/ngap/aper"

type LocationReportingRequestType struct {
	EventType                                 *EventType                    `False,`
	ReportArea                                *ReportArea                   `False,`
	AreaOfInterestList                        *AreaOfInterestList           `False,OPTIONAL`
	LocationReportingReferenceIDToBeCancelled *LocationReportingReferenceID `False,OPTIONAL`
	// IEExtensions LocationReportingRequestTypeExtIEs `False,OPTIONAL`
}

func (ie *LocationReportingRequestType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
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
	if ie.EventType != nil {
		if err = ie.EventType.Encode(w); err != nil {
			return
		}
	}
	if ie.ReportArea != nil {
		if err = ie.ReportArea.Encode(w); err != nil {
			return
		}
	}
	if ie.AreaOfInterestList != nil {
		if err = ie.AreaOfInterestList.Encode(w); err != nil {
			return
		}
	}
	if ie.LocationReportingReferenceIDToBeCancelled != nil {
		if err = ie.LocationReportingReferenceIDToBeCancelled.Encode(w); err != nil {
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
	ie.EventType = new(EventType)
	ie.ReportArea = new(ReportArea)
	ie.AreaOfInterestList = new(AreaOfInterestList)
	ie.LocationReportingReferenceIDToBeCancelled = new(LocationReportingReferenceID)
	if err = ie.EventType.Decode(r); err != nil {
		return
	}
	if err = ie.ReportArea.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.AreaOfInterestList.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.LocationReportingReferenceIDToBeCancelled.Decode(r); err != nil {
			return
		}
	}
	return
}
