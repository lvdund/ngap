package ies

import "github.com/lvdund/ngap/aper"

type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier         *QosFlowIdentifier         `False,`
	QosFlowLevelQosParameters *QosFlowLevelQosParameters `True,OPTIONAL`
	ERABID                    *ERABID                    `False,OPTIONAL`
	// IEExtensions QosFlowAddOrModifyRequestItemExtIEs `False,OPTIONAL`
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
	if ie.QosFlowIdentifier != nil {
		if err = ie.QosFlowIdentifier.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowLevelQosParameters != nil {
		if err = ie.QosFlowLevelQosParameters.Encode(w); err != nil {
			return
		}
	}
	if ie.ERABID != nil {
		if err = ie.ERABID.Encode(w); err != nil {
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
	ie.QosFlowIdentifier = new(QosFlowIdentifier)
	ie.QosFlowLevelQosParameters = new(QosFlowLevelQosParameters)
	ie.ERABID = new(ERABID)
	if err = ie.QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.QosFlowLevelQosParameters.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.ERABID.Decode(r); err != nil {
			return
		}
	}
	return
}
