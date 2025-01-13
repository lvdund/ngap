package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                      []byte                              `lb:0,ub:0,madatory`
	PDUSessionResourceInformationList []PDUSessionResourceInformationItem `lb:1,ub:maxnoofPDUSessions,optional`
	ERABInformationList               []ERABInformationItem               `lb:1,ub:maxnoofERABs,optional`
	TargetCellID                      *NGRANCGI                           `optional`
	IndexToRFSP                       *int64                              `lb:1,ub:256,optional,valExt`
	UEHistoryInformation              []LastVisitedCellItem               `lb:1,ub:maxnoofCellsinUEHistoryInfo,optional`
	// IEExtensions *SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs `optional`
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
	tmp_RRCContainer := NewOCTETSTRING(ie.RRCContainer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_RRCContainer.Encode(w); err != nil {
		err = utils.WrapError("Encode RRCContainer", err)
		return
	}
	if len(ie.PDUSessionResourceInformationList) > 0 {
		tmp := Sequence[*PDUSessionResourceInformationItem]{
			Value: []*PDUSessionResourceInformationItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext:   false,
		}
		for _, i := range ie.PDUSessionResourceInformationList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode PDUSessionResourceInformationList", err)
			return
		}
	}
	if len(ie.ERABInformationList) > 0 {
		tmp := Sequence[*ERABInformationItem]{
			Value: []*ERABInformationItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofERABs},
			ext:   false,
		}
		for _, i := range ie.ERABInformationList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode ERABInformationList", err)
			return
		}
	}
	if ie.TargetCellID != nil {
		if err = ie.TargetCellID.Encode(w); err != nil {
			err = utils.WrapError("Encode TargetCellID", err)
			return
		}
	}
	if ie.IndexToRFSP != nil {
		tmp_IndexToRFSP := NewINTEGER(*ie.IndexToRFSP, aper.Constraint{Lb: 1, Ub: 256}, true)
		if err = tmp_IndexToRFSP.Encode(w); err != nil {
			err = utils.WrapError("Encode IndexToRFSP", err)
			return
		}
	}
	if len(ie.UEHistoryInformation) > 0 {
		tmp := Sequence[*LastVisitedCellItem]{
			Value: []*LastVisitedCellItem{},
			c:     aper.Constraint{Lb: 1, Ub: maxnoofCellsinUEHistoryInfo},
			ext:   false,
		}
		for _, i := range ie.UEHistoryInformation {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode UEHistoryInformation", err)
			return
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
		err = utils.WrapError("Read RRCContainer", err)
		return
	}
	ie.RRCContainer = tmp_RRCContainer.Value
	if aper.IsBitSet(optionals, 1) {
		tmp_PDUSessionResourceInformationList := Sequence[*PDUSessionResourceInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofPDUSessions},
			ext: false,
		}
		fn := func() *PDUSessionResourceInformationItem { return new(PDUSessionResourceInformationItem) }
		if err = tmp_PDUSessionResourceInformationList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read PDUSessionResourceInformationList", err)
			return
		}
		ie.PDUSessionResourceInformationList = []PDUSessionResourceInformationItem{}
		for _, i := range tmp_PDUSessionResourceInformationList.Value {
			ie.PDUSessionResourceInformationList = append(ie.PDUSessionResourceInformationList, *i)
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_ERABInformationList := Sequence[*ERABInformationItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofERABs},
			ext: false,
		}
		fn := func() *ERABInformationItem { return new(ERABInformationItem) }
		if err = tmp_ERABInformationList.Decode(r, fn); err != nil {
			err = utils.WrapError("Read ERABInformationList", err)
			return
		}
		ie.ERABInformationList = []ERABInformationItem{}
		for _, i := range tmp_ERABInformationList.Value {
			ie.ERABInformationList = append(ie.ERABInformationList, *i)
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.TargetCellID.Decode(r); err != nil {
			err = utils.WrapError("Read TargetCellID", err)
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		tmp_IndexToRFSP := INTEGER{
			c:   aper.Constraint{Lb: 1, Ub: 256},
			ext: true,
		}
		if err = tmp_IndexToRFSP.Decode(r); err != nil {
			err = utils.WrapError("Read IndexToRFSP", err)
			return
		}
		ie.IndexToRFSP = (*int64)(&tmp_IndexToRFSP.Value)
	}
	if aper.IsBitSet(optionals, 5) {
		tmp_UEHistoryInformation := Sequence[*LastVisitedCellItem]{
			c:   aper.Constraint{Lb: 1, Ub: maxnoofCellsinUEHistoryInfo},
			ext: false,
		}
		fn := func() *LastVisitedCellItem { return new(LastVisitedCellItem) }
		if err = tmp_UEHistoryInformation.Decode(r, fn); err != nil {
			err = utils.WrapError("Read UEHistoryInformation", err)
			return
		}
		ie.UEHistoryInformation = []LastVisitedCellItem{}
		for _, i := range tmp_UEHistoryInformation.Value {
			ie.UEHistoryInformation = append(ie.UEHistoryInformation, *i)
		}
	}
	return
}
