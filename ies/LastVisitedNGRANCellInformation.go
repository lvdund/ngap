package ies

import "github.com/lvdund/ngap/aper"

type LastVisitedNGRANCellInformation struct {
	GlobalCellID                          *NGRANCGI                              `False,`
	CellType                              *CellType                              `True,`
	TimeUEStayedInCell                    *TimeUEStayedInCell                    `False,`
	TimeUEStayedInCellEnhancedGranularity *TimeUEStayedInCellEnhancedGranularity `False,OPTIONAL`
	HOCauseValue                          *Cause                                 `False,OPTIONAL`
	// IEExtensions LastVisitedNGRANCellInformationExtIEs `False,OPTIONAL`
}

func (ie *LastVisitedNGRANCellInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TimeUEStayedInCellEnhancedGranularity != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.HOCauseValue != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.GlobalCellID != nil {
		if err = ie.GlobalCellID.Encode(w); err != nil {
			return
		}
	}
	if ie.CellType != nil {
		if err = ie.CellType.Encode(w); err != nil {
			return
		}
	}
	if ie.TimeUEStayedInCell != nil {
		if err = ie.TimeUEStayedInCell.Encode(w); err != nil {
			return
		}
	}
	if ie.TimeUEStayedInCellEnhancedGranularity != nil {
		if err = ie.TimeUEStayedInCellEnhancedGranularity.Encode(w); err != nil {
			return
		}
	}
	if ie.HOCauseValue != nil {
		if err = ie.HOCauseValue.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *LastVisitedNGRANCellInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	ie.GlobalCellID = new(NGRANCGI)
	ie.CellType = new(CellType)
	ie.TimeUEStayedInCell = new(TimeUEStayedInCell)
	ie.TimeUEStayedInCellEnhancedGranularity = new(TimeUEStayedInCellEnhancedGranularity)
	ie.HOCauseValue = new(Cause)
	if err = ie.GlobalCellID.Decode(r); err != nil {
		return
	}
	if err = ie.CellType.Decode(r); err != nil {
		return
	}
	if err = ie.TimeUEStayedInCell.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.TimeUEStayedInCellEnhancedGranularity.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.HOCauseValue.Decode(r); err != nil {
			return
		}
	}
	return
}
