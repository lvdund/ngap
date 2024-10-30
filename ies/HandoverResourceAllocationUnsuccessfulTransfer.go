package ies

import "github.com/lvdund/ngap/aper"

type HandoverResourceAllocationUnsuccessfulTransfer struct {
	Cause                  *Cause                  `False,`
	CriticalityDiagnostics *CriticalityDiagnostics `True,OPTIONAL`
	// IEExtensions HandoverResourceAllocationUnsuccessfulTransferExtIEs `False,OPTIONAL`
}

func (ie *HandoverResourceAllocationUnsuccessfulTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.CriticalityDiagnostics != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.Cause != nil {
		if err = ie.Cause.Encode(w); err != nil {
			return
		}
	}
	if ie.CriticalityDiagnostics != nil {
		if err = ie.CriticalityDiagnostics.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *HandoverResourceAllocationUnsuccessfulTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.Cause = new(Cause)
	ie.CriticalityDiagnostics = new(CriticalityDiagnostics)
	if err = ie.Cause.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.CriticalityDiagnostics.Decode(r); err != nil {
			return
		}
	}
	return
}
