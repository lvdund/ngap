package ies

import "github.com/lvdund/ngap/aper"

type QosFlowItemWithDataForwarding struct {
	QosFlowIdentifier      *QosFlowIdentifier      `False,`
	DataForwardingAccepted *DataForwardingAccepted `False,OPTIONAL`
	// IEExtensions QosFlowItemWithDataForwardingExtIEs `False,OPTIONAL`
}

func (ie *QosFlowItemWithDataForwarding) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DataForwardingAccepted != nil {
		aper.SetBit(optionals, 1)
	}
	w.WriteBits(optionals, 2)
	if ie.QosFlowIdentifier != nil {
		if err = ie.QosFlowIdentifier.Encode(w); err != nil {
			return
		}
	}
	if ie.DataForwardingAccepted != nil {
		if err = ie.DataForwardingAccepted.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *QosFlowItemWithDataForwarding) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(2); err != nil {
		return
	}
	ie.QosFlowIdentifier = new(QosFlowIdentifier)
	ie.DataForwardingAccepted = new(DataForwardingAccepted)
	if err = ie.QosFlowIdentifier.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DataForwardingAccepted.Decode(r); err != nil {
			return
		}
	}
	return
}
