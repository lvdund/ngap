package ies

import "github.com/lvdund/ngap/aper"

type UEassociatedLogicalNGconnectionItem struct {
	AMFUENGAPID *AMFUENGAPID `False,OPTIONAL`
	RANUENGAPID *RANUENGAPID `False,OPTIONAL`
	// IEExtensions UEassociatedLogicalNGconnectionItemExtIEs `False,OPTIONAL`
}

func (ie *UEassociatedLogicalNGconnectionItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AMFUENGAPID != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.RANUENGAPID != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.AMFUENGAPID != nil {
		if err = ie.AMFUENGAPID.Encode(w); err != nil {
			return
		}
	}
	if ie.RANUENGAPID != nil {
		if err = ie.RANUENGAPID.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *UEassociatedLogicalNGconnectionItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	ie.AMFUENGAPID = new(AMFUENGAPID)
	ie.RANUENGAPID = new(RANUENGAPID)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.AMFUENGAPID.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.RANUENGAPID.Decode(r); err != nil {
			return
		}
	}
	return
}
