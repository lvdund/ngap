package ies

import "github.com/lvdund/ngap/aper"

const (
	SONInformationPresentNothing uint64 = iota /* No components present */
	SONInformationPresentSONInformationRequest
	SONInformationPresentSONInformationReply
	SONInformationPresentChoiceExtensions
)

type SONInformation struct {
	Choice                uint64
	SONInformationRequest *SONInformationRequest `False,,,`
	SONInformationReply   *SONInformationReply   `True,,,`
	// ChoiceExtensions *SONInformationExtIEs `False,,,`
}

func (ie *SONInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return
	}
	switch ie.Choice {
	case SONInformationPresentSONInformationRequest:
		err = ie.SONInformationRequest.Encode(w)
	case SONInformationPresentSONInformationReply:
		err = ie.SONInformationReply.Encode(w)
	}
	return
}
func (ie *SONInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case SONInformationPresentSONInformationRequest:
		var tmp SONInformationRequest
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SONInformationRequest = &tmp
	case SONInformationPresentSONInformationReply:
		var tmp SONInformationReply
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SONInformationReply = &tmp
	}
	return
}
