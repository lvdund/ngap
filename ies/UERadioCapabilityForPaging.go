package ies

import "github.com/lvdund/ngap/aper"

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *[]byte
	UERadioCapabilityForPagingOfEUTRA *[]byte
	// IEExtensions  *UERadioCapabilityForPagingExtIEs
}

func (ie *UERadioCapabilityForPaging) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.UERadioCapabilityForPagingOfNR != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.UERadioCapabilityForPagingOfEUTRA != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.UERadioCapabilityForPagingOfNR != nil {
		tmp_UERadioCapabilityForPagingOfNR := NewOCTETSTRING(*ie.UERadioCapabilityForPagingOfNR, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_UERadioCapabilityForPagingOfNR.Encode(w); err != nil {
			return
		}
	}
	if ie.UERadioCapabilityForPagingOfEUTRA != nil {
		tmp_UERadioCapabilityForPagingOfEUTRA := NewOCTETSTRING(*ie.UERadioCapabilityForPagingOfEUTRA, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_UERadioCapabilityForPagingOfEUTRA.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *UERadioCapabilityForPaging) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_UERadioCapabilityForPagingOfNR := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_UERadioCapabilityForPagingOfNR.Decode(r); err != nil {
			return
		}
		*ie.UERadioCapabilityForPagingOfNR = tmp_UERadioCapabilityForPagingOfNR.Value
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_UERadioCapabilityForPagingOfEUTRA := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_UERadioCapabilityForPagingOfEUTRA.Decode(r); err != nil {
			return
		}
		*ie.UERadioCapabilityForPagingOfEUTRA = tmp_UERadioCapabilityForPagingOfEUTRA.Value
	}
	return
}
