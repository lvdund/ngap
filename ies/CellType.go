package ies

import "github.com/lvdund/ngap/aper"

type CellType struct {
	CellSize *CellSize `False,`
	// IEExtensions CellTypeExtIEs `False,OPTIONAL`
}

func (ie *CellType) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.CellSize != nil {
		if err = ie.CellSize.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CellType) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.CellSize = new(CellSize)
	if err = ie.CellSize.Decode(r); err != nil {
		return
	}
	return
}
