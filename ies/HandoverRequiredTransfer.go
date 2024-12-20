package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
)

type HandoverRequiredTransfer struct {
	DirectForwardingPathAvailability *DirectForwardingPathAvailability `False,OPTIONAL`
	// IEExtensions HandoverRequiredTransferExtIEs `False,OPTIONAL`
}

func (ie *HandoverRequiredTransfer) Encode() (b []byte, err error) {
	w := aper.NewWriter(bytes.NewBuffer(b))
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
			return
		}
	}
	return
}
func (ie *HandoverRequiredTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.DirectForwardingPathAvailability = new(DirectForwardingPathAvailability)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DirectForwardingPathAvailability.Decode(r); err != nil {
			return
		}
	}
	return
}
