package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type CNTypeRestrictionsForEquivalentItem struct {
	PlmnIdentity []byte `lb:3,ub:3,madatory`
	CnType       CnType `madatory`
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
		err = fmt.Errorf("Encode PlmnIdentity: %w", err)
		return
	}
	if err = ie.CnType.Encode(w); err != nil {
		err = fmt.Errorf("Encode CnType: %w", err)
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
		err = fmt.Errorf("Read PlmnIdentity: %w", err)
		return
	}
	ie.PlmnIdentity = tmp_PlmnIdentity.Value
	if err = ie.CnType.Decode(r); err != nil {
		err = fmt.Errorf("Read CnType: %w", err)
		return
	}
	return
}
