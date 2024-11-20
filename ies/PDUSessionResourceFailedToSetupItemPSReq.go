package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceFailedToSetupItemPSReq struct {
	PDUSessionID                         *PDUSessionID     `False,`
	PathSwitchRequestSetupFailedTransfer *aper.OctetString `False,`
	// IEExtensions PDUSessionResourceFailedToSetupItemPSReqExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceFailedToSetupItemPSReq) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PDUSessionID != nil {
		if err = ie.PDUSessionID.Encode(w); err != nil {
			return
		}
	}
	if ie.PathSwitchRequestSetupFailedTransfer != nil {
		if err = w.WriteOctetString(*ie.PathSwitchRequestSetupFailedTransfer, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return
		}
	}
	return
}
func (ie *PDUSessionResourceFailedToSetupItemPSReq) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PDUSessionID = new(PDUSessionID)
	var o []byte
	if err = ie.PDUSessionID.Decode(r); err != nil {
		return
	}
	if o, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return
	} else {
		ie.PathSwitchRequestSetupFailedTransfer = (*aper.OctetString)(&o)
	}
	return
}
