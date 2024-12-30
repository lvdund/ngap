package ies

import "github.com/lvdund/ngap/aper"

const (
	CPTransportLayerInformationPresentNothing uint64 = iota
	CPTransportLayerInformationPresentEndpointipaddress
	CPTransportLayerInformationPresentChoiceExtensions
)

type CPTransportLayerInformation struct {
	Choice            uint64
	EndpointIPAddress *BITSTRING
	// ChoiceExtensions *CPTransportLayerInformationExtIEs
}

func (ie *CPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerInformationPresentEndpointipaddress:
		err = ie.EndpointIPAddress.Encode(w)
	}
	return
}
func (ie *CPTransportLayerInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerInformationPresentEndpointipaddress:
		var tmp BITSTRING
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EndpointIPAddress = &tmp
	}
	return
}
