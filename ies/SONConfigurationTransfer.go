package ies

import "github.com/lvdund/ngap/aper"

type SONConfigurationTransfer struct {
	TargetRANNodeID        *TargetRANNodeID        `True,`
	SourceRANNodeID        *SourceRANNodeID        `True,`
	SONInformation         *SONInformation         `False,`
	XnTNLConfigurationInfo *XnTNLConfigurationInfo `True,OPTIONAL`
	// IEExtensions SONConfigurationTransferExtIEs `False,OPTIONAL`
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
	if ie.TargetRANNodeID != nil {
		if err = ie.TargetRANNodeID.Encode(w); err != nil {
			return
		}
	}
	if ie.SourceRANNodeID != nil {
		if err = ie.SourceRANNodeID.Encode(w); err != nil {
			return
		}
	}
	if ie.SONInformation != nil {
		if err = ie.SONInformation.Encode(w); err != nil {
			return
		}
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
	ie.TargetRANNodeID = new(TargetRANNodeID)
	ie.SourceRANNodeID = new(SourceRANNodeID)
	ie.SONInformation = new(SONInformation)
	ie.XnTNLConfigurationInfo = new(XnTNLConfigurationInfo)
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
