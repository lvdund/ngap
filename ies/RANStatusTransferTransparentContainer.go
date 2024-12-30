package ies

import "github.com/lvdund/ngap/aper"

type RANStatusTransferTransparentContainer struct {
	DRBsSubjectToStatusTransferList []DRBsSubjectToStatusTransferItem
	// IEExtensions  *RANStatusTransferTransparentContainerExtIEs
}

func (ie *RANStatusTransferTransparentContainer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if len(ie.DRBsSubjectToStatusTransferList) > 0 {
		tmp := Sequence[*DRBsSubjectToStatusTransferItem]{
			Value: []*DRBsSubjectToStatusTransferItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
			ext:   false,
		}
		for _, i := range ie.DRBsSubjectToStatusTransferList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
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
	tmp_DRBsSubjectToStatusTransferList := Sequence[*DRBsSubjectToStatusTransferItem]{
		c:   aper.Constraint{Lb: 1, Ub: maxnoofDRBs},
		ext: false,
	}
	if err = tmp_DRBsSubjectToStatusTransferList.Decode(r); err != nil {
		return
	}
	ie.DRBsSubjectToStatusTransferList = []DRBsSubjectToStatusTransferItem{}
	for _, i := range tmp_DRBsSubjectToStatusTransferList.Value {
		ie.DRBsSubjectToStatusTransferList = append(ie.DRBsSubjectToStatusTransferList, *i)
	}
	return
}
