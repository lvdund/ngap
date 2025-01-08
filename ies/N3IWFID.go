package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	N3IWFIDPresentNothing uint64 = iota
	N3IWFIDPresentN3IwfId
	N3IWFIDPresentChoiceExtensions
)

type N3IWFID struct {
	Choice  uint64
	N3IWFID []byte
	// ChoiceExtensions *N3IWFIDExtIEs
}

func (ie *N3IWFID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case N3IWFIDPresentN3IwfId:
		tmp := NewBITSTRING(ie.N3IWFID, aper.Constraint{Lb: 16, Ub: 16}, false)
		err = tmp.Encode(w)
	}
	return
}
func (ie *N3IWFID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case N3IWFIDPresentN3IwfId:
		tmp := NewBITSTRING(nil, aper.Constraint{Lb: 16, Ub: 16}, false)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read N3IWFID", err)
			return
		}
		ie.N3IWFID = tmp.Value.Bytes
	}
	return
}
