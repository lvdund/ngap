package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ServiceAreaInformationItem struct {
	PLMNIdentity   []byte
	AllowedTACs    []TAC `optional`
	NotAllowedTACs []TAC `optional`
	// IEExtensions *ServiceAreaInformationItemExtIEs `optional`
}

func (ie *ServiceAreaInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.AllowedTACs != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.NotAllowedTACs != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	if ie.AllowedTACs != nil {
		if len(ie.AllowedTACs) > 0 {
			tmp := Sequence[*TAC]{
				Value: []*TAC{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofAllowedAreas},
				ext:   false,
			}
			for _, i := range ie.AllowedTACs {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				err = utils.WrapError("Read AllowedTACs", err)
				return
			}
		}
	}
	if ie.NotAllowedTACs != nil {
		if len(ie.NotAllowedTACs) > 0 {
			tmp := Sequence[*TAC]{
				Value: []*TAC{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofAllowedAreas},
				ext:   false,
			}
			for _, i := range ie.NotAllowedTACs {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				err = utils.WrapError("Read NotAllowedTACs", err)
				return
			}
		}
	}
	return
}
func (ie *ServiceAreaInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
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
	if aper.IsBitSet(optionals, 1) {
		tmp_AllowedTACs := Sequence[*TAC]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAllowedAreas},
			ext: false,
		}
		fn := func() *TAC { return new(TAC) }
		if err = tmp_AllowedTACs.Decode(r, fn); err != nil {
			err = utils.WrapError("Read AllowedTACs", err)
			return
		}
		ie.AllowedTACs = []TAC{}
		for _, i := range tmp_AllowedTACs.Value {
			ie.AllowedTACs = append(ie.AllowedTACs, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_NotAllowedTACs := Sequence[*TAC]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofAllowedAreas},
			ext: false,
		}
		fn := func() *TAC { return new(TAC) }
		if err = tmp_NotAllowedTACs.Decode(r, fn); err != nil {
			err = utils.WrapError("Read NotAllowedTACs", err)
			return
		}
		ie.NotAllowedTACs = []TAC{}
		for _, i := range tmp_NotAllowedTACs.Value {
			ie.NotAllowedTACs = append(ie.NotAllowedTACs, *i)
		}
	}
	return
}
