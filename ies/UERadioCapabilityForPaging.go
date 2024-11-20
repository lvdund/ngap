package ies

import "github.com/lvdund/ngap/aper"

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *UERadioCapabilityForPagingOfNR    `False,OPTIONAL`
	UERadioCapabilityForPagingOfEUTRA *UERadioCapabilityForPagingOfEUTRA `False,OPTIONAL`
	// IEExtensions UERadioCapabilityForPagingExtIEs `False,OPTIONAL`
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
		if err = ie.UERadioCapabilityForPagingOfNR.Encode(w); err != nil {
			return
		}
	}
	if ie.UERadioCapabilityForPagingOfEUTRA != nil {
		if err = ie.UERadioCapabilityForPagingOfEUTRA.Encode(w); err != nil {
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
	ie.UERadioCapabilityForPagingOfNR = new(UERadioCapabilityForPagingOfNR)
	ie.UERadioCapabilityForPagingOfEUTRA = new(UERadioCapabilityForPagingOfEUTRA)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.UERadioCapabilityForPagingOfNR.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.UERadioCapabilityForPagingOfEUTRA.Decode(r); err != nil {
			return
		}
	}
	return
}
