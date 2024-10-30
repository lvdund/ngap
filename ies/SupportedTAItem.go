package ies

import "github.com/lvdund/ngap/aper"

type SupportedTAItem struct {
	TAC               *TAC               `False,`
	BroadcastPLMNList *BroadcastPLMNList `False,`
	// IEExtensions SupportedTAItemExtIEs `False,OPTIONAL`
}

func (ie *SupportedTAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.TAC != nil {
		if err = ie.TAC.Encode(w); err != nil {
			return
		}
	}
	if ie.BroadcastPLMNList != nil {
		if err = ie.BroadcastPLMNList.Encode(w); err != nil {
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
	ie.TAC = new(TAC)
	ie.BroadcastPLMNList = new(BroadcastPLMNList)
	if err = ie.TAC.Decode(r); err != nil {
		return
	}
	if err = ie.BroadcastPLMNList.Decode(r); err != nil {
		return
	}
	return
}
