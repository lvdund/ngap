package ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

type PDUSessionResourceAdmittedItem struct {
	PDUSessionID                       int64  `lb:0,ub:255,madatory`
	HandoverRequestAcknowledgeTransfer []byte `lb:0,ub:0,madatory`
	// IEExtensions *PDUSessionResourceAdmittedItemExtIEs `optional`
}

func (ie *PDUSessionResourceAdmittedItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		err = fmt.Errorf("Encode PDUSessionID: %w", err)
		return
	}
	tmp_HandoverRequestAcknowledgeTransfer := NewOCTETSTRING(ie.HandoverRequestAcknowledgeTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_HandoverRequestAcknowledgeTransfer.Encode(w); err != nil {
		err = fmt.Errorf("Encode HandoverRequestAcknowledgeTransfer: %w", err)
		return
	}
	return
}
func (ie *PDUSessionResourceAdmittedItem) Decode(r *aper.AperReader) (err error) {
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
		err = fmt.Errorf("Read PDUSessionID: %w", err)
		return
	}
	ie.PDUSessionID = int64(tmp_PDUSessionID.Value)
	tmp_HandoverRequestAcknowledgeTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_HandoverRequestAcknowledgeTransfer.Decode(r); err != nil {
		err = fmt.Errorf("Read HandoverRequestAcknowledgeTransfer: %w", err)
		return
	}
	ie.HandoverRequestAcknowledgeTransfer = tmp_HandoverRequestAcknowledgeTransfer.Value
	return
}
