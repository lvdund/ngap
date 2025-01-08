package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceFailedToModifyItemModCfm struct {
	PDUSessionID                                           int64
	PDUSessionResourceModifyIndicationUnsuccessfulTransfer []byte
	// IEExtensions *PDUSessionResourceFailedToModifyItemModCfmExtIEs `optional`
}

func (ie *PDUSessionResourceFailedToModifyItemModCfm) Encode(w *aper.AperWriter) (err error) {
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
	tmp_PDUSessionResourceModifyIndicationUnsuccessfulTransfer := NewOCTETSTRING(ie.PDUSessionResourceModifyIndicationUnsuccessfulTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_PDUSessionResourceModifyIndicationUnsuccessfulTransfer.Encode(w); err != nil {
		err = utils.WrapError("Read PDUSessionResourceModifyIndicationUnsuccessfulTransfer", err)
		return
	}
	return
}
func (ie *PDUSessionResourceFailedToModifyItemModCfm) Decode(r *aper.AperReader) (err error) {
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
	tmp_PDUSessionResourceModifyIndicationUnsuccessfulTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_PDUSessionResourceModifyIndicationUnsuccessfulTransfer.Decode(r); err != nil {
		err = utils.WrapError("Read PDUSessionResourceModifyIndicationUnsuccessfulTransfer", err)
		return
	}
	ie.PDUSessionResourceModifyIndicationUnsuccessfulTransfer = tmp_PDUSessionResourceModifyIndicationUnsuccessfulTransfer.Value
	return
}
