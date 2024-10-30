package ies

import "github.com/lvdund/ngap/aper"

type UPTransportLayerInformation struct {
	Choice    uint64
	GTPTunnel *GTPTunnel `True,,,`
	// ChoiceExtensions *UPTransportLayerInformationExtIEs `False,,,`
}

func (ie *UPTransportLayerInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteChoice(ie.Choice, 0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		err = ie.GTPTunnel.Encode(w)
	}
	return
}
func (ie *UPTransportLayerInformation) Decode(r *aper.AperReader) (err error) {
	if ie.Choice, err = r.ReadChoice(0, false); err != nil {
		return
	}
	switch ie.Choice {
	case 1:
		var tmp GTPTunnel
		if err = tmp.Decode(r); err != nil {
			return
		}
		ie.GTPTunnel = &tmp
	}
	return
}
