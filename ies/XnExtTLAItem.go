package ies

import "github.com/lvdund/ngap/aper"

type XnExtTLAItem struct {
	IPsecTLA *TransportLayerAddress `False,OPTIONAL`
	GTPTLAs  *XnGTPTLAs             `False,OPTIONAL`
	// IEExtensions XnExtTLAItemExtIEs `False,OPTIONAL`
}

func (ie *XnExtTLAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.IPsecTLA != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.GTPTLAs != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.IPsecTLA != nil {
		if err = ie.IPsecTLA.Encode(w); err != nil {
			return
		}
	}
	if ie.GTPTLAs != nil {
		if err = ie.GTPTLAs.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *XnExtTLAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	ie.IPsecTLA = new(TransportLayerAddress)
	ie.GTPTLAs = new(XnGTPTLAs)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.IPsecTLA.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.GTPTLAs.Decode(r); err != nil {
			return
		}
	}
	return
}
