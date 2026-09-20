package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
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
		err = fmt.Errorf("Encode QosFlowIdentifier: %w", err)
		return
	}
	if ie.DLForwarding != nil {
		if err = ie.DLForwarding.Encode(w); err != nil {
			err = fmt.Errorf("Encode DLForwarding: %w", err)
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
		err = fmt.Errorf("Read QosFlowIdentifier: %w", err)
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp := new(DLForwarding)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read DLForwarding: %w", err)
			return
		}
		ie.DLForwarding = tmp
	}
	return
}
