package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SecondaryRATDataUsageReportTransfer struct {
	SecondaryRATUsageInformation *SecondaryRATUsageInformation `optional`
	// IEExtensions *SecondaryRATDataUsageReportTransferExtIEs `optional`
}

func (ie *SecondaryRATDataUsageReportTransfer) Encode() (b []byte, err error) {
	w := aper.NewWriter(bytes.NewBuffer(b))
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
func (ie *SecondaryRATDataUsageReportTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
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
