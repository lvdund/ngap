package ies

import "github.com/lvdund/ngap/aper"

type AllocationAndRetentionPriority struct {
	PriorityLevelARP        *PriorityLevelARP        `False,`
	PreemptionCapability    *PreemptionCapability    `False,`
	PreemptionVulnerability *PreemptionVulnerability `False,`
	// IEExtensions AllocationAndRetentionPriorityExtIEs `False,OPTIONAL`
}

func (ie *AllocationAndRetentionPriority) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PriorityLevelARP != nil {
		if err = ie.PriorityLevelARP.Encode(w); err != nil {
			return
		}
	}
	if ie.PreemptionCapability != nil {
		if err = ie.PreemptionCapability.Encode(w); err != nil {
			return
		}
	}
	if ie.PreemptionVulnerability != nil {
		if err = ie.PreemptionVulnerability.Encode(w); err != nil {
			return
		}
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
	ie.PriorityLevelARP = new(PriorityLevelARP)
	ie.PreemptionCapability = new(PreemptionCapability)
	ie.PreemptionVulnerability = new(PreemptionVulnerability)
	if err = ie.PriorityLevelARP.Decode(r); err != nil {
		return
	}
	if err = ie.PreemptionCapability.Decode(r); err != nil {
		return
	}
	if err = ie.PreemptionVulnerability.Decode(r); err != nil {
		return
	}
	return
}
