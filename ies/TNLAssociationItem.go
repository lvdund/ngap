package ies

import "github.com/lvdund/ngap/aper"

type TNLAssociationItem struct {
	TNLAssociationAddress *CPTransportLayerInformation `False,`
	Cause                 *Cause                       `False,`
	// IEExtensions TNLAssociationItemExtIEs `False,OPTIONAL`
}

func (ie *TNLAssociationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.TNLAssociationAddress != nil {
		if err = ie.TNLAssociationAddress.Encode(w); err != nil {
			return
		}
	}
	if ie.Cause != nil {
		if err = ie.Cause.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *TNLAssociationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.TNLAssociationAddress = new(CPTransportLayerInformation)
	ie.Cause = new(Cause)
	if err = ie.TNLAssociationAddress.Decode(r); err != nil {
		return
	}
	if err = ie.Cause.Decode(r); err != nil {
		return
	}
	return
}
