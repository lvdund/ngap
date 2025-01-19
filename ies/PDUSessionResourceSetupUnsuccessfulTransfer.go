package ies

import (
	"bytes"

	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceSetupUnsuccessfulTransfer struct {
	Cause                  Cause                   `madatory`
	CriticalityDiagnostics *CriticalityDiagnostics `optional`
	// IEExtensions *PDUSessionResourceSetupUnsuccessfulTransferExtIEs `optional`
}

func (ie *PDUSessionResourceSetupUnsuccessfulTransfer) Encode() (b []byte, err error) {
	w := aper.NewWriter(bytes.NewBuffer(b))
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.CriticalityDiagnostics != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.Cause.Encode(w); err != nil {
		err = utils.WrapError("Encode Cause", err)
		return
	}
	if ie.CriticalityDiagnostics != nil {
		if err = ie.CriticalityDiagnostics.Encode(w); err != nil {
			err = utils.WrapError("Encode CriticalityDiagnostics", err)
			return
		}
	}
	return
}
func (ie *PDUSessionResourceSetupUnsuccessfulTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.Cause.Decode(r); err != nil {
		err = utils.WrapError("Read Cause", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.CriticalityDiagnostics.Decode(r); err != nil {
			err = utils.WrapError("Read CriticalityDiagnostics", err)
			return
		}
	}
	return
}
