package ies

import "github.com/lvdund/ngap/aper"

const (
	N3IWFIDPresentNothing uint64 = iota
	N3IWFIDPresentN3IwfId
	N3IWFIDPresentChoiceExtensions
)

type N3IWFID struct {
	Choice  uint64
	N3IWFID *[]byte
	// ChoiceExtensions *N3IWFIDExtIEs
}

func (ie *N3IWFID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case N3IWFIDPresentN3IwfId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
			Value: aper.BitString{
				Bytes:   *ie.N3IWFID,
				NumBits: uint64(len(*ie.N3IWFID)),
			},
		}
		err = tmp.Encode(w)
	}
	return
}
func (ie *N3IWFID) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case N3IWFIDPresentN3IwfId:
		tmp := BITSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.N3IWFID = &tmp.Value.Bytes
	}
	return
}
