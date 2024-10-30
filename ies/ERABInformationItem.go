package ies

import "github.com/lvdund/ngap/aper"

type ERABInformationItem struct {
	ERABID       *ERABID       `False,`
	DLForwarding *DLForwarding `False,OPTIONAL`
	// IEExtensions ERABInformationItemExtIEs `False,OPTIONAL`
}

func (ie *ERABInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLForwarding != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.ERABID != nil {
		if err = ie.ERABID.Encode(w); err != nil {
			return
		}
	}
	if ie.DLForwarding != nil {
		if err = ie.DLForwarding.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *ERABInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.ERABID = new(ERABID)
	ie.DLForwarding = new(DLForwarding)
	if err = ie.ERABID.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.DLForwarding.Decode(r); err != nil {
			return
		}
	}
	return
}
