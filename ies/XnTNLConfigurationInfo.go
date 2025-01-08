package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type XnTNLConfigurationInfo struct {
	XnTransportLayerAddresses         []TransportLayerAddress
	XnExtendedTransportLayerAddresses []XnExtTLAItem `optional`
	// IEExtensions *XnTNLConfigurationInfoExtIEs `optional`
}

func (ie *XnTNLConfigurationInfo) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.XnExtendedTransportLayerAddresses != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if len(ie.XnTransportLayerAddresses) > 0 {
		tmp := Sequence[*TransportLayerAddress]{
			Value: []*TransportLayerAddress{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofXnTLAs},
			ext:   false,
		}
		for _, i := range ie.XnTransportLayerAddresses {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Read XnTransportLayerAddresses", err)
			return
		}
	}
	if ie.XnExtendedTransportLayerAddresses != nil {
		if len(ie.XnExtendedTransportLayerAddresses) > 0 {
			tmp := Sequence[*XnExtTLAItem]{
				Value: []*XnExtTLAItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofXnExtTLAs},
				ext:   false,
			}
			for _, i := range ie.XnExtendedTransportLayerAddresses {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				err = utils.WrapError("Read XnExtendedTransportLayerAddresses", err)
				return
			}
		}
	}
	return
}
func (ie *XnTNLConfigurationInfo) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_XnTransportLayerAddresses := Sequence[*TransportLayerAddress]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofXnTLAs},
		ext: false,
	}
	fn := func() *TransportLayerAddress { return new(TransportLayerAddress) }
	if err = tmp_XnTransportLayerAddresses.Decode(r, fn); err != nil {
		err = utils.WrapError("Read XnTransportLayerAddresses", err)
		return
	}
	ie.XnTransportLayerAddresses = []TransportLayerAddress{}
	for _, i := range tmp_XnTransportLayerAddresses.Value {
		ie.XnTransportLayerAddresses = append(ie.XnTransportLayerAddresses, *i)
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_XnExtendedTransportLayerAddresses := Sequence[*XnExtTLAItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofXnExtTLAs},
			ext: false,
		}
		fn := func() *XnExtTLAItem { return new(XnExtTLAItem) }
		if err = tmp_XnExtendedTransportLayerAddresses.Decode(r, fn); err != nil {
			err = utils.WrapError("Read XnExtendedTransportLayerAddresses", err)
			return
		}
		ie.XnExtendedTransportLayerAddresses = []XnExtTLAItem{}
		for _, i := range tmp_XnExtendedTransportLayerAddresses.Value {
			ie.XnExtendedTransportLayerAddresses = append(ie.XnExtendedTransportLayerAddresses, *i)
		}
	}
	return
}
