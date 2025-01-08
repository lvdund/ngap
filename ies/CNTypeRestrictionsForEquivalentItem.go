package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CNTypeRestrictionsForEquivalentItem struct {
	PlmnIdentity []byte
	CnType       int64
	// IEExtensions *CNTypeRestrictionsForEquivalentItemExtIEs `optional`
}

func (ie *CNTypeRestrictionsForEquivalentItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PlmnIdentity := NewOCTETSTRING(ie.PlmnIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PlmnIdentity.Encode(w); err != nil {
		err = utils.WrapError("Read PlmnIdentity", err)
		return
	}
	tmp_CnType := NewENUMERATED(ie.CnType, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_CnType.Encode(w); err != nil {
		err = utils.WrapError("Read CnType", err)
		return
	}
	return
}
func (ie *CNTypeRestrictionsForEquivalentItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PlmnIdentity := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_PlmnIdentity.Decode(r); err != nil {
		err = utils.WrapError("Read PlmnIdentity", err)
		return
	}
	ie.PlmnIdentity = tmp_PlmnIdentity.Value
	tmp_CnType := ENUMERATED{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_CnType.Decode(r); err != nil {
		err = utils.WrapError("Read CnType", err)
		return
	}
	ie.CnType = int64(tmp_CnType.Value)
	return
}
