package ies

import "github.com/lvdund/ngap/aper"

type DRBsSubjectToStatusTransferItem struct {
	DRBID       int64
	DRBStatusUL DRBStatusUL
	DRBStatusDL DRBStatusDL
	// IEExtension *DRBsSubjectToStatusTransferItemExtIEs `optional`
}

func (ie *DRBsSubjectToStatusTransferItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_DRBID := NewINTEGER(ie.DRBID, aper.Constraint{Lb: 1, Ub: 32}, false)
	if err = tmp_DRBID.Encode(w); err != nil {
		return
	}
	if err = ie.DRBStatusUL.Encode(w); err != nil {
		return
	}
	if err = ie.DRBStatusDL.Encode(w); err != nil {
		return
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
	tmp_DRBID := INTEGER{
		c:   aper.Constraint{Lb: 1, Ub: 32},
		ext: false,
	}
	if err = tmp_DRBID.Decode(r); err != nil {
		return
	}
	ie.DRBID = int64(tmp_DRBID.Value)
	if err = ie.DRBStatusUL.Decode(r); err != nil {
		return
	}
	if err = ie.DRBStatusDL.Decode(r); err != nil {
		return
	}
	return
}
