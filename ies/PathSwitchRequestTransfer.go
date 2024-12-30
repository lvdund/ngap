package ies

import "github.com/lvdund/ngap/aper"

type PathSwitchRequestTransfer struct {
	DLNGUUPTNLInformation        UPTransportLayerInformation
	DLNGUTNLInformationReused    *DLNGUTNLInformationReused
	UserPlaneSecurityInformation *UserPlaneSecurityInformation
	QosFlowAcceptedList          *[]QosFlowAcceptedItem
	// IEExtensions  *PathSwitchRequestTransferExtIEs
}

func (ie *PathSwitchRequestTransfer) Encode(w *aper.AperWriter) (err error) {
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
	if ie.QosFlowAcceptedList != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
		return
	}
	if ie.DLNGUTNLInformationReused != nil {
		if err = ie.DLNGUTNLInformationReused.Encode(w); err != nil {
			return
		}
	}
	if ie.UserPlaneSecurityInformation != nil {
		if err = ie.UserPlaneSecurityInformation.Encode(w); err != nil {
			return
		}
	}
	if ie.QosFlowAcceptedList != nil {
		if len(*ie.QosFlowAcceptedList) > 0 {
			tmp := Sequence[*QosFlowAcceptedItem]{
				Value: []*QosFlowAcceptedItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
				ext:   false,
			}
			for _, i := range *ie.QosFlowAcceptedList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	return
}
func (ie *PathSwitchRequestTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 1) {
		if err = ie.DLNGUTNLInformationReused.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.UserPlaneSecurityInformation.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_QosFlowAcceptedList := Sequence[*QosFlowAcceptedItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofQosFlows},
			ext: false,
		}
		if err = tmp_QosFlowAcceptedList.Decode(r); err != nil {
			return
		}
		ie.QosFlowAcceptedList = &[]QosFlowAcceptedItem{}
		for _, i := range tmp_QosFlowAcceptedList.Value {
			*ie.QosFlowAcceptedList = append(*ie.QosFlowAcceptedList, *i)
		}
	}
	return
}
