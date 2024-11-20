package ies

import "github.com/lvdund/ngap/aper"

type PacketErrorRate struct {
	PERExponent *aper.Integer `False,`
	// IEExtensions PacketErrorRateExtIEs `False,OPTIONAL`
}

func (ie *PacketErrorRate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.PERExponent != nil {
		if err = w.WriteInteger(int64(*ie.PERExponent), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return
		}
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
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return
	} else {
		ie.PERExponent = (*aper.Integer)(&v)
	}
	return
}
