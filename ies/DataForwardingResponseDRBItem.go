package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type DataForwardingResponseDRBItem struct {
	DRBID                        int64                        `lb:1,ub:32,madatory,valExt`
	DLForwardingUPTNLInformation *UPTransportLayerInformation `optional`
	ULForwardingUPTNLInformation *UPTransportLayerInformation `optional`
	// IEExtensions *DataForwardingResponseDRBItemExtIEs `optional`
}

func (ie *DataForwardingResponseDRBItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLForwardingUPTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ULForwardingUPTNLInformation != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, true)
	if err = tmp_DRBID.Encode(w); err != nil {
		err = utils.WrapError("Encode DRBID", err)
		return
	}
	if ie.DLForwardingUPTNLInformation != nil {
		if err = ie.DLForwardingUPTNLInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode DLForwardingUPTNLInformation", err)
			return
		}
	}
	if ie.ULForwardingUPTNLInformation != nil {
		if err = ie.ULForwardingUPTNLInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode ULForwardingUPTNLInformation", err)
			return
		}
	}
	return
}
func (ie *DataForwardingResponseDRBItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	tmp_DRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: true,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		err = utils.WrapError("Read DRBID", err)
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(UPTransportLayerInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DLForwardingUPTNLInformation", err)
			return
		}
		ie.DLForwardingUPTNLInformation = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(UPTransportLayerInformation)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ULForwardingUPTNLInformation", err)
			return
		}
		ie.ULForwardingUPTNLInformation = tmp
	}
	return
}
