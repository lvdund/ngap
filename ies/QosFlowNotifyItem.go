package ies

import "github.com/lvdund/ngap/aper"

type QosFlowNotifyItem struct {
	QosFlowIdentifier *QosFlowIdentifier `False,`
	NotificationCause *NotificationCause `False,`
	// IEExtensions QosFlowNotifyItemExtIEs `False,OPTIONAL`
}

func (ie *QosFlowNotifyItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.QosFlowIdentifier != nil {
		if err = ie.QosFlowIdentifier.Encode(w); err != nil {
			return
		}
	}
	if ie.NotificationCause != nil {
		if err = ie.NotificationCause.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QosFlowNotifyItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.QosFlowIdentifier = new(QosFlowIdentifier)
	ie.NotificationCause = new(NotificationCause)
	if err = ie.QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	if err = ie.NotificationCause.Decode(r); err != nil {
		return
	}
	return
}
