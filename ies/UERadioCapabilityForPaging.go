package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    []byte `lb:0,ub:0,optional`
	UERadioCapabilityForPagingOfEUTRA []byte `lb:0,ub:0,optional`
	// IEExtensions *UERadioCapabilityForPagingExtIEs `optional`
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
		tmp_UERadioCapabilityForPagingOfNR := NewOCTETSTRING(ie.UERadioCapabilityForPagingOfNR, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_UERadioCapabilityForPagingOfNR.Encode(w); err != nil {
			err = utils.WrapError("Encode UERadioCapabilityForPagingOfNR", err)
			return
		}
	}
	if ie.UERadioCapabilityForPagingOfEUTRA != nil {
		tmp_UERadioCapabilityForPagingOfEUTRA := NewOCTETSTRING(ie.UERadioCapabilityForPagingOfEUTRA, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_UERadioCapabilityForPagingOfEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode UERadioCapabilityForPagingOfEUTRA", err)
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
			err = utils.WrapError("Read UERadioCapabilityForPagingOfNR", err)
			return
		}
		ie.UERadioCapabilityForPagingOfNR = tmp_UERadioCapabilityForPagingOfNR.Value
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_UERadioCapabilityForPagingOfEUTRA := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_UERadioCapabilityForPagingOfEUTRA.Decode(r); err != nil {
			err = utils.WrapError("Read UERadioCapabilityForPagingOfEUTRA", err)
			return
		}
		ie.UERadioCapabilityForPagingOfEUTRA = tmp_UERadioCapabilityForPagingOfEUTRA.Value
	}
	return
}
