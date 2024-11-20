package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDCancelledNRItem struct {
	EmergencyAreaID       *EmergencyAreaID       `False,`
	CancelledCellsInEAINR *CancelledCellsInEAINR `False,`
	// IEExtensions EmergencyAreaIDCancelledNRItemExtIEs `False,OPTIONAL`
}

func (ie *EmergencyAreaIDCancelledNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.EmergencyAreaID != nil {
		if err = ie.EmergencyAreaID.Encode(w); err != nil {
			return
		}
	}
	if ie.CancelledCellsInEAINR != nil {
		if err = ie.CancelledCellsInEAINR.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EmergencyAreaIDCancelledNRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.EmergencyAreaID = new(EmergencyAreaID)
	ie.CancelledCellsInEAINR = new(CancelledCellsInEAINR)
	if err = ie.EmergencyAreaID.Decode(r); err != nil {
		return
	}
	if err = ie.CancelledCellsInEAINR.Decode(r); err != nil {
		return
	}
	return
}
