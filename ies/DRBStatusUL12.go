package ies

import "github.com/lvdund/ngap/aper"

type DRBStatusUL12 struct {
	ULCOUNTValue              COUNTValueForPDCPSN12
	ReceiveStatusOfULPDCPSDUs *[]byte
	// IEExtension  *DRBStatusUL12ExtIEs
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
		return
	}
	if ie.ReceiveStatusOfULPDCPSDUs != nil {
		tmp_ReceiveStatusOfULPDCPSDUs := NewBITSTRING(*ie.ReceiveStatusOfULPDCPSDUs, aper.Constraint{Lb: 0, Ub: 0}, true)
		if err = tmp_ReceiveStatusOfULPDCPSDUs.Encode(w); err != nil {
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
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_ReceiveStatusOfULPDCPSDUs := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_ReceiveStatusOfULPDCPSDUs.Decode(r); err != nil {
			return
		}
		*ie.ReceiveStatusOfULPDCPSDUs = tmp_ReceiveStatusOfULPDCPSDUs.Value.Bytes
	}
	return
}
