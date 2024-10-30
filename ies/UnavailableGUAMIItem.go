package ies

import "github.com/lvdund/ngap/aper"

type UnavailableGUAMIItem struct {
	GUAMI                        *GUAMI                        `True,`
	TimerApproachForGUAMIRemoval *TimerApproachForGUAMIRemoval `False,OPTIONAL`
	BackupAMFName                *AMFName                      `False,OPTIONAL`
	// IEExtensions UnavailableGUAMIItemExtIEs `False,OPTIONAL`
}

func (ie *UnavailableGUAMIItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TimerApproachForGUAMIRemoval != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.BackupAMFName != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.GUAMI != nil {
		if err = ie.GUAMI.Encode(w); err != nil {
			return
		}
	}
	if ie.TimerApproachForGUAMIRemoval != nil {
		if err = ie.TimerApproachForGUAMIRemoval.Encode(w); err != nil {
			return
		}
	}
	if ie.BackupAMFName != nil {
		if err = ie.BackupAMFName.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *UnavailableGUAMIItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	ie.GUAMI = new(GUAMI)
	ie.TimerApproachForGUAMIRemoval = new(TimerApproachForGUAMIRemoval)
	ie.BackupAMFName = new(AMFName)
	if err = ie.GUAMI.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.TimerApproachForGUAMIRemoval.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.BackupAMFName.Decode(r); err != nil {
			return
		}
	}
	return
}
