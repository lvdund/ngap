package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type ERABInformationItem struct {
	ERABID       int64
	DLForwarding *DLForwarding `optional`
	// IEExtensions *ERABInformationItemExtIEs `optional`
}

func (ie *ERABInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLForwarding != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_ERABID := NewINTEGER(ie.ERABID, aper.Constraint{Lb: 0, Ub: 15}, false)
	if err = tmp_ERABID.Encode(w); err != nil {
		err = utils.WrapError("Read ERABID", err)
		return
	}
	if ie.DLForwarding != nil {
		if err = ie.DLForwarding.Encode(w); err != nil {
			err = utils.WrapError("Read DLForwarding", err)
			return
		}
	}
	return
}
func (ie *ERABInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_ERABID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 15},
		ext: false,
	}
	if err = tmp_ERABID.Decode(r); err != nil {
		err = utils.WrapError("Read ERABID", err)
		return
	}
	ie.ERABID = int64(tmp_ERABID.Value)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DLForwarding.Decode(r); err != nil {
			err = utils.WrapError("Read DLForwarding", err)
			return
		}
	}
	return
}
