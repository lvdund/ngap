package ies

import "github.com/lvdund/ngap/aper"

type ServedGUAMIItem struct {
	GUAMI         *GUAMI   `True,`
	BackupAMFName *AMFName `False,OPTIONAL`
	// IEExtensions ServedGUAMIItemExtIEs `False,OPTIONAL`
}

func (ie *ServedGUAMIItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.BackupAMFName != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.GUAMI != nil {
		if err = ie.GUAMI.Encode(w); err != nil {
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
func (ie *ServedGUAMIItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.GUAMI = new(GUAMI)
	ie.BackupAMFName = new(AMFName)
	if err = ie.GUAMI.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.BackupAMFName.Decode(r); err != nil {
			return
		}
	}
	return
}
