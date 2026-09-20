package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type DRBStatusDL18 struct {
	DLCOUNTValue COUNTValueForPDCPSN18 `madatory`
	// IEExtension *DRBStatusDL18ExtIEs `optional`
}

func (ie *DRBStatusDL18) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.DLCOUNTValue.Encode(w); err != nil {
		err = fmt.Errorf("Encode DLCOUNTValue: %w", err)
		return
	}
	return
}
func (ie *DRBStatusDL18) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.DLCOUNTValue.Decode(r); err != nil {
		err = fmt.Errorf("Read DLCOUNTValue: %w", err)
		return
	}
	return
}
