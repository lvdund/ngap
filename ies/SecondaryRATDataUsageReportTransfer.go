package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SecondaryRATDataUsageReportTransfer struct {
	SecondaryRATUsageInformation *SecondaryRATUsageInformation `optional`
	// IEExtensions *SecondaryRATDataUsageReportTransferExtIEs `optional`
}

func (ie *SecondaryRATDataUsageReportTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.SecondaryRATUsageInformation != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.SecondaryRATUsageInformation != nil {
		if err = ie.SecondaryRATUsageInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode SecondaryRATUsageInformation", err)
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
	if aper.IsBitSet(optionals, 1) {
		if err = ie.SecondaryRATUsageInformation.Decode(r); err != nil {
			err = utils.WrapError("Read SecondaryRATUsageInformation", err)
			return
		}
	}
	return
}
