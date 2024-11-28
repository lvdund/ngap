package ies

import "github.com/lvdund/ngap/aper"

const (
	UPTransportLayerInformationPresentNothing uint64 = iota /* No components present */
	UPTransportLayerInformationPresentGTPTunnel
	UPTransportLayerInformationPresentChoiceExtensions
)

type UPTransportLayerInformation struct {
	Choice    uint64
	GTPTunnel *GTPTunnel `True,,,`
	// ChoiceExtensions *UPTransportLayerInformationExtIEs `False,,,`
}

func (ie *UPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UPTransportLayerInformationPresentGTPTunnel:
		err = ie.GTPTunnel.Encode(w)
	}
	return
}
func (ie *UPTransportLayerInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return
	}
	switch ie.Choice {
	case UPTransportLayerInformationPresentGTPTunnel:
		var tmp GTPTunnel
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GTPTunnel = &tmp
	}
	return
}
