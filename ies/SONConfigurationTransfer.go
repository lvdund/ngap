package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type SONConfigurationTransfer struct {
	TargetRANNodeID        TargetRANNodeID         `madatory`
	SourceRANNodeID        SourceRANNodeID         `madatory`
	SONInformation         SONInformation          `madatory`
	XnTNLConfigurationInfo *XnTNLConfigurationInfo `optional`
	// IEExtensions *SONConfigurationTransferExtIEs `optional`
}

func (ie *SONConfigurationTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.XnTNLConfigurationInfo != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.TargetRANNodeID.Encode(w); err != nil {
		err = fmt.Errorf("Encode TargetRANNodeID: %w", err)
		return
	}
	if err = ie.SourceRANNodeID.Encode(w); err != nil {
		err = fmt.Errorf("Encode SourceRANNodeID: %w", err)
		return
	}
	if err = ie.SONInformation.Encode(w); err != nil {
		err = fmt.Errorf("Encode SONInformation: %w", err)
		return
	}
	if ie.XnTNLConfigurationInfo != nil {
		if err = ie.XnTNLConfigurationInfo.Encode(w); err != nil {
			err = fmt.Errorf("Encode XnTNLConfigurationInfo: %w", err)
			return
		}
	}
	return
}
func (ie *SONConfigurationTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.TargetRANNodeID.Decode(r); err != nil {
		err = fmt.Errorf("Read TargetRANNodeID: %w", err)
		return
	}
	if err = ie.SourceRANNodeID.Decode(r); err != nil {
		err = fmt.Errorf("Read SourceRANNodeID: %w", err)
		return
	}
	if err = ie.SONInformation.Decode(r); err != nil {
		err = fmt.Errorf("Read SONInformation: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(XnTNLConfigurationInfo)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read XnTNLConfigurationInfo: %w", err)
			return
		}
		ie.XnTNLConfigurationInfo = tmp
	}
	return
}
