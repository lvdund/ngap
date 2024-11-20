package ies

import "github.com/lvdund/ngap/aper"

type DRBStatusUL18 struct {
	ULCOUNTValue              *COUNTValueForPDCPSN18 `True,`
	ReceiveStatusOfULPDCPSDUs *aper.BitString        `True,OPTIONAL`
	// IEExtension DRBStatusUL18ExtIEs `False,OPTIONAL`
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
	if ie.ULCOUNTValue != nil {
		if err = ie.ULCOUNTValue.Encode(w); err != nil {
			return
		}
	}
	if ie.ReceiveStatusOfULPDCPSDUs != nil {
		if err = w.WriteBitString(ie.ReceiveStatusOfULPDCPSDUs.Bytes, uint(ie.ReceiveStatusOfULPDCPSDUs.NumBits), &aper.Constraint{Lb: 1, Ub: 131072}, true); err != nil {
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
	ie.ULCOUNTValue = new(COUNTValueForPDCPSN18)
	var b []byte
	var n uint
	if err = ie.ULCOUNTValue.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if b, n, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 131072}, true); err != nil {
			return
		} else {
			ie.ReceiveStatusOfULPDCPSDUs = &aper.BitString{Bytes: b, NumBits: uint64(n)}
		}
	}
	return
}
