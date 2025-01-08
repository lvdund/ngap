package ies

import "github.com/lvdund/ngap/aper"

type EmergencyFallbackIndicator struct {
	EmergencyFallbackRequestIndicator EmergencyFallbackRequestIndicator
	EmergencyServiceTargetCN          *EmergencyServiceTargetCN `optional`
	// IEExtensions *EmergencyFallbackIndicatorExtIEs `optional`
}

func (ie *EmergencyFallbackIndicator) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.EmergencyServiceTargetCN != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if err = ie.EmergencyFallbackRequestIndicator.Encode(w); err != nil {
		return
	}
	if ie.EmergencyServiceTargetCN != nil {
		if err = ie.EmergencyServiceTargetCN.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *EmergencyFallbackIndicator) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	if err = ie.EmergencyFallbackRequestIndicator.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.EmergencyServiceTargetCN.Decode(r); err != nil {
			return
		}
	}
	return
}
