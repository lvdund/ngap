package ies

import "github.com/lvdund/ngap/aper"

type FiveGSTMSI struct {
	AMFSetID   []byte
	AMFPointer []byte
	FiveGTMSI  []byte
	// IEExtensions *FiveGSTMSIExtIEs `optional`
}

func (ie *FiveGSTMSI) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_AMFSetID := NewBITSTRING(ie.AMFSetID, aper.Constraint{Lb: 10, Ub: 10}, false)
	if err = tmp_AMFSetID.Encode(w); err != nil {
		return
	}
	tmp_AMFPointer := NewBITSTRING(ie.AMFPointer, aper.Constraint{Lb: 6, Ub: 6}, false)
	if err = tmp_AMFPointer.Encode(w); err != nil {
		return
	}
	tmp_FiveGTMSI := NewOCTETSTRING(ie.FiveGTMSI, aper.Constraint{Lb: 4, Ub: 4}, false)
	if err = tmp_FiveGTMSI.Encode(w); err != nil {
		return
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
	tmp_AMFSetID := BITSTRING{
		c:   aper.Constraint{Lb: 10, Ub: 10},
		ext: false,
	}
	if err = tmp_AMFSetID.Decode(r); err != nil {
		return
	}
	ie.AMFSetID = tmp_AMFSetID.Value.Bytes
	tmp_AMFPointer := BITSTRING{
		c:   aper.Constraint{Lb: 6, Ub: 6},
		ext: false,
	}
	if err = tmp_AMFPointer.Decode(r); err != nil {
		return
	}
	ie.AMFPointer = tmp_AMFPointer.Value.Bytes
	tmp_FiveGTMSI := OCTETSTRING{
		c:   aper.Constraint{Lb: 4, Ub: 4},
		ext: false,
	}
	if err = tmp_FiveGTMSI.Decode(r); err != nil {
		return
	}
	ie.FiveGTMSI = tmp_FiveGTMSI.Value
	return
}
