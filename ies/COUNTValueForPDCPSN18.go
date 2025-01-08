package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type COUNTValueForPDCPSN18 struct {
	PDCPSN18    int64
	HFNPDCPSN18 int64
	// IEExtensions *COUNTValueForPDCPSN18ExtIEs `optional`
}

func (ie *COUNTValueForPDCPSN18) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDCPSN18 := NewINTEGER(ie.PDCPSN18, aper.Constraint{Lb: 0, Ub: 262143}, false)
	if err = tmp_PDCPSN18.Encode(w); err != nil {
		err = utils.WrapError("Read PDCPSN18", err)
		return
	}
	tmp_HFNPDCPSN18 := NewINTEGER(ie.HFNPDCPSN18, aper.Constraint{Lb: 0, Ub: 16383}, false)
	if err = tmp_HFNPDCPSN18.Encode(w); err != nil {
		err = utils.WrapError("Read HFNPDCPSN18", err)
		return
	}
	return
}
func (ie *COUNTValueForPDCPSN18) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PDCPSN18 := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 262143},
		ext: false,
	}
	if err = tmp_PDCPSN18.Decode(r); err != nil {
		err = utils.WrapError("Read PDCPSN18", err)
		return
	}
	ie.PDCPSN18 = int64(tmp_PDCPSN18.Value)
	tmp_HFNPDCPSN18 := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 16383},
		ext: false,
	}
	if err = tmp_HFNPDCPSN18.Decode(r); err != nil {
		err = utils.WrapError("Read HFNPDCPSN18", err)
		return
	}
	ie.HFNPDCPSN18 = int64(tmp_HFNPDCPSN18.Value)
	return
}
