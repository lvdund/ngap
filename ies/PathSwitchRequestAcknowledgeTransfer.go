package ies

import "github.com/lvdund/ngap/aper"

type PathSwitchRequestAcknowledgeTransfer struct {
	ULNGUUPTNLInformation *UPTransportLayerInformation `False,OPTIONAL`
	SecurityIndication    *SecurityIndication          `True,OPTIONAL`
	// IEExtensions PathSwitchRequestAcknowledgeTransferExtIEs `False,OPTIONAL`
}

func (ie *PathSwitchRequestAcknowledgeTransfer) Encode(w *aper.AperWriter) (err error) {
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
			return
		}
	}
	if ie.SecurityIndication != nil {
		if err = ie.SecurityIndication.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PathSwitchRequestAcknowledgeTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	ie.ULNGUUPTNLInformation = new(UPTransportLayerInformation)
	ie.SecurityIndication = new(SecurityIndication)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.ULNGUUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.SecurityIndication.Decode(r); err != nil {
			return
		}
	}
	return
}
