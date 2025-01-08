package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type HandoverRequiredTransfer struct {
	DirectForwardingPathAvailability *DirectForwardingPathAvailability `optional`
	// IEExtensions *HandoverRequiredTransferExtIEs `optional`
}

func (ie *HandoverRequiredTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DirectForwardingPathAvailability != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.DirectForwardingPathAvailability != nil {
		if err = ie.DirectForwardingPathAvailability.Encode(w); err != nil {
			err = utils.WrapError("Read DirectForwardingPathAvailability", err)
			return
		}
	}
	return
}
func (ie *HandoverRequiredTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DirectForwardingPathAvailability.Decode(r); err != nil {
			err = utils.WrapError("Read DirectForwardingPathAvailability", err)
			return
		}
	}
	return
}
