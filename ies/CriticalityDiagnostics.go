package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                 `optional`
	TriggeringMessage         *TriggeringMessage             `optional`
	ProcedureCriticality      *Criticality                   `optional`
	IEsCriticalityDiagnostics []CriticalityDiagnosticsIEItem `lb:1,ub:maxnoofErrors,optional`
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
			err = fmt.Errorf("Encode ProcedureCode: %w", err)
			return
		}
	}
	if ie.TriggeringMessage != nil {
		if err = ie.TriggeringMessage.Encode(w); err != nil {
			err = fmt.Errorf("Encode TriggeringMessage: %w", err)
			return
		}
	}
	if ie.ProcedureCriticality != nil {
		if err = ie.ProcedureCriticality.Encode(w); err != nil {
			err = fmt.Errorf("Encode ProcedureCriticality: %w", err)
			return
		}
	}
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
			err = fmt.Errorf("Encode IEsCriticalityDiagnostics: %w", err)
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
	if aper.IsBitSet(optionals, 1) {
		tmp := new(ProcedureCode)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read ProcedureCode: %w", err)
			return
		}
		ie.ProcedureCode = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(TriggeringMessage)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read TriggeringMessage: %w", err)
			return
		}
		ie.TriggeringMessage = tmp
	}
	if aper.IsBitSet(optionals, 3) {
		tmp := new(Criticality)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read ProcedureCriticality: %w", err)
			return
		}
		ie.ProcedureCriticality = tmp
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_IEsCriticalityDiagnostics := Sequence[*CriticalityDiagnosticsIEItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofErrors},
			ext: false,
		}
		fn := func() *CriticalityDiagnosticsIEItem { return new(CriticalityDiagnosticsIEItem) }
		if err = tmp_IEsCriticalityDiagnostics.Decode(r, fn); err != nil {
			err = fmt.Errorf("Read IEsCriticalityDiagnostics: %w", err)
			return
		}
		ie.IEsCriticalityDiagnostics = []CriticalityDiagnosticsIEItem{}
		for _, i := range tmp_IEsCriticalityDiagnostics.Value {
			ie.IEsCriticalityDiagnostics = append(ie.IEsCriticalityDiagnostics, *i)
		}
	}
	return
}
