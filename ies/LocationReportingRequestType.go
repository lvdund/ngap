package ies

import "github.com/lvdund/ngap/aper"

type LocationReportingRequestType struct {
	EventType                                 EventType
	ReportArea                                ReportArea
	AreaOfInterestList                        *[]AreaOfInterestItem
	LocationReportingReferenceIDToBeCancelled *int64
	// IEExtensions  *LocationReportingRequestTypeExtIEs
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
		return
	}
	if err = ie.ReportArea.Encode(w); err != nil {
		return
	}
	if ie.AreaOfInterestList != nil {
		if len(*ie.AreaOfInterestList) > 0 {
			tmp := Sequence[*AreaOfInterestItem]{
				Value: []*AreaOfInterestItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofAoI},
				ext:   false,
			}
			for _, i := range *ie.AreaOfInterestList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.LocationReportingReferenceIDToBeCancelled != nil {
		tmp_LocationReportingReferenceIDToBeCancelled := NewINTEGER(*ie.LocationReportingReferenceIDToBeCancelled, aper.Constraint{Lb: 1, Ub: 64}, true)
		if err = tmp_LocationReportingReferenceIDToBeCancelled.Encode(w); err != nil {
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
		return
	}
	if err = ie.ReportArea.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_AreaOfInterestList := Sequence[*AreaOfInterestItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAoI},
			ext: false,
		}
		if err = tmp_AreaOfInterestList.Decode(r); err != nil {
			return
		}
		ie.AreaOfInterestList = &[]AreaOfInterestItem{}
		for _, i := range tmp_AreaOfInterestList.Value {
			*ie.AreaOfInterestList = append(*ie.AreaOfInterestList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_LocationReportingReferenceIDToBeCancelled := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 64},
			ext: false,
		}
		if err = tmp_LocationReportingReferenceIDToBeCancelled.Decode(r); err != nil {
			return
		}
		*ie.LocationReportingReferenceIDToBeCancelled = int64(tmp_LocationReportingReferenceIDToBeCancelled.Value)
	}
	return
}
