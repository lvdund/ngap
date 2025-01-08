package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type UENGAPIDpair struct {
	AMFUENGAPID int64
	RANUENGAPID int64
	// IEExtensions *UENGAPIDpairExtIEs `optional`
}

func (ie *UENGAPIDpair) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_AMFUENGAPID := NewINTEGER(ie.AMFUENGAPID, aper.Constraint{Lb: 0, Ub: 1099511627775}, false)
	if err = tmp_AMFUENGAPID.Encode(w); err != nil {
		err = utils.WrapError("Read AMFUENGAPID", err)
		return
	}
	tmp_RANUENGAPID := NewINTEGER(ie.RANUENGAPID, aper.Constraint{Lb: 0, Ub: 4294967295}, false)
	if err = tmp_RANUENGAPID.Encode(w); err != nil {
		err = utils.WrapError("Read RANUENGAPID", err)
		return
	}
	return
}
func (ie *UENGAPIDpair) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_AMFUENGAPID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 1099511627775},
		ext: false,
	}
	if err = tmp_AMFUENGAPID.Decode(r); err != nil {
		err = utils.WrapError("Read AMFUENGAPID", err)
		return
	}
	ie.AMFUENGAPID = int64(tmp_AMFUENGAPID.Value)
	tmp_RANUENGAPID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4294967295},
		ext: false,
	}
	if err = tmp_RANUENGAPID.Decode(r); err != nil {
		err = utils.WrapError("Read RANUENGAPID", err)
		return
	}
	ie.RANUENGAPID = int64(tmp_RANUENGAPID.Value)
	return
}
