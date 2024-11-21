package ies

import "github.com/lvdund/ngap/aper"

type AreaOfInterest struct {
	AreaOfInterestTAIList     *AreaOfInterestTAIList     `False,OPTIONAL`
	AreaOfInterestCellList    *AreaOfInterestCellList    `False,OPTIONAL`
	AreaOfInterestRANNodeList *AreaOfInterestRANNodeList `False,OPTIONAL`
	// IEExtensions AreaOfInterestExtIEs `False,OPTIONAL`
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
		if err = ie.AreaOfInterestTAIList.Encode(w); err != nil {
			return
		}
	}
	if ie.AreaOfInterestCellList != nil {
		if err = ie.AreaOfInterestCellList.Encode(w); err != nil {
			return
		}
	}
	if ie.AreaOfInterestRANNodeList != nil {
		if err = ie.AreaOfInterestRANNodeList.Encode(w); err != nil {
			return
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
	ie.AreaOfInterestTAIList = new(AreaOfInterestTAIList)
	ie.AreaOfInterestCellList = new(AreaOfInterestCellList)
	ie.AreaOfInterestRANNodeList = new(AreaOfInterestRANNodeList)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.AreaOfInterestTAIList.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.AreaOfInterestCellList.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.AreaOfInterestRANNodeList.Decode(r); err != nil {
			return
		}
	}
	return
}
