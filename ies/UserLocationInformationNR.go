package ies

import "github.com/lvdund/ngap/aper"

type UserLocationInformationNR struct {
	NRCGI     NRCGI
	TAI       TAI
	TimeStamp []byte
	// IEExtensions *UserLocationInformationNRExtIEs `optional`
}

func (ie *UserLocationInformationNR) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TimeStamp != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.NRCGI.Encode(w); err != nil {
		return
	}
	if err = ie.TAI.Encode(w); err != nil {
		return
	}
	if ie.TimeStamp != nil {
		tmp_TimeStamp := NewOCTETSTRING(ie.TimeStamp, aper.Constraint{Lb: 4, Ub: 4}, false)
		if err = tmp_TimeStamp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *UserLocationInformationNR) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.NRCGI.Decode(r); err != nil {
		return
	}
	if err = ie.TAI.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_TimeStamp := OCTETSTRING{
			c:   aper.Constraint{Lb: 4, Ub: 4},
			ext: false,
		}
		if err = tmp_TimeStamp.Decode(r); err != nil {
			return
		}
		ie.TimeStamp = tmp_TimeStamp.Value
	}
	return
}
