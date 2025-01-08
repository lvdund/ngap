package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type BroadcastPLMNItem struct {
	PLMNIdentity        []byte
	TAISliceSupportList []SliceSupportItem
	// IEExtensions *BroadcastPLMNItemExtIEs `optional`
}

func (ie *BroadcastPLMNItem) Encode(w *aper.AperWriter) (err error) {
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
	if len(ie.TAISliceSupportList) > 0 {
		tmp := Sequence[*SliceSupportItem]{
			Value: []*SliceSupportItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
			ext:   false,
		}
		for _, i := range ie.TAISliceSupportList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Read TAISliceSupportList", err)
			return
		}
	}
	return
}
func (ie *BroadcastPLMNItem) Decode(r *aper.AperReader) (err error) {
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
	tmp_TAISliceSupportList := Sequence[*SliceSupportItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofSliceItems},
		ext: false,
	}
	fn := func() *SliceSupportItem { return new(SliceSupportItem) }
	if err = tmp_TAISliceSupportList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read TAISliceSupportList", err)
		return
	}
	ie.TAISliceSupportList = []SliceSupportItem{}
	for _, i := range tmp_TAISliceSupportList.Value {
		ie.TAISliceSupportList = append(ie.TAISliceSupportList, *i)
	}
	return
}
