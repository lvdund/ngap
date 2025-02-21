package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

const (
	CPTransportLayerInformationPresentNothing uint64 = iota
	CPTransportLayerInformationPresentEndpointipaddress
	CPTransportLayerInformationPresentChoiceExtensions
)

type CPTransportLayerInformation struct {
	Choice            uint64
	EndpointIPAddress *aper.BitString
	// ChoiceExtensions *CPTransportLayerInformationExtIEs
}

func (ie *CPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case CPTransportLayerInformationPresentEndpointipaddress:
		tmp := NewBITSTRING(*ie.EndpointIPAddress, aper.Constraint{Lb: 1, Ub: 160}, false)
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
		tmp := BITSTRING{c: aper.Constraint{Lb: 1, Ub: 160}, ext: false}
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read EndpointIPAddress", err)
			return
		}
		ie.EndpointIPAddress = &aper.BitString{Bytes: tmp.Value.Bytes, NumBits: tmp.Value.NumBits}
	}
	return
}
