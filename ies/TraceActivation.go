package ies

import "github.com/lvdund/ngap/aper"

type TraceActivation struct {
	NGRANTraceID                   *NGRANTraceID          `False,`
	InterfacesToTrace              *InterfacesToTrace     `False,`
	TraceDepth                     *TraceDepth            `False,`
	TraceCollectionEntityIPAddress *TransportLayerAddress `False,`
	// IEExtensions TraceActivationExtIEs `False,OPTIONAL`
}

func (ie *TraceActivation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.NGRANTraceID != nil {
		if err = ie.NGRANTraceID.Encode(w); err != nil {
			return
		}
	}
	if ie.InterfacesToTrace != nil {
		if err = ie.InterfacesToTrace.Encode(w); err != nil {
			return
		}
	}
	if ie.TraceDepth != nil {
		if err = ie.TraceDepth.Encode(w); err != nil {
			return
		}
	}
	if ie.TraceCollectionEntityIPAddress != nil {
		if err = ie.TraceCollectionEntityIPAddress.Encode(w); err != nil {
			return
		}
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
	ie.NGRANTraceID = new(NGRANTraceID)
	ie.InterfacesToTrace = new(InterfacesToTrace)
	ie.TraceDepth = new(TraceDepth)
	ie.TraceCollectionEntityIPAddress = new(TransportLayerAddress)
	if err = ie.NGRANTraceID.Decode(r); err != nil {
		return
	}
	if err = ie.InterfacesToTrace.Decode(r); err != nil {
		return
	}
	if err = ie.TraceDepth.Decode(r); err != nil {
		return
	}
	if err = ie.TraceCollectionEntityIPAddress.Decode(r); err != nil {
		return
	}
	return
}
