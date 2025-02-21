package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GUAMI struct {
	PLMNIdentity []byte         `lb:3,ub:3,madatory`
	AMFRegionID  aper.BitString `lb:8,ub:8,madatory`
	AMFSetID     aper.BitString `lb:10,ub:10,madatory`
	AMFPointer   aper.BitString `lb:6,ub:6,madatory`
	// IEExtensions *GUAMIExtIEs `optional`
}

func (ie *GUAMI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PLMNIdentity := NewOCTETSTRING(ie.PLMNIdentity, aper.Constraint{Lb: 3, Ub: 3}, false)
	if err = tmp_PLMNIdentity.Encode(w); err != nil {
		err = utils.WrapError("Encode PLMNIdentity", err)
		return
	}
	tmp_AMFRegionID := NewBITSTRING(ie.AMFRegionID, aper.Constraint{Lb: 8, Ub: 8}, false)
	if err = tmp_AMFRegionID.Encode(w); err != nil {
		err = utils.WrapError("Encode AMFRegionID", err)
		return
	}
	tmp_AMFSetID := NewBITSTRING(ie.AMFSetID, aper.Constraint{Lb: 10, Ub: 10}, false)
	if err = tmp_AMFSetID.Encode(w); err != nil {
		err = utils.WrapError("Encode AMFSetID", err)
		return
	}
	tmp_AMFPointer := NewBITSTRING(ie.AMFPointer, aper.Constraint{Lb: 6, Ub: 6}, false)
	if err = tmp_AMFPointer.Encode(w); err != nil {
		err = utils.WrapError("Encode AMFPointer", err)
		return
	}
	return
}
func (ie *GUAMI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PLMNIdentity := OCTETSTRING{
		c:   aper.Constraint{Lb: 3, Ub: 3},
		ext: false,
	}
	if err = tmp_PLMNIdentity.Decode(r); err != nil {
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	ie.PLMNIdentity = tmp_PLMNIdentity.Value
	tmp_AMFRegionID := BITSTRING{
		c:   aper.Constraint{Lb: 8, Ub: 8},
		ext: false,
	}
	if err = tmp_AMFRegionID.Decode(r); err != nil {
		err = utils.WrapError("Read AMFRegionID", err)
		return
	}
	ie.AMFRegionID = aper.BitString{Bytes: tmp_AMFRegionID.Value.Bytes, NumBits: tmp_AMFRegionID.Value.NumBits}
	tmp_AMFSetID := BITSTRING{
		c:   aper.Constraint{Lb: 10, Ub: 10},
		ext: false,
	}
	if err = tmp_AMFSetID.Decode(r); err != nil {
		err = utils.WrapError("Read AMFSetID", err)
		return
	}
	ie.AMFSetID = aper.BitString{Bytes: tmp_AMFSetID.Value.Bytes, NumBits: tmp_AMFSetID.Value.NumBits}
	tmp_AMFPointer := BITSTRING{
		c:   aper.Constraint{Lb: 6, Ub: 6},
		ext: false,
	}
	if err = tmp_AMFPointer.Decode(r); err != nil {
		err = utils.WrapError("Read AMFPointer", err)
		return
	}
	ie.AMFPointer = aper.BitString{Bytes: tmp_AMFPointer.Value.Bytes, NumBits: tmp_AMFPointer.Value.NumBits}
	return
}
