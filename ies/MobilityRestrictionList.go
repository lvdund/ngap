package ies

import "github.com/lvdund/ngap/aper"

type MobilityRestrictionList struct {
	ServingPLMN              *PLMNIdentity             `False,`
	EquivalentPLMNs          *EquivalentPLMNs          `False,OPTIONAL`
	RATRestrictions          *RATRestrictions          `False,OPTIONAL`
	ForbiddenAreaInformation *ForbiddenAreaInformation `False,OPTIONAL`
	ServiceAreaInformation   *ServiceAreaInformation   `False,OPTIONAL`
	// IEExtensions MobilityRestrictionListExtIEs `False,OPTIONAL`
}

func (ie *MobilityRestrictionList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.EquivalentPLMNs != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.RATRestrictions != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.ForbiddenAreaInformation != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.ServiceAreaInformation != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if ie.ServingPLMN != nil {
		if err = ie.ServingPLMN.Encode(w); err != nil {
			return
		}
	}
	if ie.EquivalentPLMNs != nil {
		if err = ie.EquivalentPLMNs.Encode(w); err != nil {
			return
		}
	}
	if ie.RATRestrictions != nil {
		if err = ie.RATRestrictions.Encode(w); err != nil {
			return
		}
	}
	if ie.ForbiddenAreaInformation != nil {
		if err = ie.ForbiddenAreaInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.ServiceAreaInformation != nil {
		if err = ie.ServiceAreaInformation.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *MobilityRestrictionList) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	ie.ServingPLMN = new(PLMNIdentity)
	ie.EquivalentPLMNs = new(EquivalentPLMNs)
	ie.RATRestrictions = new(RATRestrictions)
	ie.ForbiddenAreaInformation = new(ForbiddenAreaInformation)
	ie.ServiceAreaInformation = new(ServiceAreaInformation)
	if err = ie.ServingPLMN.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.EquivalentPLMNs.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.RATRestrictions.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.ForbiddenAreaInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.ServiceAreaInformation.Decode(r); err != nil {
			return
		}
	}
	return
}
