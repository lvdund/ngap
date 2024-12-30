package ies

import "github.com/lvdund/ngap/aper"

type SONConfigurationTransfer struct {
	TargetRANNodeID        TargetRANNodeID
	SourceRANNodeID        SourceRANNodeID
	SONInformation         SONInformation
	XnTNLConfigurationInfo *XnTNLConfigurationInfo
	// IEExtensions  *SONConfigurationTransferExtIEs
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
		return
	}
	if err = ie.SourceRANNodeID.Encode(w); err != nil {
		return
	}
	if err = ie.SONInformation.Encode(w); err != nil {
		return
	}
	if ie.XnTNLConfigurationInfo != nil {
		if err = ie.XnTNLConfigurationInfo.Encode(w); err != nil {
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
		return
	}
	if err = ie.SourceRANNodeID.Decode(r); err != nil {
		return
	}
	if err = ie.SONInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.XnTNLConfigurationInfo.Decode(r); err != nil {
			return
		}
	}
	return
}
