package ies

import "github.com/lvdund/ngap/aper"

type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier         int64
	QosFlowLevelQosParameters *QosFlowLevelQosParameters
	ERABID                    *int64
	// IEExtensions  *QosFlowAddOrModifyRequestItemExtIEs
}

func (ie *QosFlowAddOrModifyRequestItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.QosFlowLevelQosParameters != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ERABID != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	tmp_QosFlowIdentifier := NewINTEGER(ie.QosFlowIdentifier, aper.Constraint{Lb: 0, Ub: 63}, false)
	if err = tmp_QosFlowIdentifier.Encode(w); err != nil {
		return
	}
	if ie.QosFlowLevelQosParameters != nil {
		if err = ie.QosFlowLevelQosParameters.Encode(w); err != nil {
			return
		}
	}
	if ie.ERABID != nil {
		tmp_ERABID := NewINTEGER(*ie.ERABID, aper.Constraint{Lb: 0, Ub: 15}, false)
		if err = tmp_ERABID.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QosFlowAddOrModifyRequestItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	tmp_QosFlowIdentifier := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 63},
		ext: false,
	}
	if err = tmp_QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	ie.QosFlowIdentifier = int64(tmp_QosFlowIdentifier.Value)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.QosFlowLevelQosParameters.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_ERABID := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 15},
			ext: false,
		}
		if err = tmp_ERABID.Decode(r); err != nil {
			return
		}
		*ie.ERABID = int64(tmp_ERABID.Value)
	}
	return
}
