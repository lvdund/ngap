package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceToReleaseItemHOCmd struct {
	PDUSessionID                            int64  `lb:0,ub:255,madatory`
	HandoverPreparationUnsuccessfulTransfer []byte `lb:0,ub:0,madatory`
	// IEExtensions *PDUSessionResourceToReleaseItemHOCmdExtIEs `optional`
}

func (ie *PDUSessionResourceToReleaseItemHOCmd) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		err = utils.WrapError("Encode PDUSessionID", err)
		return
	}
	tmp_HandoverPreparationUnsuccessfulTransfer := NewOCTETSTRING(ie.HandoverPreparationUnsuccessfulTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_HandoverPreparationUnsuccessfulTransfer.Encode(w); err != nil {
		err = utils.WrapError("Encode HandoverPreparationUnsuccessfulTransfer", err)
		return
	}
	return
}
func (ie *PDUSessionResourceToReleaseItemHOCmd) Decode(r *aper.AperReader) (err error) {
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
	tmp_HandoverPreparationUnsuccessfulTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_HandoverPreparationUnsuccessfulTransfer.Decode(r); err != nil {
		err = utils.WrapError("Read HandoverPreparationUnsuccessfulTransfer", err)
		return
	}
	ie.HandoverPreparationUnsuccessfulTransfer = tmp_HandoverPreparationUnsuccessfulTransfer.Value
	return
}
