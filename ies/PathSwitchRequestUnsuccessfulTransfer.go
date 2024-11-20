package ies

import "github.com/lvdund/ngap/aper"

type PathSwitchRequestUnsuccessfulTransfer struct {
	Cause *Cause `False,`
	// IEExtensions PathSwitchRequestUnsuccessfulTransferExtIEs `False,OPTIONAL`
}

func (ie *PathSwitchRequestUnsuccessfulTransfer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.Cause != nil {
		if err = ie.Cause.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *PathSwitchRequestUnsuccessfulTransfer) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.Cause = new(Cause)
	if err = ie.Cause.Decode(r); err != nil {
		return
	}
	return
}
