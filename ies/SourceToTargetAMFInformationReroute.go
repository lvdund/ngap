package ies

import "github.com/lvdund/ngap/aper"

type SourceToTargetAMFInformationReroute struct {
	ConfiguredNSSAI     *ConfiguredNSSAI     `False,OPTIONAL`
	RejectedNSSAIinPLMN *RejectedNSSAIinPLMN `False,OPTIONAL`
	RejectedNSSAIinTA   *RejectedNSSAIinTA   `False,OPTIONAL`
	// IEExtensions SourceToTargetAMFInformationRerouteExtIEs `False,OPTIONAL`
}

func (ie *SourceToTargetAMFInformationReroute) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ConfiguredNSSAI != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.RejectedNSSAIinPLMN != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.RejectedNSSAIinTA != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.ConfiguredNSSAI != nil {
		if err = ie.ConfiguredNSSAI.Encode(w); err != nil {
			return
		}
	}
	if ie.RejectedNSSAIinPLMN != nil {
		if err = ie.RejectedNSSAIinPLMN.Encode(w); err != nil {
			return
		}
	}
	if ie.RejectedNSSAIinTA != nil {
		if err = ie.RejectedNSSAIinTA.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *SourceToTargetAMFInformationReroute) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	ie.ConfiguredNSSAI = new(ConfiguredNSSAI)
	ie.RejectedNSSAIinPLMN = new(RejectedNSSAIinPLMN)
	ie.RejectedNSSAIinTA = new(RejectedNSSAIinTA)
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ConfiguredNSSAI.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.RejectedNSSAIinPLMN.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.RejectedNSSAIinTA.Decode(r); err != nil {
			return
		}
	}
	return
}
