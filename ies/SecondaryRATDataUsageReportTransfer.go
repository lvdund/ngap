package ies

import "github.com/lvdund/ngap/aper"

type SecondaryRATDataUsageReportTransfer struct {
	SecondaryRATUsageInformation *SecondaryRATUsageInformation `False,OPTIONAL`
	// IEExtensions SecondaryRATDataUsageReportTransferExtIEs `False,OPTIONAL`
}

func (ie *SecondaryRATDataUsageReportTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SecondaryRATUsageInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.SecondaryRATUsageInformation != nil {
		if err = ie.SecondaryRATUsageInformation.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *SecondaryRATDataUsageReportTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.SecondaryRATUsageInformation = new(SecondaryRATUsageInformation)
	if aper.IsBitSet(optionals, 2) {
		if err = ie.SecondaryRATUsageInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
