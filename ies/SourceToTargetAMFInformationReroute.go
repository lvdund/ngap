package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SourceToTargetAMFInformationReroute struct {
	ConfiguredNSSAI     []byte `lb:128,ub:128,optional`
	RejectedNSSAIinPLMN []byte `lb:32,ub:32,optional`
	RejectedNSSAIinTA   []byte `lb:32,ub:32,optional`
	// IEExtensions *SourceToTargetAMFInformationRerouteExtIEs `optional`
}

func (ie *SourceToTargetAMFInformationReroute) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
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
		tmp_ConfiguredNSSAI := NewOCTETSTRING(ie.ConfiguredNSSAI, aper.Constraint{Lb: 128, Ub: 128}, false)
		if err = tmp_ConfiguredNSSAI.Encode(w); err != nil {
			err = utils.WrapError("Encode ConfiguredNSSAI", err)
			return
		}
	}
	if ie.RejectedNSSAIinPLMN != nil {
		tmp_RejectedNSSAIinPLMN := NewOCTETSTRING(ie.RejectedNSSAIinPLMN, aper.Constraint{Lb: 32, Ub: 32}, false)
		if err = tmp_RejectedNSSAIinPLMN.Encode(w); err != nil {
			err = utils.WrapError("Encode RejectedNSSAIinPLMN", err)
			return
		}
	}
	if ie.RejectedNSSAIinTA != nil {
		tmp_RejectedNSSAIinTA := NewOCTETSTRING(ie.RejectedNSSAIinTA, aper.Constraint{Lb: 32, Ub: 32}, false)
		if err = tmp_RejectedNSSAIinTA.Encode(w); err != nil {
			err = utils.WrapError("Encode RejectedNSSAIinTA", err)
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
	if aper.IsBitSet(optionals, 1) {
		tmp_ConfiguredNSSAI := OCTETSTRING{
			c:   aper.Constraint{Lb: 128, Ub: 128},
			ext: false,
		}
		if err = tmp_ConfiguredNSSAI.Decode(r); err != nil {
			err = utils.WrapError("Read ConfiguredNSSAI", err)
			return
		}
		ie.ConfiguredNSSAI = tmp_ConfiguredNSSAI.Value
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_RejectedNSSAIinPLMN := OCTETSTRING{
			c:   aper.Constraint{Lb: 32, Ub: 32},
			ext: false,
		}
		if err = tmp_RejectedNSSAIinPLMN.Decode(r); err != nil {
			err = utils.WrapError("Read RejectedNSSAIinPLMN", err)
			return
		}
		ie.RejectedNSSAIinPLMN = tmp_RejectedNSSAIinPLMN.Value
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_RejectedNSSAIinTA := OCTETSTRING{
			c:   aper.Constraint{Lb: 32, Ub: 32},
			ext: false,
		}
		if err = tmp_RejectedNSSAIinTA.Decode(r); err != nil {
			err = utils.WrapError("Read RejectedNSSAIinTA", err)
			return
		}
		ie.RejectedNSSAIinTA = tmp_RejectedNSSAIinTA.Value
	}
	return
}
