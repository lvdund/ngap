package ies

import "github.com/lvdund/ngap/aper"

type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                `False,OPTIONAL`
	TriggeringMessage         *TriggeringMessage            `False,OPTIONAL`
	ProcedureCriticality      *Criticality                  `False,OPTIONAL`
	IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList `False,OPTIONAL`
	// IEExtensions CriticalityDiagnosticsExtIEs `False,OPTIONAL`
}

func (ie *CriticalityDiagnostics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ProcedureCode != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.TriggeringMessage != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.ProcedureCriticality != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.IEsCriticalityDiagnostics != nil {
		aper.SetBit(optionals, 4)
	}
	w.WriteBits(optionals, 5)
	if ie.ProcedureCode != nil {
		if err = ie.ProcedureCode.Encode(w); err != nil {
			return
		}
	}
	if ie.TriggeringMessage != nil {
		if err = ie.TriggeringMessage.Encode(w); err != nil {
			return
		}
	}
	if ie.ProcedureCriticality != nil {
		if err = ie.ProcedureCriticality.Encode(w); err != nil {
			return
		}
	}
	if ie.IEsCriticalityDiagnostics != nil {
		if err = ie.IEsCriticalityDiagnostics.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *CriticalityDiagnostics) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(5); err != nil {
		return
	}
	ie.ProcedureCode = new(ProcedureCode)
	ie.TriggeringMessage = new(TriggeringMessage)
	ie.ProcedureCriticality = new(Criticality)
	ie.IEsCriticalityDiagnostics = new(CriticalityDiagnosticsIEList)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.ProcedureCode.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.TriggeringMessage.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.ProcedureCriticality.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.IEsCriticalityDiagnostics.Decode(r); err != nil {
			return
		}
	}
	return
}
