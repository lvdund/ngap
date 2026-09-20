package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type XnExtTLAItem struct {
	IPsecTLA *aper.BitString         `lb:1,ub:160,optional,valExt`
	GTPTLAs  []TransportLayerAddress `lb:1,ub:maxnoofXnGTPTLAs,optional`
	// IEExtensions *XnExtTLAItemExtIEs `optional`
}

func (ie *XnExtTLAItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.IPsecTLA != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.GTPTLAs != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.IPsecTLA != nil {
		tmp_IPsecTLA := NewBITSTRING(*ie.IPsecTLA, aper.Constraint{Lb: 1, Ub: 160}, true)
		if err = tmp_IPsecTLA.Encode(w); err != nil {
			err = fmt.Errorf("Encode IPsecTLA: %w", err)
			return
		}
	}
	if len(ie.GTPTLAs) > 0 {
		tmp := Sequence[*TransportLayerAddress]{
			Value: []*TransportLayerAddress{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofXnGTPTLAs},
			ext:   false,
		}
		for _, i := range ie.GTPTLAs {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode GTPTLAs: %w", err)
			return
		}
	}
	return
}
func (ie *XnExtTLAItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_IPsecTLA := BITSTRING{
			c:   aper.Constraint{Lb: 1, Ub: 160},
			ext: true,
		}
		if err = tmp_IPsecTLA.Decode(r); err != nil {
			err = fmt.Errorf("Read IPsecTLA: %w", err)
			return
		}
		ie.IPsecTLA = &aper.BitString{
			Bytes:   tmp_IPsecTLA.Value.Bytes,
			NumBits: tmp_IPsecTLA.Value.NumBits,
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_GTPTLAs := Sequence[*TransportLayerAddress]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofXnGTPTLAs},
			ext: false,
		}
		fn := func() *TransportLayerAddress { return new(TransportLayerAddress) }
		if err = tmp_GTPTLAs.Decode(r, fn); err != nil {
			err = fmt.Errorf("Read GTPTLAs: %w", err)
			return
		}
		ie.GTPTLAs = []TransportLayerAddress{}
		for _, i := range tmp_GTPTLAs.Value {
			ie.GTPTLAs = append(ie.GTPTLAs, *i)
		}
	}
	return
}
