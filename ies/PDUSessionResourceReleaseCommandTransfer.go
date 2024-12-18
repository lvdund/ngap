package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceReleaseCommandTransfer struct {
	Cause *Cause `False,`
	// IEExtensions PDUSessionResourceReleaseCommandTransferExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceReleaseCommandTransfer) Encode() (b []byte, err error) {
	w := aper.NewWriter(bytes.NewBuffer(b))
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.Cause != nil {
		if err = ie.Cause.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceReleaseCommandTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.Cause = new(Cause)
	if err = ie.Cause.Decode(r); err != nil {
		return
	}
	return
}
