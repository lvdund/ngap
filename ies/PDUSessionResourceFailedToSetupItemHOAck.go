package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceFailedToSetupItemHOAck struct {
	PDUSessionID                                   int64
	HandoverResourceAllocationUnsuccessfulTransfer []byte
	// IEExtensions *PDUSessionResourceFailedToSetupItemHOAckExtIEs `optional`
}

func (ie *PDUSessionResourceFailedToSetupItemHOAck) Encode(w *aper.AperWriter) (err error) {
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
	tmp_HandoverResourceAllocationUnsuccessfulTransfer := NewOCTETSTRING(ie.HandoverResourceAllocationUnsuccessfulTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_HandoverResourceAllocationUnsuccessfulTransfer.Encode(w); err != nil {
		err = utils.WrapError("Read HandoverResourceAllocationUnsuccessfulTransfer", err)
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToSetupItemHOAck) Decode(r *aper.AperReader) (err error) {
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
	tmp_HandoverResourceAllocationUnsuccessfulTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_HandoverResourceAllocationUnsuccessfulTransfer.Decode(r); err != nil {
		err = utils.WrapError("Read HandoverResourceAllocationUnsuccessfulTransfer", err)
		return
	}
	ie.HandoverResourceAllocationUnsuccessfulTransfer = tmp_HandoverResourceAllocationUnsuccessfulTransfer.Value
	return
}
