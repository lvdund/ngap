package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceNotifyReleasedTransfer struct {
	Cause Cause `madatory`
	// IEExtensions *PDUSessionResourceNotifyReleasedTransferExtIEs `optional`
}

func (ie *PDUSessionResourceNotifyReleasedTransfer) Encode(w *aper.AperWriter) (err error) {
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
func (ie *PDUSessionResourceNotifyReleasedTransfer) Decode(r *aper.AperReader) (err error) {
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
