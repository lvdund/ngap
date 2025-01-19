package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceModifyIndicationUnsuccessfulTransfer struct {
	Cause Cause `madatory`
	// IEExtensions *PDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs `optional`
}

func (ie *PDUSessionResourceModifyIndicationUnsuccessfulTransfer) Encode() (b []byte, err error) {
	w := aper.NewWriter(bytes.NewBuffer(b))
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if err = ie.Cause.Encode(w); err != nil {
		err = utils.WrapError("Encode Cause", err)
		return
	}
	return
}
func (ie *PDUSessionResourceModifyIndicationUnsuccessfulTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	if err = ie.Cause.Decode(r); err != nil {
		err = utils.WrapError("Read Cause", err)
		return
	}
	return
}
