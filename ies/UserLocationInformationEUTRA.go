package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UserLocationInformationEUTRA struct {
	EUTRACGI  EUTRACGI `madatory`
	TAI       TAI      `madatory`
	TimeStamp []byte   `lb:4,ub:4,optional`
	// IEExtensions *UserLocationInformationEUTRAExtIEs `optional`
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
	if err = ie.EUTRACGI.Encode(w); err != nil {
		err = utils.WrapError("Encode EUTRACGI", err)
		return
	}
	if err = ie.TAI.Encode(w); err != nil {
		err = utils.WrapError("Encode TAI", err)
		return
	}
	if ie.TimeStamp != nil {
		tmp_TimeStamp := NewOCTETSTRING(ie.TimeStamp, aper.Constraint{Lb: 4, Ub: 4}, false)
		if err = tmp_TimeStamp.Encode(w); err != nil {
			err = utils.WrapError("Encode TimeStamp", err)
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
	if err = ie.EUTRACGI.Decode(r); err != nil {
		err = utils.WrapError("Read EUTRACGI", err)
		return
	}
	if err = ie.TAI.Decode(r); err != nil {
		err = utils.WrapError("Read TAI", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_TimeStamp := OCTETSTRING{
			c:   aper.Constraint{Lb: 4, Ub: 4},
			ext: false,
		}
		if err = tmp_TimeStamp.Decode(r); err != nil {
			err = utils.WrapError("Read TimeStamp", err)
			return
		}
		ie.TimeStamp = tmp_TimeStamp.Value
	}
	return
}
