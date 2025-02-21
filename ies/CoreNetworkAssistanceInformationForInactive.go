package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type CoreNetworkAssistanceInformationForInactive struct {
	UEIdentityIndexValue            UEIdentityIndexValue     `madatory`
	UESpecificDRX                   *PagingDRX               `optional`
	PeriodicRegistrationUpdateTimer aper.BitString           `lb:8,ub:8,madatory`
	MICOModeIndication              *MICOModeIndication      `optional`
	TAIListForInactive              []TAIListForInactiveItem `lb:1,ub:maxnoofTAIforInactive,madatory`
	ExpectedUEBehaviour             *ExpectedUEBehaviour     `optional`
	// IEExtensions *CoreNetworkAssistanceInformationForInactiveExtIEs `optional`
}

func (ie *CoreNetworkAssistanceInformationForInactive) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.UESpecificDRX != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.MICOModeIndication != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.ExpectedUEBehaviour != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if err = ie.UEIdentityIndexValue.Encode(w); err != nil {
		err = utils.WrapError("Encode UEIdentityIndexValue", err)
		return
	}
	if ie.UESpecificDRX != nil {
		if err = ie.UESpecificDRX.Encode(w); err != nil {
			err = utils.WrapError("Encode UESpecificDRX", err)
			return
		}
	}
	tmp_PeriodicRegistrationUpdateTimer := NewBITSTRING(ie.PeriodicRegistrationUpdateTimer, aper.Constraint{Lb: 8, Ub: 8}, false)
	if err = tmp_PeriodicRegistrationUpdateTimer.Encode(w); err != nil {
		err = utils.WrapError("Encode PeriodicRegistrationUpdateTimer", err)
		return
	}
	if ie.MICOModeIndication != nil {
		if err = ie.MICOModeIndication.Encode(w); err != nil {
			err = utils.WrapError("Encode MICOModeIndication", err)
			return
		}
	}
	if len(ie.TAIListForInactive) > 0 {
		tmp := Sequence[*TAIListForInactiveItem]{
			Value: []*TAIListForInactiveItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofTAIforInactive},
			ext:   false,
		}
		for _, i := range ie.TAIListForInactive {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode TAIListForInactive", err)
			return
		}
	} else {
		err = utils.WrapError("TAIListForInactive is nil", err)
		return
	}
	if ie.ExpectedUEBehaviour != nil {
		if err = ie.ExpectedUEBehaviour.Encode(w); err != nil {
			err = utils.WrapError("Encode ExpectedUEBehaviour", err)
			return
		}
	}
	return
}
func (ie *CoreNetworkAssistanceInformationForInactive) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if err = ie.UEIdentityIndexValue.Decode(r); err != nil {
		err = utils.WrapError("Read UEIdentityIndexValue", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(PagingDRX)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read UESpecificDRX", err)
			return
		}
		ie.UESpecificDRX = tmp
	}
	tmp_PeriodicRegistrationUpdateTimer := BITSTRING{
		c:   aper.Constraint{Lb: 8, Ub: 8},
		ext: false,
	}
	if err = tmp_PeriodicRegistrationUpdateTimer.Decode(r); err != nil {
		err = utils.WrapError("Read PeriodicRegistrationUpdateTimer", err)
		return
	}
	ie.PeriodicRegistrationUpdateTimer = aper.BitString{Bytes: tmp_PeriodicRegistrationUpdateTimer.Value.Bytes, NumBits: tmp_PeriodicRegistrationUpdateTimer.Value.NumBits}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(MICOModeIndication)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read MICOModeIndication", err)
			return
		}
		ie.MICOModeIndication = tmp
	}
	tmp_TAIListForInactive := Sequence[*TAIListForInactiveItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforInactive},
		ext: false,
	}
	fn := func() *TAIListForInactiveItem { return new(TAIListForInactiveItem) }
	if err = tmp_TAIListForInactive.Decode(r, fn); err != nil {
		err = utils.WrapError("Read TAIListForInactive", err)
		return
	}
	ie.TAIListForInactive = []TAIListForInactiveItem{}
	for _, i := range tmp_TAIListForInactive.Value {
		ie.TAIListForInactive = append(ie.TAIListForInactive, *i)
	}
	if aper.IsBitSet(optionals, 3) {
		tmp := new(ExpectedUEBehaviour)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read ExpectedUEBehaviour", err)
			return
		}
		ie.ExpectedUEBehaviour = tmp
	}
	return
}
