package ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type HandoverCommandTransfer struct {
	DLForwardingUPTNLInformation  *UPTransportLayerInformation    `optional`
	QosFlowToBeForwardedList      []QosFlowToBeForwardedItem      `lb:1,ub:maxnoofQosFlows,optional`
	DataForwardingResponseDRBList []DataForwardingResponseDRBItem `lb:1,ub:maxnoofDRBs,optional`
	// IEExtensions *HandoverCommandTransferExtIEs `optional`
}

func (ie *HandoverCommandTransfer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err = w.WriteBool(aper.Zero); err != nil {
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
			err = fmt.Errorf("Encode DLForwardingUPTNLInformation: %w", err)
			return
		}
	}
	if len(ie.QosFlowToBeForwardedList) > 0 {
		tmp := Sequence[*QosFlowToBeForwardedItem]{
			Value: []*QosFlowToBeForwardedItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowToBeForwardedList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode QosFlowToBeForwardedList: %w", err)
			return
		}
	}
	if len(ie.DataForwardingResponseDRBList) > 0 {
		tmp := Sequence[*DataForwardingResponseDRBItem]{
			Value: []*DataForwardingResponseDRBItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext:   false,
		}
		for _, i := range ie.DataForwardingResponseDRBList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode DataForwardingResponseDRBList: %w", err)
			return
		}
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *HandoverCommandTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(UPTransportLayerInformation)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read DLForwardingUPTNLInformation: %w", err)
			return
		}
		ie.DLForwardingUPTNLInformation = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_QosFlowToBeForwardedList := Sequence[*QosFlowToBeForwardedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		fn := func() *QosFlowToBeForwardedItem { return new(QosFlowToBeForwardedItem) }
		if err = tmp_QosFlowToBeForwardedList.Decode(r, fn); err != nil {
			err = fmt.Errorf("Read QosFlowToBeForwardedList: %w", err)
			return
		}
		ie.QosFlowToBeForwardedList = []QosFlowToBeForwardedItem{}
		for _, i := range tmp_QosFlowToBeForwardedList.Value {
			ie.QosFlowToBeForwardedList = append(ie.QosFlowToBeForwardedList, *i)
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_DataForwardingResponseDRBList := Sequence[*DataForwardingResponseDRBItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext: false,
		}
		fn := func() *DataForwardingResponseDRBItem { return new(DataForwardingResponseDRBItem) }
		if err = tmp_DataForwardingResponseDRBList.Decode(r, fn); err != nil {
			err = fmt.Errorf("Read DataForwardingResponseDRBList: %w", err)
			return
		}
		ie.DataForwardingResponseDRBList = []DataForwardingResponseDRBItem{}
		for _, i := range tmp_DataForwardingResponseDRBList.Value {
			ie.DataForwardingResponseDRBList = append(ie.DataForwardingResponseDRBList, *i)
		}
	}
	return
}
