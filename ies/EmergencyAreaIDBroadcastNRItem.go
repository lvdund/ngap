package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDBroadcastNRItem struct {
	EmergencyAreaID       *EmergencyAreaID       `False,`
	CompletedCellsInEAINR *CompletedCellsInEAINR `False,`
	// IEExtensions EmergencyAreaIDBroadcastNRItemExtIEs `False,OPTIONAL`
}

func (ie *EmergencyAreaIDBroadcastNRItem) Encode(w *aper.AperWriter) (err error) {
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
	if ie.CompletedCellsInEAINR != nil {
		if err = ie.CompletedCellsInEAINR.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EmergencyAreaIDBroadcastNRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.EmergencyAreaID = new(EmergencyAreaID)
	ie.CompletedCellsInEAINR = new(CompletedCellsInEAINR)
	if err = ie.EmergencyAreaID.Decode(r); err != nil {
		return
	}
	if err = ie.CompletedCellsInEAINR.Decode(r); err != nil {
		return
	}
	return
}
