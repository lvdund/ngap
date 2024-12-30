package ies

import "github.com/lvdund/ngap/aper"

type ExpectedUEBehaviour struct {
	ExpectedUEActivityBehaviour *ExpectedUEActivityBehaviour
	ExpectedHOInterval          *ExpectedHOInterval
	ExpectedUEMobility          *ExpectedUEMobility
	ExpectedUEMovingTrajectory  *[]ExpectedUEMovingTrajectoryItem
	// IEExtensions  *ExpectedUEBehaviourExtIEs
}

func (ie *ExpectedUEBehaviour) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ExpectedUEActivityBehaviour != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ExpectedHOInterval != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.ExpectedUEMobility != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.ExpectedUEMovingTrajectory != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if ie.ExpectedUEActivityBehaviour != nil {
		if err = ie.ExpectedUEActivityBehaviour.Encode(w); err != nil {
			return
		}
	}
	if ie.ExpectedHOInterval != nil {
		if err = ie.ExpectedHOInterval.Encode(w); err != nil {
			return
		}
	}
	if ie.ExpectedUEMobility != nil {
		if err = ie.ExpectedUEMobility.Encode(w); err != nil {
			return
		}
	}
	if ie.ExpectedUEMovingTrajectory != nil {
		if len(*ie.ExpectedUEMovingTrajectory) > 0 {
			tmp := Sequence[*ExpectedUEMovingTrajectoryItem]{
				Value: []*ExpectedUEMovingTrajectoryItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofCellsUEMovingTrajectory},
				ext:   false,
			}
			for _, i := range *ie.ExpectedUEMovingTrajectory {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	return
}
func (ie *ExpectedUEBehaviour) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.ExpectedUEActivityBehaviour.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ExpectedHOInterval.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.ExpectedUEMobility.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_ExpectedUEMovingTrajectory := Sequence[*ExpectedUEMovingTrajectoryItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsUEMovingTrajectory},
			ext: false,
		}
		if err = tmp_ExpectedUEMovingTrajectory.Decode(r); err != nil {
			return
		}
		ie.ExpectedUEMovingTrajectory = &[]ExpectedUEMovingTrajectoryItem{}
		for _, i := range tmp_ExpectedUEMovingTrajectory.Value {
			*ie.ExpectedUEMovingTrajectory = append(*ie.ExpectedUEMovingTrajectory, *i)
		}
	}
	return
}
