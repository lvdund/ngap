package ies

import "github.com/lvdund/ngap/aper"

type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                 `optional`
	TriggeringMessage         *TriggeringMessage             `optional`
	ProcedureCriticality      *Criticality                   `optional`
	IEsCriticalityDiagnostics []CriticalityDiagnosticsIEItem `optional`
	// IEExtensions *CriticalityDiagnosticsExtIEs `optional`
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
		if len(ie.IEsCriticalityDiagnostics) > 0 {
			tmp := Sequence[*CriticalityDiagnosticsIEItem]{
				Value: []*CriticalityDiagnosticsIEItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofErrors},
				ext:   false,
			}
			for _, i := range ie.IEsCriticalityDiagnostics {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
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
		tmp_IEsCriticalityDiagnostics := Sequence[*CriticalityDiagnosticsIEItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofErrors},
			ext: false,
		}
		fn := func() *CriticalityDiagnosticsIEItem { return new(CriticalityDiagnosticsIEItem) }
		if err = tmp_IEsCriticalityDiagnostics.Decode(r, fn); err != nil {
			return
		}
		ie.IEsCriticalityDiagnostics = []CriticalityDiagnosticsIEItem{}
		for _, i := range tmp_IEsCriticalityDiagnostics.Value {
			ie.IEsCriticalityDiagnostics = append(ie.IEsCriticalityDiagnostics, *i)
		}
	}
	return
}
