package ies

import "github.com/lvdund/ngap/aper"

type FiveGSTMSI struct {
	AMFSetID   *AMFSetID   `False,`
	AMFPointer *AMFPointer `False,`
	FiveGTMSI  *FiveGTMSI  `False,`
	// IEExtensions FiveGSTMSIExtIEs `False,OPTIONAL`
}

func (ie *FiveGSTMSI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.AMFSetID != nil {
		if err = ie.AMFSetID.Encode(w); err != nil {
			return
		}
	}
	if ie.AMFPointer != nil {
		if err = ie.AMFPointer.Encode(w); err != nil {
			return
		}
	}
	if ie.FiveGTMSI != nil {
		if err = ie.FiveGTMSI.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *FiveGSTMSI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.AMFSetID = new(AMFSetID)
	ie.AMFPointer = new(AMFPointer)
	ie.FiveGTMSI = new(FiveGTMSI)
	if err = ie.AMFSetID.Decode(r); err != nil {
		return
	}
	if err = ie.AMFPointer.Decode(r); err != nil {
		return
	}
	if err = ie.FiveGTMSI.Decode(r); err != nil {
		return
	}
	return
}
