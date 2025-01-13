package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PDUSessionResourceSetupItemHOReq struct {
	PDUSessionID            int64  `lb:0,ub:255,madatory`
	SNSSAI                  SNSSAI `madatory`
	HandoverRequestTransfer []byte `lb:0,ub:0,madatory`
	// IEExtensions *PDUSessionResourceSetupItemHOReqExtIEs `optional`
}

func (ie *PDUSessionResourceSetupItemHOReq) Encode(w *aper.AperWriter) (err error) {
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
	if err = ie.SNSSAI.Encode(w); err != nil {
		err = utils.WrapError("Encode SNSSAI", err)
		return
	}
	tmp_HandoverRequestTransfer := NewOCTETSTRING(ie.HandoverRequestTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_HandoverRequestTransfer.Encode(w); err != nil {
		err = utils.WrapError("Encode HandoverRequestTransfer", err)
		return
	}
	return
}
func (ie *PDUSessionResourceSetupItemHOReq) Decode(r *aper.AperReader) (err error) {
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
	if err = ie.SNSSAI.Decode(r); err != nil {
		err = utils.WrapError("Read SNSSAI", err)
		return
	}
	tmp_HandoverRequestTransfer := OCTETSTRING{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_HandoverRequestTransfer.Decode(r); err != nil {
		err = utils.WrapError("Read HandoverRequestTransfer", err)
		return
	}
	ie.HandoverRequestTransfer = tmp_HandoverRequestTransfer.Value
	return
}
