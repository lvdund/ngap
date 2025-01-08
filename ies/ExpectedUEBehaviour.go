package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ExpectedUEBehaviour struct {
	ExpectedUEActivityBehaviour *ExpectedUEActivityBehaviour     `optional`
	ExpectedHOInterval          *ExpectedHOInterval              `optional`
	ExpectedUEMobility          *ExpectedUEMobility              `optional`
	ExpectedUEMovingTrajectory  []ExpectedUEMovingTrajectoryItem `optional`
	// IEExtensions *ExpectedUEBehaviourExtIEs `optional`
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
			err = utils.WrapError("Read ExpectedUEActivityBehaviour", err)
			return
		}
	}
	if ie.ExpectedHOInterval != nil {
		if err = ie.ExpectedHOInterval.Encode(w); err != nil {
			err = utils.WrapError("Read ExpectedHOInterval", err)
			return
		}
	}
	if ie.ExpectedUEMobility != nil {
		if err = ie.ExpectedUEMobility.Encode(w); err != nil {
			err = utils.WrapError("Read ExpectedUEMobility", err)
			return
		}
	}
	if ie.ExpectedUEMovingTrajectory != nil {
		if len(ie.ExpectedUEMovingTrajectory) > 0 {
			tmp := Sequence[*ExpectedUEMovingTrajectoryItem]{
				Value: []*ExpectedUEMovingTrajectoryItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofCellsUEMovingTrajectory},
				ext:   false,
			}
			for _, i := range ie.ExpectedUEMovingTrajectory {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				err = utils.WrapError("Read ExpectedUEMovingTrajectory", err)
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
			err = utils.WrapError("Read ExpectedUEActivityBehaviour", err)
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ExpectedHOInterval.Decode(r); err != nil {
			err = utils.WrapError("Read ExpectedHOInterval", err)
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.ExpectedUEMobility.Decode(r); err != nil {
			err = utils.WrapError("Read ExpectedUEMobility", err)
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_ExpectedUEMovingTrajectory := Sequence[*ExpectedUEMovingTrajectoryItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsUEMovingTrajectory},
			ext: false,
		}
		fn := func() *ExpectedUEMovingTrajectoryItem { return new(ExpectedUEMovingTrajectoryItem) }
		if err = tmp_ExpectedUEMovingTrajectory.Decode(r, fn); err != nil {
			err = utils.WrapError("Read ExpectedUEMovingTrajectory", err)
			return
		}
		ie.ExpectedUEMovingTrajectory = []ExpectedUEMovingTrajectoryItem{}
		for _, i := range tmp_ExpectedUEMovingTrajectory.Value {
			ie.ExpectedUEMovingTrajectory = append(ie.ExpectedUEMovingTrajectory, *i)
		}
	}
	return
}
