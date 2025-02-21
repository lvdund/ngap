package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DRBStatusUL18 struct {
	ULCOUNTValue              COUNTValueForPDCPSN18 `madatory`
	ReceiveStatusOfULPDCPSDUs *aper.BitString       `lb:1,ub:131072,optional`
	// IEExtension *DRBStatusUL18ExtIEs `optional`
}

func (ie *DRBStatusUL18) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ReceiveStatusOfULPDCPSDUs != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.ULCOUNTValue.Encode(w); err != nil {
		err = utils.WrapError("Encode ULCOUNTValue", err)
		return
	}
	if ie.ReceiveStatusOfULPDCPSDUs != nil {
		tmp_ReceiveStatusOfULPDCPSDUs := NewBITSTRING(*ie.ReceiveStatusOfULPDCPSDUs, aper.Constraint{Lb: 1, Ub: 131072}, false)
		if err = tmp_ReceiveStatusOfULPDCPSDUs.Encode(w); err != nil {
			err = utils.WrapError("Encode ReceiveStatusOfULPDCPSDUs", err)
			return
		}
	}
	return
}
func (ie *DRBStatusUL18) Decode(r *aper.AperReader) (err error) {
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
			c:   aper.Constraint{Lb: 1, Ub: 131072},
			ext: false,
		}
		if err = tmp_ReceiveStatusOfULPDCPSDUs.Decode(r); err != nil {
			err = utils.WrapError("Read ReceiveStatusOfULPDCPSDUs", err)
			return
		}
		ie.ReceiveStatusOfULPDCPSDUs = &aper.BitString{
			Bytes:   tmp_ReceiveStatusOfULPDCPSDUs.Value.Bytes,
			NumBits: tmp_ReceiveStatusOfULPDCPSDUs.Value.NumBits,
		}
	}
	return
}
