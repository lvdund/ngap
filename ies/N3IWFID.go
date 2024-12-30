package ies

import "github.com/lvdund/ngap/aper"

const (
	N3IWFIDPresentNothing uint64 = iota
	N3IWFIDPresentN3IwfId
	N3IWFIDPresentChoiceExtensions
)

type N3IWFID struct {
	Choice  uint64
	N3IWFID *BITSTRING
	// ChoiceExtensions *N3IWFIDExtIEs
}

func (ie *N3IWFID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case N3IWFIDPresentN3IwfId:
		err = ie.N3IWFID.Encode(w)
	}
	return
}
func (ie *N3IWFID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case N3IWFIDPresentN3IwfId:
		var tmp BITSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.N3IWFID = &tmp
	}
	return
}
