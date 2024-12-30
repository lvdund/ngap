package ies

import "github.com/lvdund/ngap/aper"

type CoreNetworkAssistanceInformationForInactive struct {
	UEIdentityIndexValue            UEIdentityIndexValue
	UESpecificDRX                   *PagingDRX
	PeriodicRegistrationUpdateTimer *[]byte
	MICOModeIndication              *MICOModeIndication
	TAIListForInactive              *[]TAIListForInactiveItem
	ExpectedUEBehaviour             *ExpectedUEBehaviour
	// IEExtensions  *CoreNetworkAssistanceInformationForInactiveExtIEs
}

func (ie *CoreNetworkAssistanceInformationForInactive) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.UESpecificDRX != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.PeriodicRegistrationUpdateTimer != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.MICOModeIndication != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.TAIListForInactive != nil {
		aper.SetBit(optionals, 4)
	}
	if ie.ExpectedUEBehaviour != nil {
		aper.SetBit(optionals, 5)
	}
	w.WriteBits(optionals, 6)
	if err = ie.UEIdentityIndexValue.Encode(w); err != nil {
		return
	}
	if ie.UESpecificDRX != nil {
		if err = ie.UESpecificDRX.Encode(w); err != nil {
			return
		}
	}
	if ie.PeriodicRegistrationUpdateTimer != nil {
		tmp_PeriodicRegistrationUpdateTimer := NewBITSTRING(*ie.PeriodicRegistrationUpdateTimer, aper.Constraint{Lb: 8, Ub: 8}, false)
		if err = tmp_PeriodicRegistrationUpdateTimer.Encode(w); err != nil {
			return
		}
	}
	if ie.MICOModeIndication != nil {
		if err = ie.MICOModeIndication.Encode(w); err != nil {
			return
		}
	}
	if ie.TAIListForInactive != nil {
		if len(*ie.TAIListForInactive) > 0 {
			tmp := Sequence[*TAIListForInactiveItem]{
				Value: []*TAIListForInactiveItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofTAIforInactive},
				ext:   false,
			}
			for _, i := range *ie.TAIListForInactive {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.ExpectedUEBehaviour != nil {
		if err = ie.ExpectedUEBehaviour.Encode(w); err != nil {
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
	if optionals, err = r.ReadBits(6); err != nil {
		return
	}
	if err = ie.UEIdentityIndexValue.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.UESpecificDRX.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_PeriodicRegistrationUpdateTimer := BITSTRING{
			c:   aper.Constraint{Lb: 8, Ub: 8},
			ext: false,
		}
		if err = tmp_PeriodicRegistrationUpdateTimer.Decode(r); err != nil {
			return
		}
		*ie.PeriodicRegistrationUpdateTimer = tmp_PeriodicRegistrationUpdateTimer.Value.Bytes
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.MICOModeIndication.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_TAIListForInactive := Sequence[*TAIListForInactiveItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofTAIforInactive},
			ext: false,
		}
		if err = tmp_TAIListForInactive.Decode(r); err != nil {
			return
		}
		ie.TAIListForInactive = &[]TAIListForInactiveItem{}
		for _, i := range tmp_TAIListForInactive.Value {
			*ie.TAIListForInactive = append(*ie.TAIListForInactive, *i)
		}
	}
	if aper.IsBitSet(optionals, 5) {
		if err = ie.ExpectedUEBehaviour.Decode(r); err != nil {
			return
		}
	}
	return
}
