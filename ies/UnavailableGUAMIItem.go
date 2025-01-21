package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UnavailableGUAMIItem struct {
	GUAMI                        GUAMI                         `madatory`
	TimerApproachForGUAMIRemoval *TimerApproachForGUAMIRemoval `optional`
	BackupAMFName                []byte                        `lb:1,ub:150,optional,valExt`
	// IEExtensions *UnavailableGUAMIItemExtIEs `optional`
}

func (ie *UnavailableGUAMIItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
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
	if err = ie.GUAMI.Encode(w); err != nil {
		err = utils.WrapError("Encode GUAMI", err)
		return
	}
	if ie.TimerApproachForGUAMIRemoval != nil {
		if err = ie.TimerApproachForGUAMIRemoval.Encode(w); err != nil {
			err = utils.WrapError("Encode TimerApproachForGUAMIRemoval", err)
			return
		}
	}
	if ie.BackupAMFName != nil {
		tmp_BackupAMFName := NewOCTETSTRING(ie.BackupAMFName, aper.Constraint{Lb: 1, Ub: 150}, true)
		if err = tmp_BackupAMFName.Encode(w); err != nil {
			err = utils.WrapError("Encode BackupAMFName", err)
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
	if err = ie.GUAMI.Decode(r); err != nil {
		err = utils.WrapError("Read GUAMI", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(TimerApproachForGUAMIRemoval)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read TimerApproachForGUAMIRemoval", err)
			return
		}
		ie.TimerApproachForGUAMIRemoval = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_BackupAMFName := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: true,
		}
		if err = tmp_BackupAMFName.Decode(r); err != nil {
			err = utils.WrapError("Read BackupAMFName", err)
			return
		}
		ie.BackupAMFName = tmp_BackupAMFName.Value
	}
	return
}
