package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterest struct {
	AreaOfInterestTAIList     *[]AreaOfInterestTAIItem
	AreaOfInterestCellList    *[]AreaOfInterestCellItem
	AreaOfInterestRANNodeList *[]AreaOfInterestRANNodeItem
	// IEExtensions  *AreaOfInterestExtIEs
}

func (ie *AreaOfInterest) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AreaOfInterestTAIList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.AreaOfInterestCellList != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.AreaOfInterestRANNodeList != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.AreaOfInterestTAIList != nil {
		if len(*ie.AreaOfInterestTAIList) > 0 {
			tmp := Sequence[*AreaOfInterestTAIItem]{
				Value: []*AreaOfInterestTAIItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofTAIinAoI},
				ext:   false,
			}
			for _, i := range *ie.AreaOfInterestTAIList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.AreaOfInterestCellList != nil {
		if len(*ie.AreaOfInterestCellList) > 0 {
			tmp := Sequence[*AreaOfInterestCellItem]{
				Value: []*AreaOfInterestCellItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofCellinAoI},
				ext:   false,
			}
			for _, i := range *ie.AreaOfInterestCellList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.AreaOfInterestRANNodeList != nil {
		if len(*ie.AreaOfInterestRANNodeList) > 0 {
			tmp := Sequence[*AreaOfInterestRANNodeItem]{
				Value: []*AreaOfInterestRANNodeItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofRANNodeinAoI},
				ext:   false,
			}
			for _, i := range *ie.AreaOfInterestRANNodeList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	return
}
func (ie *AreaOfInterest) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_AreaOfInterestTAIList := Sequence[*AreaOfInterestTAIItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIinAoI},
			ext: false,
		}
		if err = tmp_AreaOfInterestTAIList.Decode(r); err != nil {
			return
		}
		ie.AreaOfInterestTAIList = &[]AreaOfInterestTAIItem{}
		for _, i := range tmp_AreaOfInterestTAIList.Value {
			*ie.AreaOfInterestTAIList = append(*ie.AreaOfInterestTAIList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_AreaOfInterestCellList := Sequence[*AreaOfInterestCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellinAoI},
			ext: false,
		}
		if err = tmp_AreaOfInterestCellList.Decode(r); err != nil {
			return
		}
		ie.AreaOfInterestCellList = &[]AreaOfInterestCellItem{}
		for _, i := range tmp_AreaOfInterestCellList.Value {
			*ie.AreaOfInterestCellList = append(*ie.AreaOfInterestCellList, *i)
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_AreaOfInterestRANNodeList := Sequence[*AreaOfInterestRANNodeItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofRANNodeinAoI},
			ext: false,
		}
		if err = tmp_AreaOfInterestRANNodeList.Decode(r); err != nil {
			return
		}
		ie.AreaOfInterestRANNodeList = &[]AreaOfInterestRANNodeItem{}
		for _, i := range tmp_AreaOfInterestRANNodeList.Value {
			*ie.AreaOfInterestRANNodeList = append(*ie.AreaOfInterestRANNodeList, *i)
		}
	}
	return
}
