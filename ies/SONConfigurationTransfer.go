package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
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
		err = utils.WrapError("Encode TargetRANNodeID", err)
		return
	}
	if err = ie.SourceRANNodeID.Encode(w); err != nil {
		err = utils.WrapError("Encode SourceRANNodeID", err)
		return
	}
	if err = ie.SONInformation.Encode(w); err != nil {
		err = utils.WrapError("Encode SONInformation", err)
		return
	}
	if ie.XnTNLConfigurationInfo != nil {
		if err = ie.XnTNLConfigurationInfo.Encode(w); err != nil {
			err = utils.WrapError("Encode XnTNLConfigurationInfo", err)
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
		err = utils.WrapError("Read TargetRANNodeID", err)
		return
	}
	if err = ie.SourceRANNodeID.Decode(r); err != nil {
		err = utils.WrapError("Read SourceRANNodeID", err)
		return
	}
	if err = ie.SONInformation.Decode(r); err != nil {
		err = utils.WrapError("Read SONInformation", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(XnTNLConfigurationInfo)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read XnTNLConfigurationInfo", err)
			return
		}
		ie.XnTNLConfigurationInfo = tmp
	}
	return
}
