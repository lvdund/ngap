package ies

import "github.com/lvdund/ngap/aper"

type MobilityRestrictionList struct {
	ServingPLMN              []byte
	EquivalentPLMNs          []PLMNIdentity                 `optional`
	RATRestrictions          []RATRestrictionsItem          `optional`
	ForbiddenAreaInformation []ForbiddenAreaInformationItem `optional`
	ServiceAreaInformation   []ServiceAreaInformationItem   `optional`
	// IEExtensions *MobilityRestrictionListExtIEs `optional`
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
	tmp_ServingPLMN := NewOCTETSTRING(ie.ServingPLMN, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_ServingPLMN.Encode(w); err != nil {
		return
	}
	if ie.EquivalentPLMNs != nil {
		if len(ie.EquivalentPLMNs) > 0 {
			tmp := Sequence[*PLMNIdentity]{
				Value: []*PLMNIdentity{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofEPLMNs},
				ext:   false,
			}
			for _, i := range ie.EquivalentPLMNs {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.RATRestrictions != nil {
		if len(ie.RATRestrictions) > 0 {
			tmp := Sequence[*RATRestrictionsItem]{
				Value: []*RATRestrictionsItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofEPLMNsPlusOne},
				ext:   false,
			}
			for _, i := range ie.RATRestrictions {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.ForbiddenAreaInformation != nil {
		if len(ie.ForbiddenAreaInformation) > 0 {
			tmp := Sequence[*ForbiddenAreaInformationItem]{
				Value: []*ForbiddenAreaInformationItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofEPLMNsPlusOne},
				ext:   false,
			}
			for _, i := range ie.ForbiddenAreaInformation {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.ServiceAreaInformation != nil {
		if len(ie.ServiceAreaInformation) > 0 {
			tmp := Sequence[*ServiceAreaInformationItem]{
				Value: []*ServiceAreaInformationItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofEPLMNsPlusOne},
				ext:   false,
			}
			for _, i := range ie.ServiceAreaInformation {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
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
	tmp_ServingPLMN := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_ServingPLMN.Decode(r); err != nil {
		return
	}
	ie.ServingPLMN = tmp_ServingPLMN.Value
	if aper.IsBitSet(optionals, 1) {
		tmp_EquivalentPLMNs := Sequence[*PLMNIdentity]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEPLMNs},
			ext: false,
		}
		fn := func() *PLMNIdentity { return new(PLMNIdentity) }
		if err = tmp_EquivalentPLMNs.Decode(r, fn); err != nil {
			return
		}
		ie.EquivalentPLMNs = []PLMNIdentity{}
		for _, i := range tmp_EquivalentPLMNs.Value {
			ie.EquivalentPLMNs = append(ie.EquivalentPLMNs, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_RATRestrictions := Sequence[*RATRestrictionsItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEPLMNsPlusOne},
			ext: false,
		}
		fn := func() *RATRestrictionsItem { return new(RATRestrictionsItem) }
		if err = tmp_RATRestrictions.Decode(r, fn); err != nil {
			return
		}
		ie.RATRestrictions = []RATRestrictionsItem{}
		for _, i := range tmp_RATRestrictions.Value {
			ie.RATRestrictions = append(ie.RATRestrictions, *i)
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_ForbiddenAreaInformation := Sequence[*ForbiddenAreaInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEPLMNsPlusOne},
			ext: false,
		}
		fn := func() *ForbiddenAreaInformationItem { return new(ForbiddenAreaInformationItem) }
		if err = tmp_ForbiddenAreaInformation.Decode(r, fn); err != nil {
			return
		}
		ie.ForbiddenAreaInformation = []ForbiddenAreaInformationItem{}
		for _, i := range tmp_ForbiddenAreaInformation.Value {
			ie.ForbiddenAreaInformation = append(ie.ForbiddenAreaInformation, *i)
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_ServiceAreaInformation := Sequence[*ServiceAreaInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofEPLMNsPlusOne},
			ext: false,
		}
		fn := func() *ServiceAreaInformationItem { return new(ServiceAreaInformationItem) }
		if err = tmp_ServiceAreaInformation.Decode(r, fn); err != nil {
			return
		}
		ie.ServiceAreaInformation = []ServiceAreaInformationItem{}
		for _, i := range tmp_ServiceAreaInformation.Value {
			ie.ServiceAreaInformation = append(ie.ServiceAreaInformation, *i)
		}
	}
	return
}
