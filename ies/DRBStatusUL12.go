package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBStatusUL12 struct {
	ULCOUNTValue              COUNTValueForPDCPSN12
	ReceiveStatusOfULPDCPSDUs []byte
	// IEExtension *DRBStatusUL12ExtIEs `optional`
}

func (ie *DRBStatusUL12) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ReceiveStatusOfULPDCPSDUs != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.ULCOUNTValue.Encode(w); err != nil {
		err = utils.WrapError("Read ULCOUNTValue", err)
		return
	}
	if ie.ReceiveStatusOfULPDCPSDUs != nil {
		tmp_ReceiveStatusOfULPDCPSDUs := NewBITSTRING(ie.ReceiveStatusOfULPDCPSDUs, aper.Constraint{Lb: 1, Ub: 2048}, false)
		if err = tmp_ReceiveStatusOfULPDCPSDUs.Encode(w); err != nil {
			err = utils.WrapError("Read ReceiveStatusOfULPDCPSDUs", err)
			return
		}
	}
	return
}
func (ie *DRBStatusUL12) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.ULCOUNTValue.Decode(r); err != nil {
		err = utils.WrapError("Read ULCOUNTValue", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_ReceiveStatusOfULPDCPSDUs := BITSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 2048},
			ext: false,
		}
		if err = tmp_ReceiveStatusOfULPDCPSDUs.Decode(r); err != nil {
			err = utils.WrapError("Read ReceiveStatusOfULPDCPSDUs", err)
			return
		}
		ie.ReceiveStatusOfULPDCPSDUs = tmp_ReceiveStatusOfULPDCPSDUs.Value.Bytes
	}
	return
}
