package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   []QosFlowNotifyItem    `optional`
	QosFlowReleasedList []QosFlowWithCauseItem `optional`
	// IEExtensions *PDUSessionResourceNotifyTransferExtIEs `optional`
}

func (ie *PDUSessionResourceNotifyTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.QosFlowNotifyList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.QosFlowReleasedList != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if ie.QosFlowNotifyList != nil {
		if len(ie.QosFlowNotifyList) > 0 {
			tmp := Sequence[*QosFlowNotifyItem]{
				Value: []*QosFlowNotifyItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range ie.QosFlowNotifyList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.QosFlowReleasedList != nil {
		if len(ie.QosFlowReleasedList) > 0 {
			tmp := Sequence[*QosFlowWithCauseItem]{
				Value: []*QosFlowWithCauseItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range ie.QosFlowReleasedList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	return
}
func (ie *PDUSessionResourceNotifyTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp_QosFlowNotifyList := Sequence[*QosFlowNotifyItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowNotifyItem { return new(QosFlowNotifyItem) }
		if err = tmp_QosFlowNotifyList.Decode(r, fn); err != nil {
			return
		}
		ie.QosFlowNotifyList = []QosFlowNotifyItem{}
		for _, i := range tmp_QosFlowNotifyList.Value {
			ie.QosFlowNotifyList = append(ie.QosFlowNotifyList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_QosFlowReleasedList := Sequence[*QosFlowWithCauseItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowWithCauseItem { return new(QosFlowWithCauseItem) }
		if err = tmp_QosFlowReleasedList.Decode(r, fn); err != nil {
			return
		}
		ie.QosFlowReleasedList = []QosFlowWithCauseItem{}
		for _, i := range tmp_QosFlowReleasedList.Value {
			ie.QosFlowReleasedList = append(ie.QosFlowReleasedList, *i)
		}
	}
	return
}
