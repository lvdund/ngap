package ies

import "github.com/lvdund/ngap/aper"

type GUAMI struct {
	PLMNIdentity *PLMNIdentity `False,`
	AMFRegionID  *AMFRegionID  `False,`
	AMFSetID     *AMFSetID     `False,`
	AMFPointer   *AMFPointer   `False,`
	// IEExtensions GUAMIExtIEs `False,OPTIONAL`
}

func (ie *GUAMI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PLMNIdentity != nil {
		if err = ie.PLMNIdentity.Encode(w); err != nil {
			return
		}
	}
	if ie.AMFRegionID != nil {
		if err = ie.AMFRegionID.Encode(w); err != nil {
			return
		}
	}
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
	return
}
func (ie *GUAMI) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.PLMNIdentity = new(PLMNIdentity)
	ie.AMFRegionID = new(AMFRegionID)
	ie.AMFSetID = new(AMFSetID)
	ie.AMFPointer = new(AMFPointer)
	if err = ie.PLMNIdentity.Decode(r); err != nil {
		return
	}
	if err = ie.AMFRegionID.Decode(r); err != nil {
		return
	}
	if err = ie.AMFSetID.Decode(r); err != nil {
		return
	}
	if err = ie.AMFPointer.Decode(r); err != nil {
		return
	}
	return
}
