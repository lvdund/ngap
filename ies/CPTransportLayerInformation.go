package ies

import "github.com/lvdund/ngap/aper"

const (
	CPTransportLayerInformationPresentNothing uint64 = iota
	CPTransportLayerInformationPresentEndpointipaddress
	CPTransportLayerInformationPresentChoiceExtensions
)

type CPTransportLayerInformation struct {
	Choice            uint64
	EndpointIPAddress []byte
	// ChoiceExtensions *CPTransportLayerInformationExtIEs
}

func (ie *CPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerInformationPresentEndpointipaddress:
		tmp := NewBITSTRING(ie.EndpointIPAddress, aper.Constraint{Lb: 1, Ub: 160}, false)
		err = tmp.Encode(w)
	}
	return
}
func (ie *CPTransportLayerInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerInformationPresentEndpointipaddress:
		tmp := NewBITSTRING(nil, aper.Constraint{Lb: 1, Ub: 160}, false)
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.EndpointIPAddress = tmp.Value.Bytes
	}
	return
}
