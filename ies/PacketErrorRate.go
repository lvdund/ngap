package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type PacketErrorRate struct {
	PERScalar   int64 `lb:0,ub:9,madatory,valExt`
	PERExponent int64 `lb:0,ub:9,madatory,valExt`
	// IEExtensions *PacketErrorRateExtIEs `optional,valExt`
}

func (ie *PacketErrorRate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_PERScalar := NewINTEGER(ie.PERScalar, aper.Constraint{Lb: 0, Ub: 9}, true)
	if err = tmp_PERScalar.Encode(w); err != nil {
		err = utils.WrapError("Encode PERScalar", err)
		return
	}
	tmp_PERExponent := NewINTEGER(ie.PERExponent, aper.Constraint{Lb: 0, Ub: 9}, true)
	if err = tmp_PERExponent.Encode(w); err != nil {
		err = utils.WrapError("Encode PERExponent", err)
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
		c:   aper.Constraint{Lb: 0, Ub: 9},
		ext: true,
	}
	if err = tmp_PERScalar.Decode(r); err != nil {
		err = utils.WrapError("Read PERScalar", err)
		return
	}
	ie.PERScalar = int64(tmp_PERScalar.Value)
	tmp_PERExponent := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 9},
		ext: true,
	}
	if err = tmp_PERExponent.Decode(r); err != nil {
		err = utils.WrapError("Read PERExponent", err)
		return
	}
	ie.PERExponent = int64(tmp_PERExponent.Value)
	return
}
