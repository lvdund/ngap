package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type EndpointIPAddressAndPort struct {
	EndpointIPAddress []byte
	PortNumber        []byte
	// IEExtensions *EndpointIPAddressAndPortExtIEs `optional`
}

func (ie *EndpointIPAddressAndPort) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_EndpointIPAddress := NewBITSTRING(ie.EndpointIPAddress, aper.Constraint{Lb: 1, Ub: 160}, false)
	if err = tmp_EndpointIPAddress.Encode(w); err != nil {
		err = utils.WrapError("Read EndpointIPAddress", err)
		return
	}
	tmp_PortNumber := NewOCTETSTRING(ie.PortNumber, aper.Constraint{Lb: 2, Ub: 2}, false)
	if err = tmp_PortNumber.Encode(w); err != nil {
		err = utils.WrapError("Read PortNumber", err)
		return
	}
	return
}
func (ie *EndpointIPAddressAndPort) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_EndpointIPAddress := BITSTRING{
		c:   aper.Constraint{Lb: 1, Ub: 160},
		ext: false,
	}
	if err = tmp_EndpointIPAddress.Decode(r); err != nil {
		err = utils.WrapError("Read EndpointIPAddress", err)
		return
	}
	ie.EndpointIPAddress = tmp_EndpointIPAddress.Value.Bytes
	tmp_PortNumber := OCTETSTRING{
		c:   aper.Constraint{Lb: 2, Ub: 2},
		ext: false,
	}
	if err = tmp_PortNumber.Decode(r); err != nil {
		err = utils.WrapError("Read PortNumber", err)
		return
	}
	ie.PortNumber = tmp_PortNumber.Value
	return
}
