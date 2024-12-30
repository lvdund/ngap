package ies

import "github.com/lvdund/ngap/aper"

type AllocationAndRetentionPriority struct {
	PriorityLevelARP        int64
	PreemptionCapability    PreemptionCapability
	PreemptionVulnerability PreemptionVulnerability
	// IEExtensions  *AllocationAndRetentionPriorityExtIEs
}

func (ie *AllocationAndRetentionPriority) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PriorityLevelARP := NewINTEGER(ie.PriorityLevelARP, aper.Constraint{Lb: 1, Ub: 15}, true)
	if err = tmp_PriorityLevelARP.Encode(w); err != nil {
		return
	}
	if err = ie.PreemptionCapability.Encode(w); err != nil {
		return
	}
	if err = ie.PreemptionVulnerability.Encode(w); err != nil {
		return
	}
	return
}
func (ie *AllocationAndRetentionPriority) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PriorityLevelARP := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 15},
		ext: false,
	}
	if err = tmp_PriorityLevelARP.Decode(r); err != nil {
		return
	}
	ie.PriorityLevelARP = int64(tmp_PriorityLevelARP.Value)
	if err = ie.PreemptionCapability.Decode(r); err != nil {
		return
	}
	if err = ie.PreemptionVulnerability.Decode(r); err != nil {
		return
	}
	return
}
