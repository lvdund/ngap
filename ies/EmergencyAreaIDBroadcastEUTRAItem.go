package ies

import "github.com/lvdund/ngap/aper"

type EmergencyAreaIDBroadcastEUTRAItem struct {
	EmergencyAreaID          *EmergencyAreaID          `False,`
	CompletedCellsInEAIEUTRA *CompletedCellsInEAIEUTRA `False,`
	// IEExtensions EmergencyAreaIDBroadcastEUTRAItemExtIEs `False,OPTIONAL`
}

func (ie *EmergencyAreaIDBroadcastEUTRAItem) Encode(w *aper.AperWriter) (err error) {
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
	if ie.CompletedCellsInEAIEUTRA != nil {
		if err = ie.CompletedCellsInEAIEUTRA.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EmergencyAreaIDBroadcastEUTRAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.EmergencyAreaID = new(EmergencyAreaID)
	ie.CompletedCellsInEAIEUTRA = new(CompletedCellsInEAIEUTRA)
	if err = ie.EmergencyAreaID.Decode(r); err != nil {
		return
	}
	if err = ie.CompletedCellsInEAIEUTRA.Decode(r); err != nil {
		return
	}
	return
}
