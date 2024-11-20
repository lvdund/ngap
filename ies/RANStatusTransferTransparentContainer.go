package ies

import "github.com/lvdund/ngap/aper"

type RANStatusTransferTransparentContainer struct {
	DRBsSubjectToStatusTransferList *DRBsSubjectToStatusTransferList `False,`
	// IEExtensions RANStatusTransferTransparentContainerExtIEs `False,OPTIONAL`
}

func (ie *RANStatusTransferTransparentContainer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.DRBsSubjectToStatusTransferList != nil {
		if err = ie.DRBsSubjectToStatusTransferList.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *RANStatusTransferTransparentContainer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.DRBsSubjectToStatusTransferList = new(DRBsSubjectToStatusTransferList)
	if err = ie.DRBsSubjectToStatusTransferList.Decode(r); err != nil {
		return
	}
	return
}
