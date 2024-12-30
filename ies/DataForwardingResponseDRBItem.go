package ies

import "github.com/lvdund/ngap/aper"

type DataForwardingResponseDRBItem struct {
	DRBID                        int64
	DLForwardingUPTNLInformation *UPTransportLayerInformation
	ULForwardingUPTNLInformation *UPTransportLayerInformation
	// IEExtensions  *DataForwardingResponseDRBItemExtIEs
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
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_DRBID.Encode(w); err != nil {
		return
	}
	if ie.DLForwardingUPTNLInformation != nil {
		if err = ie.DLForwardingUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.ULForwardingUPTNLInformation != nil {
		if err = ie.ULForwardingUPTNLInformation.Encode(w); err != nil {
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
		ext: false,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DLForwardingUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ULForwardingUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
