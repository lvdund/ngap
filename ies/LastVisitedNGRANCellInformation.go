package ies

import "github.com/lvdund/ngap/aper"

type LastVisitedNGRANCellInformation struct {
	GlobalCellID                          NGRANCGI
	CellType                              CellType
	TimeUEStayedInCell                    int64
	TimeUEStayedInCellEnhancedGranularity *int64
	HOCauseValue                          *Cause
	// IEExtensions  *LastVisitedNGRANCellInformationExtIEs
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
	if err = ie.GlobalCellID.Encode(w); err != nil {
		return
	}
	if err = ie.CellType.Encode(w); err != nil {
		return
	}
	tmp_TimeUEStayedInCell := NewINTEGER(ie.TimeUEStayedInCell, aper.Constraint{Lb: 0, Ub: 4095}, false)
	if err = tmp_TimeUEStayedInCell.Encode(w); err != nil {
		return
	}
	if ie.TimeUEStayedInCellEnhancedGranularity != nil {
		tmp_TimeUEStayedInCellEnhancedGranularity := NewINTEGER(*ie.TimeUEStayedInCellEnhancedGranularity, aper.Constraint{Lb: 0, Ub: 40950}, false)
		if err = tmp_TimeUEStayedInCellEnhancedGranularity.Encode(w); err != nil {
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
	if err = ie.GlobalCellID.Decode(r); err != nil {
		return
	}
	if err = ie.CellType.Decode(r); err != nil {
		return
	}
	tmp_TimeUEStayedInCell := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: false,
	}
	if err = tmp_TimeUEStayedInCell.Decode(r); err != nil {
		return
	}
	ie.TimeUEStayedInCell = int64(tmp_TimeUEStayedInCell.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_TimeUEStayedInCellEnhancedGranularity := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 40950},
			ext: false,
		}
		if err = tmp_TimeUEStayedInCellEnhancedGranularity.Decode(r); err != nil {
			return
		}
		*ie.TimeUEStayedInCellEnhancedGranularity = int64(tmp_TimeUEStayedInCellEnhancedGranularity.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.HOCauseValue.Decode(r); err != nil {
			return
		}
	}
	return
}
