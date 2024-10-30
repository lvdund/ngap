package ies

import "github.com/lvdund/ngap/aper"

type PDUSessionResourceReleaseResponseTransfer struct {
	// IEExtensions PDUSessionResourceReleaseResponseTransferExtIEs `False,OPTIONAL`
}

func (ie *PDUSessionResourceReleaseResponseTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	return
}
func (ie *PDUSessionResourceReleaseResponseTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	return
}
