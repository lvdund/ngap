package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          *EmergencyAreaID          `False,`
	CancelledCellsInEAIEUTRA *CancelledCellsInEAIEUTRA `False,`
	// IEExtensions EmergencyAreaIDCancelledEUTRAItemExtIEs `False,OPTIONAL`
}

func (ie *EmergencyAreaIDCancelledEUTRAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.EmergencyAreaID != nil {
		if err = ie.EmergencyAreaID.Encode(w); err != nil {
			return
		}
	}
	if ie.CancelledCellsInEAIEUTRA != nil {
		if err = ie.CancelledCellsInEAIEUTRA.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EmergencyAreaIDCancelledEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.EmergencyAreaID = new(EmergencyAreaID)
	ie.CancelledCellsInEAIEUTRA = new(CancelledCellsInEAIEUTRA)
	if err = ie.EmergencyAreaID.Decode(r); err != nil {
		return
	}
	if err = ie.CancelledCellsInEAIEUTRA.Decode(r); err != nil {
		return
	}
	return
}
