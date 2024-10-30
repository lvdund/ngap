package ies

import "github.com/lvdund/ngap/aper"

type CPTransportLayerInformation struct {
	Choice            uint64
	EndpointIPAddress *TransportLayerAddress `False,,,`
	// ChoiceExtensions *CPTransportLayerInformationExtIEs `False,,,`
}

func (ie *CPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.EndpointIPAddress.Encode(w)
	}
	return
}
func (ie *CPTransportLayerInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp TransportLayerAddress
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EndpointIPAddress = &tmp
	}
	return
}
