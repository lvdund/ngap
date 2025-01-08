package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceSecondaryRATUsageItem struct {
	PDUSessionID                        int64
	SecondaryRATDataUsageReportTransfer []byte
	// IEExtensions *PDUSessionResourceSecondaryRATUsageItemExtIEs `optional`
}

func (ie *PDUSessionResourceSecondaryRATUsageItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		err = utils.WrapError("Read PDUSessionID", err)
		return
	}
	tmp_SecondaryRATDataUsageReportTransfer := NewOCTETSTRING(ie.SecondaryRATDataUsageReportTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_SecondaryRATDataUsageReportTransfer.Encode(w); err != nil {
		err = utils.WrapError("Read SecondaryRATDataUsageReportTransfer", err)
		return
	}
	return
}
func (ie *PDUSessionResourceSecondaryRATUsageItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PDUSessionID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_PDUSessionID.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionID", err)
		return
	}
	ie.PDUSessionID = int64(tmp_PDUSessionID.Value)
	tmp_SecondaryRATDataUsageReportTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_SecondaryRATDataUsageReportTransfer.Decode(r); err != nil {
		err = utils.WrapError("Read SecondaryRATDataUsageReportTransfer", err)
		return
	}
	ie.SecondaryRATDataUsageReportTransfer = tmp_SecondaryRATDataUsageReportTransfer.Value
	return
}
