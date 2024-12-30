package ies

import "github.com/lvdund/ngap/aper"

type TraceActivation struct {
	NGRANTraceID                   []byte
	InterfacesToTrace              []byte
	TraceDepth                     TraceDepth
	TraceCollectionEntityIPAddress []byte
	// IEExtensions  *TraceActivationExtIEs
}

func (ie *TraceActivation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_NGRANTraceID := NewOCTETSTRING(ie.NGRANTraceID, aper.Constraint{Lb: 8, Ub: 8}, true)
	if err = tmp_NGRANTraceID.Encode(w); err != nil {
		return
	}
	tmp_InterfacesToTrace := NewBITSTRING(ie.InterfacesToTrace, aper.Constraint{Lb: 8, Ub: 8}, true)
	if err = tmp_InterfacesToTrace.Encode(w); err != nil {
		return
	}
	if err = ie.TraceDepth.Encode(w); err != nil {
		return
	}
	tmp_TraceCollectionEntityIPAddress := NewBITSTRING(ie.TraceCollectionEntityIPAddress, aper.Constraint{Lb: 1, Ub: 160}, true)
	if err = tmp_TraceCollectionEntityIPAddress.Encode(w); err != nil {
		return
	}
	return
}
func (ie *TraceActivation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_NGRANTraceID := OCTETSTRING{
		c:   aper.Constraint{Lb: 8, Ub: 8},
		ext: false,
	}
	if err = tmp_NGRANTraceID.Decode(r); err != nil {
		return
	}
	ie.NGRANTraceID = tmp_NGRANTraceID.Value
	tmp_InterfacesToTrace := BITSTRING{
		c:   aper.Constraint{Lb: 8, Ub: 8},
		ext: false,
	}
	if err = tmp_InterfacesToTrace.Decode(r); err != nil {
		return
	}
	ie.InterfacesToTrace = tmp_InterfacesToTrace.Value.Bytes
	if err = ie.TraceDepth.Decode(r); err != nil {
		return
	}
	tmp_TraceCollectionEntityIPAddress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: false,
	}
	if err = tmp_TraceCollectionEntityIPAddress.Decode(r); err != nil {
		return
	}
	ie.TraceCollectionEntityIPAddress = tmp_TraceCollectionEntityIPAddress.Value.Bytes
	return
}
