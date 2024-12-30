package ies

import "github.com/lvdund/ngap/aper"

type SupportedTAItem struct {
	TAC               []byte
	BroadcastPLMNList []BroadcastPLMNItem
	// IEExtensions  *SupportedTAItemExtIEs
}

func (ie *SupportedTAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_TAC := NewOCTETSTRING(ie.TAC, aper.Constraint{Lb: 3, Ub: 3}, true)
	if err = tmp_TAC.Encode(w); err != nil {
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
			return
		}
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
		return
	}
	ie.TAC = tmp_TAC.Value
	tmp_BroadcastPLMNList := Sequence[*BroadcastPLMNItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofBPLMNs},
		ext: false,
	}
	if err = tmp_BroadcastPLMNList.Decode(r); err != nil {
		return
	}
	ie.BroadcastPLMNList = []BroadcastPLMNItem{}
	for _, i := range tmp_BroadcastPLMNList.Value {
		ie.BroadcastPLMNList = append(ie.BroadcastPLMNList, *i)
	}
	return
}
