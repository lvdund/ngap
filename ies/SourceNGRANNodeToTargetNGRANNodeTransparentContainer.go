package ies

import "github.com/lvdund/ngap/aper"

type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                      *RRCContainer                      `False,`
	PDUSessionResourceInformationList *PDUSessionResourceInformationList `False,OPTIONAL`
	ERABInformationList               *ERABInformationList               `False,OPTIONAL`
	TargetCellID                      *NGRANCGI                          `False,`
	IndexToRFSP                       *IndexToRFSP                       `False,OPTIONAL`
	UEHistoryInformation              *UEHistoryInformation              `False,`
	// IEExtensions SourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs `False,OPTIONAL`
}

func (ie *SourceNGRANNodeToTargetNGRANNodeTransparentContainer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PDUSessionResourceInformationList != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.ERABInformationList != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.IndexToRFSP != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	if ie.RRCContainer != nil {
		if err = ie.RRCContainer.Encode(w); err != nil {
			return
		}
	}
	if ie.PDUSessionResourceInformationList != nil {
		if err = ie.PDUSessionResourceInformationList.Encode(w); err != nil {
			return
		}
	}
	if ie.ERABInformationList != nil {
		if err = ie.ERABInformationList.Encode(w); err != nil {
			return
		}
	}
	if ie.TargetCellID != nil {
		if err = ie.TargetCellID.Encode(w); err != nil {
			return
		}
	}
	if ie.IndexToRFSP != nil {
		if err = ie.IndexToRFSP.Encode(w); err != nil {
			return
		}
	}
	if ie.UEHistoryInformation != nil {
		if err = ie.UEHistoryInformation.Encode(w); err != nil {
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
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	ie.RRCContainer = new(RRCContainer)
	ie.PDUSessionResourceInformationList = new(PDUSessionResourceInformationList)
	ie.ERABInformationList = new(ERABInformationList)
	ie.TargetCellID = new(NGRANCGI)
	ie.IndexToRFSP = new(IndexToRFSP)
	ie.UEHistoryInformation = new(UEHistoryInformation)
	if err = ie.RRCContainer.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.PDUSessionResourceInformationList.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.ERABInformationList.Decode(r); err != nil {
			return
		}
	}
	if err = ie.TargetCellID.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.IndexToRFSP.Decode(r); err != nil {
			return
		}
	}
	if err = ie.UEHistoryInformation.Decode(r); err != nil {
		return
	}
	return
}
