package ies

import "github.com/lvdund/ngap/aper"

type CoreNetworkAssistanceInformationForInactive struct {
	UEIdentityIndexValue            *UEIdentityIndexValue            `False,`
	UESpecificDRX                   *PagingDRX                       `False,OPTIONAL`
	PeriodicRegistrationUpdateTimer *PeriodicRegistrationUpdateTimer `False,`
	MICOModeIndication              *MICOModeIndication              `False,OPTIONAL`
	TAIListForInactive              *TAIListForInactive              `False,`
	ExpectedUEBehaviour             *ExpectedUEBehaviour             `True,OPTIONAL`
	// IEExtensions CoreNetworkAssistanceInformationForInactiveExtIEs `False,OPTIONAL`
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
	if ie.UEIdentityIndexValue != nil {
		if err = ie.UEIdentityIndexValue.Encode(w); err != nil {
			return
		}
	}
	if ie.UESpecificDRX != nil {
		if err = ie.UESpecificDRX.Encode(w); err != nil {
			return
		}
	}
	if ie.PeriodicRegistrationUpdateTimer != nil {
		if err = ie.PeriodicRegistrationUpdateTimer.Encode(w); err != nil {
			return
		}
	}
	if ie.MICOModeIndication != nil {
		if err = ie.MICOModeIndication.Encode(w); err != nil {
			return
		}
	}
	if ie.TAIListForInactive != nil {
		if err = ie.TAIListForInactive.Encode(w); err != nil {
			return
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
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	ie.UEIdentityIndexValue = new(UEIdentityIndexValue)
	ie.UESpecificDRX = new(PagingDRX)
	ie.PeriodicRegistrationUpdateTimer = new(PeriodicRegistrationUpdateTimer)
	ie.MICOModeIndication = new(MICOModeIndication)
	ie.TAIListForInactive = new(TAIListForInactive)
	ie.ExpectedUEBehaviour = new(ExpectedUEBehaviour)
	if err = ie.UEIdentityIndexValue.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.UESpecificDRX.Decode(r); err != nil {
			return
		}
	}
	if err = ie.PeriodicRegistrationUpdateTimer.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.MICOModeIndication.Decode(r); err != nil {
			return
		}
	}
	if err = ie.TAIListForInactive.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.ExpectedUEBehaviour.Decode(r); err != nil {
			return
		}
	}
	return
}
