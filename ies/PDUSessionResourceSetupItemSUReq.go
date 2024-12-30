package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceSetupItemSUReq struct {
	PDUSessionID                           int64
	PDUSessionNASPDU                       *[]byte
	SNSSAI                                 *SNSSAI
	PDUSessionResourceSetupRequestTransfer *[]byte
	// IEExtensions  *PDUSessionResourceSetupItemSUReqExtIEs
}

func (ie *PDUSessionResourceSetupItemSUReq) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.PDUSessionNASPDU != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.SNSSAI != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.PDUSessionResourceSetupRequestTransfer != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	tmp_PDUSessionID := NewINTEGER(ie.PDUSessionID, aper.Constraint{Lb: 0, Ub: 255}, false)
	if err = tmp_PDUSessionID.Encode(w); err != nil {
		return
	}
	if ie.PDUSessionNASPDU != nil {
		tmp_PDUSessionNASPDU := NewOCTETSTRING(*ie.PDUSessionNASPDU, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_PDUSessionNASPDU.Encode(w); err != nil {
			return
		}
	}
	if ie.SNSSAI != nil {
		if err = ie.SNSSAI.Encode(w); err != nil {
			return
		}
	}
	if ie.PDUSessionResourceSetupRequestTransfer != nil {
		tmp_PDUSessionResourceSetupRequestTransfer := NewOCTETSTRING(*ie.PDUSessionResourceSetupRequestTransfer, aper.Constraint{Lb: 0, Ub: 0}, false)
		if err = tmp_PDUSessionResourceSetupRequestTransfer.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceSetupItemSUReq) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	tmp_PDUSessionID := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 255},
		ext: false,
	}
	if err = tmp_PDUSessionID.Decode(r); err != nil {
		return
	}
	ie.PDUSessionID = int64(tmp_PDUSessionID.Value)
	if aper.IsBitSet(optionals, 1) {
		tmp_PDUSessionNASPDU := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_PDUSessionNASPDU.Decode(r); err != nil {
			return
		}
		*ie.PDUSessionNASPDU = tmp_PDUSessionNASPDU.Value
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.SNSSAI.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_PDUSessionResourceSetupRequestTransfer := OCTETSTRING{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		if err = tmp_PDUSessionResourceSetupRequestTransfer.Decode(r); err != nil {
			return
		}
		*ie.PDUSessionResourceSetupRequestTransfer = tmp_PDUSessionResourceSetupRequestTransfer.Value
	}
	return
}
