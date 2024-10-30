package ies

import "github.com/lvdund/ngap/aper"

type DRBStatusDL18 struct {
	DLCOUNTValue *COUNTValueForPDCPSN18 `True,`
	// IEExtension DRBStatusDL18ExtIEs `False,OPTIONAL`
}

func (ie *DRBStatusDL18) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.DLCOUNTValue != nil {
		if err = ie.DLCOUNTValue.Encode(w); err != nil {
			return
		}
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
	ie.DLCOUNTValue = new(COUNTValueForPDCPSN18)
	if err = ie.DLCOUNTValue.Decode(r); err != nil {
		return
	}
	return
}
