package ies

import "github.com/lvdund/ngap/aper"

type TargetRANNodeID struct {
	GlobalRANNodeID *GlobalRANNodeID `False,`
	SelectedTAI     *TAI             `True,`
	// IEExtensions TargetRANNodeIDExtIEs `False,OPTIONAL`
}

func (ie *TargetRANNodeID) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.GlobalRANNodeID != nil {
		if err = ie.GlobalRANNodeID.Encode(w); err != nil {
			return
		}
	}
	if ie.SelectedTAI != nil {
		if err = ie.SelectedTAI.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TargetRANNodeID) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.GlobalRANNodeID = new(GlobalRANNodeID)
	ie.SelectedTAI = new(TAI)
	if err = ie.GlobalRANNodeID.Decode(r); err != nil {
		return
	}
	if err = ie.SelectedTAI.Decode(r); err != nil {
		return
	}
	return
}
