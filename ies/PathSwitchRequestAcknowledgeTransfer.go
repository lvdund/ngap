package ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PathSwitchRequestAcknowledgeTransfer struct {
	ULNGUUPTNLInformation *UPTransportLayerInformation `optional`
	SecurityIndication    *SecurityIndication          `optional`
	// IEExtensions *PathSwitchRequestAcknowledgeTransferExtIEs `optional`
}

func (ie *PathSwitchRequestAcknowledgeTransfer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ULNGUUPTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SecurityIndication != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.ULNGUUPTNLInformation != nil {
		if err = ie.ULNGUUPTNLInformation.Encode(w); err != nil {
			err = fmt.Errorf("Encode ULNGUUPTNLInformation: %w", err)
			return
		}
	}
	if ie.SecurityIndication != nil {
		if err = ie.SecurityIndication.Encode(w); err != nil {
			err = fmt.Errorf("Encode SecurityIndication: %w", err)
			return
		}
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *PathSwitchRequestAcknowledgeTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(UPTransportLayerInformation)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read ULNGUUPTNLInformation: %w", err)
			return
		}
		ie.ULNGUUPTNLInformation = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(SecurityIndication)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read SecurityIndication: %w", err)
			return
		}
		ie.SecurityIndication = tmp
	}
	return
}
