package ies

import "github.com/lvdund/ngap/aper"

type UEassociatedLogicalNGconnectionItem struct {
	AMFUENGAPID *int64
	RANUENGAPID *int64
	// IEExtensions  *UEassociatedLogicalNGconnectionItemExtIEs
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
		tmp_AMFUENGAPID := NewINTEGER(*ie.AMFUENGAPID, aper.Constraint{Lb: 0, Ub: 1099511627775}, false)
		if err = tmp_AMFUENGAPID.Encode(w); err != nil {
			return
		}
	}
	if ie.RANUENGAPID != nil {
		tmp_RANUENGAPID := NewINTEGER(*ie.RANUENGAPID, aper.Constraint{Lb: 0, Ub: 4294967295}, false)
		if err = tmp_RANUENGAPID.Encode(w); err != nil {
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
	if aper.IsBitSet(optionals, 1) {
		tmp_AMFUENGAPID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
			ext: false,
		}
		if err = tmp_AMFUENGAPID.Decode(r); err != nil {
			return
		}
		*ie.AMFUENGAPID = int64(tmp_AMFUENGAPID.Value)
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_RANUENGAPID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 4294967295},
			ext: false,
		}
		if err = tmp_RANUENGAPID.Decode(r); err != nil {
			return
		}
		*ie.RANUENGAPID = int64(tmp_RANUENGAPID.Value)
	}
	return
}
