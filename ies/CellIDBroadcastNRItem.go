package ies

import "github.com/lvdund/ngap/aper"

type CellIDBroadcastNRItem struct {
	NRCGI NRCGI
	// IEExtensions  *CellIDBroadcastNRItemExtIEs
}

func (ie *CellIDBroadcastNRItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.NRCGI.Encode(w); err != nil {
		return
	}
	return
}
func (ie *CellIDBroadcastNRItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.NRCGI.Decode(r); err != nil {
		return
	}
	return
}
