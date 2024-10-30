package ies

import "github.com/lvdund/ngap/aper"

type COUNTValueForPDCPSN18 struct {
	PDCPSN18    *aper.Integer `False,`
	HFNPDCPSN18 *aper.Integer `False,`
	// IEExtensions COUNTValueForPDCPSN18ExtIEs `False,OPTIONAL`
}

func (ie *COUNTValueForPDCPSN18) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PDCPSN18 != nil {
		if err = w.WriteInteger(int64(*ie.PDCPSN18), &aper.Constraint{Lb: 0, Ub: 262143}, false); err != nil {
			return
		}
	}
	if ie.HFNPDCPSN18 != nil {
		if err = w.WriteInteger(int64(*ie.HFNPDCPSN18), &aper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
			return
		}
	}
	return
}
func (ie *COUNTValueForPDCPSN18) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 262143}, false); err != nil {
		return
	} else {
		ie.PDCPSN18 = (*aper.Integer)(&v)
	}
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
		return
	} else {
		ie.HFNPDCPSN18 = (*aper.Integer)(&v)
	}
	return
}
