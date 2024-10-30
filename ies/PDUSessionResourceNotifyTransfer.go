package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   *QosFlowNotifyList    `False,OPTIONAL`
	QosFlowReleasedList *QosFlowListWithCause `False,OPTIONAL`
	// IEExtensions PDUSessionResourceNotifyTransferExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceNotifyTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
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
		if err = ie.QosFlowNotifyList.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowReleasedList != nil {
		if err = ie.QosFlowReleasedList.Encode(w); err != nil {
			return
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
	ie.QosFlowNotifyList = new(QosFlowNotifyList)
	ie.QosFlowReleasedList = new(QosFlowListWithCause)
	if aper.IsBitSet(optionals, 2) {
		if err = ie.QosFlowNotifyList.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.QosFlowReleasedList.Decode(r); err != nil {
			return
		}
	}
	return
}
