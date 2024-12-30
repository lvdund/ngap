package ies

import "github.com/lvdund/ngap/aper"

type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                      []byte
	PDUSessionResourceInformationList *[]PDUSessionResourceInformationItem
	ERABInformationList               *[]ERABInformationItem
	TargetCellID                      *NGRANCGI
	IndexToRFSP                       *int64
	UEHistoryInformation              *[]LastVisitedCellItem
	// IEExtensions  *SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs
}

func (ie *SourceNGRANNodeToTargetNGRANNodeTransparentContainer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PDUSessionResourceInformationList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ERABInformationList != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.TargetCellID != nil {
		aper.SetBit(optionals, 3)
	}
	if ie.IndexToRFSP != nil {
		aper.SetBit(optionals, 4)
	}
	if ie.UEHistoryInformation != nil {
		aper.SetBit(optionals, 5)
	}
	w.WriteBits(optionals, 6)
	tmp_RRCContainer := NewOCTETSTRING(ie.RRCContainer, aper.Constraint{Lb: 0, Ub: 0}, true)
	if err = tmp_RRCContainer.Encode(w); err != nil {
		return
	}
	if ie.PDUSessionResourceInformationList != nil {
		if len(*ie.PDUSessionResourceInformationList) > 0 {
			tmp := Sequence[*PDUSessionResourceInformationItem]{
				Value: []*PDUSessionResourceInformationItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
				ext:   false,
			}
			for _, i := range *ie.PDUSessionResourceInformationList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.ERABInformationList != nil {
		if len(*ie.ERABInformationList) > 0 {
			tmp := Sequence[*ERABInformationItem]{
				Value: []*ERABInformationItem{},
				c:     aper.Constraint{Lb: 0, Ub: 0},
				ext:   false,
			}
			for _, i := range *ie.ERABInformationList {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	if ie.TargetCellID != nil {
		if err = ie.TargetCellID.Encode(w); err != nil {
			return
		}
	}
	if ie.IndexToRFSP != nil {
		tmp_IndexToRFSP := NewINTEGER(*ie.IndexToRFSP, aper.Constraint{Lb: 1, Ub: 256}, true)
		if err = tmp_IndexToRFSP.Encode(w); err != nil {
			return
		}
	}
	if ie.UEHistoryInformation != nil {
		if len(*ie.UEHistoryInformation) > 0 {
			tmp := Sequence[*LastVisitedCellItem]{
				Value: []*LastVisitedCellItem{},
				c:     aper.Constraint{Lb: 1, Ub: maxnoofCellsinUEHistoryInfo},
				ext:   false,
			}
			for _, i := range *ie.UEHistoryInformation {
				tmp.Value = append(tmp.Value, &i)
			}
			if err = tmp.Encode(w); err != nil {
				return
			}
		}
	}
	return
}
func (ie *SourceNGRANNodeToTargetNGRANNodeTransparentContainer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(6); err != nil {
		return
	}
	tmp_RRCContainer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_RRCContainer.Decode(r); err != nil {
		return
	}
	ie.RRCContainer = tmp_RRCContainer.Value
	if aper.IsBitSet(optionals, 1) {
		tmp_PDUSessionResourceInformationList := Sequence[*PDUSessionResourceInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		if err = tmp_PDUSessionResourceInformationList.Decode(r); err != nil {
			return
		}
		ie.PDUSessionResourceInformationList = &[]PDUSessionResourceInformationItem{}
		for _, i := range tmp_PDUSessionResourceInformationList.Value {
			*ie.PDUSessionResourceInformationList = append(*ie.PDUSessionResourceInformationList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_ERABInformationList := Sequence[*ERABInformationItem]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_ERABInformationList.Decode(r); err != nil {
			return
		}
		ie.ERABInformationList = &[]ERABInformationItem{}
		for _, i := range tmp_ERABInformationList.Value {
			*ie.ERABInformationList = append(*ie.ERABInformationList, *i)
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.TargetCellID.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_IndexToRFSP := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: false,
		}
		if err = tmp_IndexToRFSP.Decode(r); err != nil {
			return
		}
		*ie.IndexToRFSP = int64(tmp_IndexToRFSP.Value)
	}
	if aper.IsBitSet(optionals, 5) {
		tmp_UEHistoryInformation := Sequence[*LastVisitedCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsinUEHistoryInfo},
			ext: false,
		}
		if err = tmp_UEHistoryInformation.Decode(r); err != nil {
			return
		}
		ie.UEHistoryInformation = &[]LastVisitedCellItem{}
		for _, i := range tmp_UEHistoryInformation.Value {
			*ie.UEHistoryInformation = append(*ie.UEHistoryInformation, *i)
		}
	}
	return
}
