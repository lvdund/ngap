package ies

import "github.com/lvdund/ngap/aper"

type COUNTValueForPDCPSN12 struct {
	PDCPSN12    *aper.Integer `False,`
	HFNPDCPSN12 *aper.Integer `False,`
	// IEExtensions COUNTValueForPDCPSN12ExtIEs `False,OPTIONAL`
}

func (ie *COUNTValueForPDCPSN12) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PDCPSN12 != nil {
		if err = w.WriteInteger(int64(*ie.PDCPSN12), &aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return
		}
	}
	if ie.HFNPDCPSN12 != nil {
		if err = w.WriteInteger(int64(*ie.HFNPDCPSN12), &aper.Constraint{Lb: 0, Ub: 1048575}, false); err != nil {
			return
		}
	}
	return
}
func (ie *COUNTValueForPDCPSN12) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return
	} else {
		ie.PDCPSN12 = (*aper.Integer)(&v)
	}
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1048575}, false); err != nil {
		return
	} else {
		ie.HFNPDCPSN12 = (*aper.Integer)(&v)
	}
	return
}
