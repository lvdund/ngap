package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SupportedTAItem struct {
	TAC               []byte              `lb:3,ub:3,madatory`
	BroadcastPLMNList []BroadcastPLMNItem `lb:1,ub:maxnoofBPLMNs,madatory`
	// IEExtensions *SupportedTAItemExtIEs `optional`
}

func (ie *SupportedTAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_TAC := NewOCTETSTRING(ie.TAC, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_TAC.Encode(w); err != nil {
		err = utils.WrapError("Encode TAC", err)
		return
	}
	if len(ie.BroadcastPLMNList) > 0 {
		tmp := Sequence[*BroadcastPLMNItem]{
			Value: []*BroadcastPLMNItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofBPLMNs},
			ext:   false,
		}
		for _, i := range ie.BroadcastPLMNList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode BroadcastPLMNList", err)
			return
		}
	} else {
		err = utils.WrapError("BroadcastPLMNList is nil", err)
		return
	}
	return
}
func (ie *SupportedTAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_TAC := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_TAC.Decode(r); err != nil {
		err = utils.WrapError("Read TAC", err)
		return
	}
	ie.TAC = tmp_TAC.Value
	tmp_BroadcastPLMNList := Sequence[*BroadcastPLMNItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofBPLMNs},
		ext: false,
	}
	fn := func() *BroadcastPLMNItem { return new(BroadcastPLMNItem) }
	if err = tmp_BroadcastPLMNList.Decode(r, fn); err != nil {
		err = utils.WrapError("Read BroadcastPLMNList", err)
		return
	}
	ie.BroadcastPLMNList = []BroadcastPLMNItem{}
	for _, i := range tmp_BroadcastPLMNList.Value {
		ie.BroadcastPLMNList = append(ie.BroadcastPLMNList, *i)
	}
	return
}
