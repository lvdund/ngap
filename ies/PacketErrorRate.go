package ies

import "github.com/lvdund/ngap/aper"

type PacketErrorRate struct {
	PERScalar   int64
	PERExponent int64
	// IEExtensions  *PacketErrorRateExtIEs
}

func (ie *PacketErrorRate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PERScalar := NewINTEGER(ie.PERScalar, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_PERScalar.Encode(w); err != nil {
		return
	}
	tmp_PERExponent := NewINTEGER(ie.PERExponent, aper.Constraint{Lb: 0, Ub: 0}, false)
	if err = tmp_PERExponent.Encode(w); err != nil {
		return
	}
	return
}
func (ie *PacketErrorRate) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_PERScalar := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_PERScalar.Decode(r); err != nil {
		return
	}
	ie.PERScalar = int64(tmp_PERScalar.Value)
	tmp_PERExponent := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 0},
		ext: false,
	}
	if err = tmp_PERExponent.Decode(r); err != nil {
		return
	}
	ie.PERExponent = int64(tmp_PERExponent.Value)
	return
}
