package ies

import "github.com/lvdund/ngap/aper"

const (
	CPTransportLayerInformationPresentNothing uint64 = iota /* No components present */
	CPTransportLayerInformationPresentEndpointIPAddress
	CPTransportLayerInformationPresentChoiceExtensions
)

type CPTransportLayerInformation struct {
	Choice            uint64
	EndpointIPAddress *TransportLayerAddress `False,,,`
	// ChoiceExtensions *CPTransportLayerInformationExtIEs `False,,,`
}

func (ie *CPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerInformationPresentEndpointIPAddress:
		err = ie.EndpointIPAddress.Encode(w)
	}
	return
}
func (ie *CPTransportLayerInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerInformationPresentEndpointIPAddress:
		var tmp TransportLayerAddress
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EndpointIPAddress = &tmp
	}
	return
}
