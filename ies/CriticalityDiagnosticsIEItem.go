package ies

import "github.com/lvdund/ngap/aper"

type CriticalityDiagnosticsIEItem struct {
	IECriticality *Criticality  `False,`
	IEID          *ProtocolIEID `False,`
	TypeOfError   *TypeOfError  `False,`
	// IEExtensions CriticalityDiagnosticsIEItemExtIEs `False,OPTIONAL`
}

func (ie *CriticalityDiagnosticsIEItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.IECriticality != nil {
		if err = ie.IECriticality.Encode(w); err != nil {
			return
		}
	}
	if ie.IEID != nil {
		if err = ie.IEID.Encode(w); err != nil {
			return
		}
	}
	if ie.TypeOfError != nil {
		if err = ie.TypeOfError.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CriticalityDiagnosticsIEItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.IECriticality = new(Criticality)
	ie.IEID = new(ProtocolIEID)
	ie.TypeOfError = new(TypeOfError)
	if err = ie.IECriticality.Decode(r); err != nil {
		return
	}
	if err = ie.IEID.Decode(r); err != nil {
		return
	}
	if err = ie.TypeOfError.Decode(r); err != nil {
		return
	}
	return
}
