package ies

import "github.com/lvdund/ngap/aper"

type PathSwitchRequestTransfer struct {
	DLNGUUPTNLInformation        *UPTransportLayerInformation  `False,`
	DLNGUTNLInformationReused    *DLNGUTNLInformationReused    `False,OPTIONAL`
	UserPlaneSecurityInformation *UserPlaneSecurityInformation `True,OPTIONAL`
	QosFlowAcceptedList          *QosFlowAcceptedList          `False,`
	// IEExtensions PathSwitchRequestTransferExtIEs `False,OPTIONAL`
}

func (ie *PathSwitchRequestTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
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
	if ie.DLNGUUPTNLInformation != nil {
		if err = ie.DLNGUUPTNLInformation.Encode(w); err != nil {
			return
		}
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
		if err = ie.QosFlowAcceptedList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PathSwitchRequestTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(3); err != nil {
		return
	}
	ie.DLNGUUPTNLInformation = new(UPTransportLayerInformation)
	ie.DLNGUTNLInformationReused = new(DLNGUTNLInformationReused)
	ie.UserPlaneSecurityInformation = new(UserPlaneSecurityInformation)
	ie.QosFlowAcceptedList = new(QosFlowAcceptedList)
	if err = ie.DLNGUUPTNLInformation.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.DLNGUTNLInformationReused.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.UserPlaneSecurityInformation.Decode(r); err != nil {
			return
		}
	}
	if err = ie.QosFlowAcceptedList.Decode(r); err != nil {
		return
	}
	return
}
