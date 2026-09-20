package ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PathSwitchRequestTransfer struct {
	DLNGUUPTNLInformation        UPTransportLayerInformation   `madatory`
	DLNGUTNLInformationReused    *DLNGUTNLInformationReused    `optional`
	UserPlaneSecurityInformation *UserPlaneSecurityInformation `optional`
	QosFlowAcceptedList          []QosFlowAcceptedItem         `lb:1,ub:maxnoofQosFlows,madatory`
	// IEExtensions *PathSwitchRequestTransferExtIEs `optional`
}

func (ie *PathSwitchRequestTransfer) Encode() (b []byte, err error) {
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.DLNGUTNLInformationReused != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.UserPlaneSecurityInformation != nil {
		aper.SetBit(optionals, 2)
	}
	w.WriteBits(optionals, 3)
	if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
		err = fmt.Errorf("Encode DLNGUUPTNLInformation: %w", err)
		return
	}
	if ie.DLNGUTNLInformationReused != nil {
		if err = ie.DLNGUTNLInformationReused.Encode(w); err != nil {
			err = fmt.Errorf("Encode DLNGUTNLInformationReused: %w", err)
			return
		}
	}
	if ie.UserPlaneSecurityInformation != nil {
		if err = ie.UserPlaneSecurityInformation.Encode(w); err != nil {
			err = fmt.Errorf("Encode UserPlaneSecurityInformation: %w", err)
			return
		}
	}
	if len(ie.QosFlowAcceptedList) > 0 {
		tmp := Sequence[*QosFlowAcceptedItem]{
			Value: []*QosFlowAcceptedItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext:   false,
		}
		for _, i := range ie.QosFlowAcceptedList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = fmt.Errorf("Encode QosFlowAcceptedList: %w", err)
			return
		}
	} else {
		if err != nil {
			err = fmt.Errorf("QosFlowAcceptedList is nil: %w", err)
		}
		return
	}
	err = w.Close()
	b = buf.Bytes()
	return
}
func (ie *PathSwitchRequestTransfer) Decode(wire []byte) (err error) {
	r := aper.NewReader(bytes.NewBuffer(wire))
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		err = fmt.Errorf("Read DLNGUUPTNLInformation: %w", err)
		return
	}
	if aper.IsBitSet(optionals, 1) {
		tmp := new(DLNGUTNLInformationReused)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read DLNGUTNLInformationReused: %w", err)
			return
		}
		ie.DLNGUTNLInformationReused = tmp
	}
	if aper.IsBitSet(optionals, 2) {
		tmp := new(UserPlaneSecurityInformation)
		if err = tmp.Decode(r); err != nil {
			err = fmt.Errorf("Read UserPlaneSecurityInformation: %w", err)
			return
		}
		ie.UserPlaneSecurityInformation = tmp
	}
	tmp_QosFlowAcceptedList := Sequence[*QosFlowAcceptedItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
		ext: false,
	}
	fn := func() *QosFlowAcceptedItem { return new(QosFlowAcceptedItem) }
	if err = tmp_QosFlowAcceptedList.Decode(r, fn); err != nil {
		err = fmt.Errorf("Read QosFlowAcceptedList: %w", err)
		return
	}
	ie.QosFlowAcceptedList = []QosFlowAcceptedItem{}
	for _, i := range tmp_QosFlowAcceptedList.Value {
		ie.QosFlowAcceptedList = append(ie.QosFlowAcceptedList, *i)
	}
	return
}
