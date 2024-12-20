package ies

import "github.com/lvdund/ngap/aper"

type UserLocationInformationEUTRA struct {
	EUTRACGI  *EUTRACGI  `True,`
	TAI       *TAI       `True,`
	TimeStamp *TimeStamp `False,OPTIONAL`
	// IEExtensions UserLocationInformationEUTRAExtIEs `False,OPTIONAL`
}

func (ie *UserLocationInformationEUTRA) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TimeStamp != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.EUTRACGI != nil {
		if err = ie.EUTRACGI.Encode(w); err != nil {
			return
		}
	}
	if ie.TAI != nil {
		if err = ie.TAI.Encode(w); err != nil {
			return
		}
	}
	if ie.TimeStamp != nil {
		if err = ie.TimeStamp.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *UserLocationInformationEUTRA) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.EUTRACGI = new(EUTRACGI)
	ie.TAI = new(TAI)
	ie.TimeStamp = new(TimeStamp)
	if err = ie.EUTRACGI.Decode(r); err != nil {
		return
	}
	if err = ie.TAI.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.TimeStamp.Decode(r); err != nil {
			return
		}
	}
	return
}
