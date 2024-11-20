package ies

import "github.com/lvdund/ngap/aper"

type DRBsSubjectToStatusTransferItem struct {
	DRBID       *DRBID       `False,`
	DRBStatusUL *DRBStatusUL `False,`
	DRBStatusDL *DRBStatusDL `False,`
	// IEExtension DRBsSubjectToStatusTransferItemExtIEs `False,OPTIONAL`
}

func (ie *DRBsSubjectToStatusTransferItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.DRBID != nil {
		if err = ie.DRBID.Encode(w); err != nil {
			return
		}
	}
	if ie.DRBStatusUL != nil {
		if err = ie.DRBStatusUL.Encode(w); err != nil {
			return
		}
	}
	if ie.DRBStatusDL != nil {
		if err = ie.DRBStatusDL.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *DRBsSubjectToStatusTransferItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.DRBID = new(DRBID)
	ie.DRBStatusUL = new(DRBStatusUL)
	ie.DRBStatusDL = new(DRBStatusDL)
	if err = ie.DRBID.Decode(r); err != nil {
		return
	}
	if err = ie.DRBStatusUL.Decode(r); err != nil {
		return
	}
	if err = ie.DRBStatusDL.Decode(r); err != nil {
		return
	}
	return
}
