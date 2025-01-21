package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type LastVisitedNGRANCellInformation struct {
	GlobalCellID                          NGRANCGI `madatory`
	CellType                              CellType `madatory`
	TimeUEStayedInCell                    int64    `lb:0,ub:4095,madatory`
	TimeUEStayedInCellEnhancedGranularity *int64   `lb:0,ub:40950,optional`
	HOCauseValue                          *Cause   `optional`
	// IEExtensions *LastVisitedNGRANCellInformationExtIEs `optional`
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
		err = utils.WrapError("Encode GlobalCellID", err)
		return
	}
	if err = ie.CellType.Encode(w); err != nil {
		err = utils.WrapError("Encode CellType", err)
		return
	}
	tmp_TimeUEStayedInCell := NewINTEGER(ie.TimeUEStayedInCell, aper.Constraint{Lb: 0, Ub: 4095}, false)
	if err = tmp_TimeUEStayedInCell.Encode(w); err != nil {
		err = utils.WrapError("Encode TimeUEStayedInCell", err)
		return
	}
	if ie.TimeUEStayedInCellEnhancedGranularity != nil {
		tmp_TimeUEStayedInCellEnhancedGranularity := NewINTEGER(*ie.TimeUEStayedInCellEnhancedGranularity, aper.Constraint{Lb: 0, Ub: 40950}, false)
		if err = tmp_TimeUEStayedInCellEnhancedGranularity.Encode(w); err != nil {
			err = utils.WrapError("Encode TimeUEStayedInCellEnhancedGranularity", err)
			return
		}
	}
	if ie.HOCauseValue != nil {
		if err = ie.HOCauseValue.Encode(w); err != nil {
			err = utils.WrapError("Encode HOCauseValue", err)
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
		err = utils.WrapError("Read GlobalCellID", err)
		return
	}
	if err = ie.CellType.Decode(r); err != nil {
		err = utils.WrapError("Read CellType", err)
		return
	}
	tmp_TimeUEStayedInCell := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4095},
		ext: false,
	}
	if err = tmp_TimeUEStayedInCell.Decode(r); err != nil {
		err = utils.WrapError("Read TimeUEStayedInCell", err)
		return
	}
	ie.TimeUEStayedInCell = int64(tmp_TimeUEStayedInCell.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_TimeUEStayedInCellEnhancedGranularity := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 40950},
			ext: false,
		}
		if err = tmp_TimeUEStayedInCellEnhancedGranularity.Decode(r); err != nil {
			err = utils.WrapError("Read TimeUEStayedInCellEnhancedGranularity", err)
			return
		}
		ie.TimeUEStayedInCellEnhancedGranularity = (*int64)(&tmp_TimeUEStayedInCellEnhancedGranularity.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(Cause)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read HOCauseValue", err)
			return
		}
		ie.HOCauseValue = tmp
	}
	return
}
