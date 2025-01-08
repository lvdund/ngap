package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	SONInformationPresentNothing uint64 = iota
	SONInformationPresentSoninformationrequest
	SONInformationPresentSoninformationreply
	SONInformationPresentChoiceExtensions
)

type SONInformation struct {
	Choice                uint64
	SONInformationRequest *SONInformationRequest
	SONInformationReply   *SONInformationReply
	// ChoiceExtensions *SONInformationExtIEs
}

func (ie *SONInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case SONInformationPresentSoninformationrequest:
		err = ie.SONInformationRequest.Encode(w)
	case SONInformationPresentSoninformationreply:
		err = ie.SONInformationReply.Encode(w)
	}
	return
}
func (ie *SONInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case SONInformationPresentSoninformationrequest:
		var tmp SONInformationRequest
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SONInformationRequest", err)
			return
		}
		ie.SONInformationRequest = &tmp
	case SONInformationPresentSoninformationreply:
		var tmp SONInformationReply
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read SONInformationReply", err)
			return
		}
		ie.SONInformationReply = &tmp
	}
	return
}
