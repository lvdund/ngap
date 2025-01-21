package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type QosFlowInformationItem struct {
	QosFlowIdentifier int64         `lb:0,ub:63,madatory,valExt`
	DLForwarding      *DLForwarding `optional`
	// IEExtensions *QosFlowInformationItemExtIEs `optional`
}

func (ie *QosFlowInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLForwarding != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_QosFlowIdentifier := NewINTEGER(ie.QosFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, true)
	if err = tmp_QosFlowIdentifier.Encode(w); err != nil {
		err = utils.WrapError("Encode QosFlowIdentifier", err)
		return
	}
	if ie.DLForwarding != nil {
		if err = ie.DLForwarding.Encode(w); err != nil {
			err = utils.WrapError("Encode DLForwarding", err)
			return
		}
	}
	return
}
func (ie *QosFlowInformationItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	tmp_QosFlowIdentifier := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 63},
		ext: true,
	}
	if err = tmp_QosFlowIdentifier.Decode(r); err != nil {
		err = utils.WrapError("Read QosFlowIdentifier", err)
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(DLForwarding)
		if err = tmp.Decode(r); err != nil {
			err = utils.WrapError("Read DLForwarding", err)
			return
		}
		ie.DLForwarding = tmp
	}
	return
}
