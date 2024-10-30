package ies

import "github.com/lvdund/ngap/aper"

type SONInformation struct {
	Choice                uint64
	SONInformationRequest *SONInformationRequest `False,,,`
	SONInformationReply   *SONInformationReply   `True,,,`
	// ChoiceExtensions *SONInformationExtIEs `False,,,`
}

func (ie *SONInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.SONInformationRequest.Encode(w)
	case 2:
		err = ie.SONInformationReply.Encode(w)
	}
	return
}
func (ie *SONInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp SONInformationRequest
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SONInformationRequest = &tmp
	case 2:
		var tmp SONInformationReply
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.SONInformationReply = &tmp
	}
	return
}
