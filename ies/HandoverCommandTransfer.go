package ies

import "github.com/lvdund/ngap/aper"

type HandoverCommandTransfer struct {
	DLForwardingUPTNLInformation  *UPTransportLayerInformation   `False,OPTIONAL`
	QosFlowToBeForwardedList      *QosFlowToBeForwardedList      `False,OPTIONAL`
	DataForwardingResponseDRBList *DataForwardingResponseDRBList `False,OPTIONAL`
	// IEExtensions HandoverCommandTransferExtIEs `False,OPTIONAL`
}

func (ie *HandoverCommandTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLForwardingUPTNLInformation != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.QosFlowToBeForwardedList != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.DataForwardingResponseDRBList != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.DLForwardingUPTNLInformation != nil {
		if err = ie.DLForwardingUPTNLInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowToBeForwardedList != nil {
		if err = ie.QosFlowToBeForwardedList.Encode(w); err != nil {
			return
		}
	}
	if ie.DataForwardingResponseDRBList != nil {
		if err = ie.DataForwardingResponseDRBList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *HandoverCommandTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	ie.DLForwardingUPTNLInformation = new(UPTransportLayerInformation)
	ie.QosFlowToBeForwardedList = new(QosFlowToBeForwardedList)
	ie.DataForwardingResponseDRBList = new(DataForwardingResponseDRBList)
	if aper.IsBitSet(optionals, 2) {
		if err = ie.DLForwardingUPTNLInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.QosFlowToBeForwardedList.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.DataForwardingResponseDRBList.Decode(r); err != nil {
			return
		}
	}
	return
}
