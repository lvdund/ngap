package ies

import "github.com/lvdund/ngap/aper"

type ExpectedUEMovingTrajectoryItem struct {
	NGRANCGI         *NGRANCGI     `False,`
	TimeStayedInCell *aper.Integer `True,OPTIONAL`
	// IEExtensions ExpectedUEMovingTrajectoryItemExtIEs `False,OPTIONAL`
}

func (ie *ExpectedUEMovingTrajectoryItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.TimeStayedInCell != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.NGRANCGI != nil {
		if err = ie.NGRANCGI.Encode(w); err != nil {
			return
		}
	}
	if ie.TimeStayedInCell != nil {
		if err = w.WriteInteger(int64(*ie.TimeStayedInCell), &aper.Constraint{Lb: 0, Ub: 4095}, true); err != nil {
			return
		}
	}
	return
}
func (ie *ExpectedUEMovingTrajectoryItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.NGRANCGI = new(NGRANCGI)
	var v int64
	if err = ie.NGRANCGI.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4095}, true); err != nil {
			return
		} else {
			ie.TimeStayedInCell = (*aper.Integer)(&v)
		}
	}
	return
}
