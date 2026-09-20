package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         int64                     `lb:0,ub:63,madatory,valExt`
	QosFlowLevelQosParameters QosFlowLevelQosParameters `madatory`
	ERABID                    *int64                    `lb:0,ub:15,optional,valExt`
	// IEExtensions *QosFlowSetupRequestItemExtIEs `optional`
}

func (ie *QosFlowSetupRequestItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.ERABID != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	tmp_QosFlowIdentifier := NewINTEGER(ie.QosFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, true)
	if err = tmp_QosFlowIdentifier.Encode(w); err != nil {
		err = fmt.Errorf("Encode QosFlowIdentifier: %w", err)
		return
	}
	if err = ie.QosFlowLevelQosParameters.Encode(w); err != nil {
		err = fmt.Errorf("Encode QosFlowLevelQosParameters: %w", err)
		return
	}
	if ie.ERABID != nil {
		tmp_ERABID := NewINTEGER(*ie.ERABID, aper.Constraint{Lb: 0, Ub: 15}, true)
		if err = tmp_ERABID.Encode(w); err != nil {
			err = fmt.Errorf("Encode ERABID: %w", err)
			return
		}
	}
	return
}
func (ie *QosFlowSetupRequestItem) Decode(r *aper.AperReader) (err error) {
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
	if err = ie.QosFlowLevelQosParameters.Decode(r); err != nil {
		err = fmt.Errorf("Read QosFlowLevelQosParameters: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_ERABID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 15},
			ext: true,
		}
		if err = tmp_ERABID.Decode(r); err != nil {
			err = fmt.Errorf("Read ERABID: %w", err)
			return
		}
		ie.ERABID = (*int64)(&tmp_ERABID.Value)
	}
	return
}
