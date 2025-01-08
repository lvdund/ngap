package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type GUAMI struct {
	PLMNIdentity []byte
	AMFRegionID  []byte
	AMFSetID     []byte
	AMFPointer   []byte
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
		err = utils.WrapError("Read PLMNIdentity", err)
		return
	}
	tmp_AMFRegionID := NewBITSTRING(ie.AMFRegionID, aper.Constraint{Lb: 8, Ub: 8}, false)
	if err = tmp_AMFRegionID.Encode(w); err != nil {
		err = utils.WrapError("Read AMFRegionID", err)
		return
	}
	tmp_AMFSetID := NewBITSTRING(ie.AMFSetID, aper.Constraint{Lb: 10, Ub: 10}, false)
	if err = tmp_AMFSetID.Encode(w); err != nil {
		err = utils.WrapError("Read AMFSetID", err)
		return
	}
	tmp_AMFPointer := NewBITSTRING(ie.AMFPointer, aper.Constraint{Lb: 6, Ub: 6}, false)
	if err = tmp_AMFPointer.Encode(w); err != nil {
		err = utils.WrapError("Read AMFPointer", err)
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
	ie.AMFRegionID = tmp_AMFRegionID.Value.Bytes
	tmp_AMFSetID := BITSTRING{
		c:   aper.Constraint{Lb: 10, Ub: 10},
		ext: false,
	}
	if err = tmp_AMFSetID.Decode(r); err != nil {
		err = utils.WrapError("Read AMFSetID", err)
		return
	}
	ie.AMFSetID = tmp_AMFSetID.Value.Bytes
	tmp_AMFPointer := BITSTRING{
		c:   aper.Constraint{Lb: 6, Ub: 6},
		ext: false,
	}
	if err = tmp_AMFPointer.Decode(r); err != nil {
		err = utils.WrapError("Read AMFPointer", err)
		return
	}
	ie.AMFPointer = tmp_AMFPointer.Value.Bytes
	return
}
