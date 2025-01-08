package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ForbiddenAreaInformationItem struct {
	PLMNIdentity  []byte
	ForbiddenTACs []TAC
	// IEExtensions *ForbiddenAreaInformationItemExtIEs `optional`
}

func (ie *ForbiddenAreaInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	if len(ie.ForbiddenTACs) > 0 {
		tmp := Sequence[*TAC]{
			Value: []*TAC{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofForbTACs},
			ext:   false,
		}
		for _, i := range ie.ForbiddenTACs {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Read ForbiddenTACs", err)
			return
		}
	}
	return
}
func (ie *ForbiddenAreaInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PLMNIdentity := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_PLMNIdentity.Decode(r); err != nil {
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	tmp_ForbiddenTACs := Sequence[*TAC]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofForbTACs},
		ext: false,
	}
	fn := func() *TAC { return new(TAC) }
	if err = tmp_ForbiddenTACs.Decode(r, fn); err != nil {
		err = utils.WrapError("Read ForbiddenTACs", err)
		return
	}
	ie.ForbiddenTACs = []TAC{}
	for _, i := range tmp_ForbiddenTACs.Value {
		ie.ForbiddenTACs = append(ie.ForbiddenTACs, *i)
	}
	return
}
