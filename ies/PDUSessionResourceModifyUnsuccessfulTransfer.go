package ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceModifyUnsuccessfulTransfer struct {
	Cause                  Cause                   `madatory`
	CriticalityDiagnostics *CriticalityDiagnostics `optional`
	// IEExtensions *PDUSessionResourceModifyUnsuccessfulTransferExtIEs `optional`
}

func (ie *PDUSessionResourceModifyUnsuccessfulTransfer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.CriticalityDiagnostics != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.Cause.Encode(w); err != nil {
		err = fmt.Errorf("Encode Cause: %w", err)
		return
	}
	if ie.CriticalityDiagnostics != nil {
		if err = ie.CriticalityDiagnostics.Encode(w); err != nil {
			err = fmt.Errorf("Encode CriticalityDiagnostics: %w", err)
			return
		}
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *PDUSessionResourceModifyUnsuccessfulTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.Cause.Decode(r); err != nil {
		err = fmt.Errorf("Read Cause: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(CriticalityDiagnostics)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read CriticalityDiagnostics: %w", err)
			return
		}
		ie.CriticalityDiagnostics = tmp
	}
	return
}
