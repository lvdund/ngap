package ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceNotifyReleasedTransfer struct {
	Cause Cause `madatory`
	// IEExtensions *PDUSessionResourceNotifyReleasedTransferExtIEs `optional`
}

func (ie *PDUSessionResourceNotifyReleasedTransfer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.Cause.Encode(w); err != nil {
		err = fmt.Errorf("Encode Cause: %w", err)
		return
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *PDUSessionResourceNotifyReleasedTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.Cause.Decode(r); err != nil {
		err = fmt.Errorf("Read Cause: %w", err)
		return
	}
	return
}
