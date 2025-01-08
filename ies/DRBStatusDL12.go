package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBStatusDL12 struct {
	DLCOUNTValue COUNTValueForPDCPSN12
	// IEExtension *DRBStatusDL12ExtIEs `optional`
}

func (ie *DRBStatusDL12) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.DLCOUNTValue.Encode(w); err != nil {
		err = utils.WrapError("Read DLCOUNTValue", err)
		return
	}
	return
}
func (ie *DRBStatusDL12) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.DLCOUNTValue.Decode(r); err != nil {
		err = utils.WrapError("Read DLCOUNTValue", err)
		return
	}
	return
}
