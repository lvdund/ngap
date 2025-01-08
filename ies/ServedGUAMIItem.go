package ies

import "github.com/lvdund/ngap/aper"

type ServedGUAMIItem struct {
	GUAMI         GUAMI
	BackupAMFName []byte
	// IEExtensions *ServedGUAMIItemExtIEs `optional`
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
	if err = ie.GUAMI.Encode(w); err != nil {
		return
	}
	if ie.BackupAMFName != nil {
		tmp_BackupAMFName := NewOCTETSTRING(ie.BackupAMFName, aper.Constraint{Lb: 1, Ub: 150}, false)
		if err = tmp_BackupAMFName.Encode(w); err != nil {
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
	if err = ie.GUAMI.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_BackupAMFName := OCTETSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 150},
			ext: false,
		}
		if err = tmp_BackupAMFName.Decode(r); err != nil {
			return
		}
		ie.BackupAMFName = tmp_BackupAMFName.Value
	}
	return
}
