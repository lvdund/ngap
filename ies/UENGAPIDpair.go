package ies

import "github.com/lvdund/ngap/aper"

type UENGAPIDpair struct {
	AMFUENGAPID *AMFUENGAPID `False,`
	RANUENGAPID *RANUENGAPID `False,`
	// IEExtensions UENGAPIDpairExtIEs `False,OPTIONAL`
}

func (ie *UENGAPIDpair) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.AMFUENGAPID != nil {
		if err = ie.AMFUENGAPID.Encode(w); err != nil {
			return
		}
	}
	if ie.RANUENGAPID != nil {
		if err = ie.RANUENGAPID.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *UENGAPIDpair) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	ie.AMFUENGAPID = new(AMFUENGAPID)
	ie.RANUENGAPID = new(RANUENGAPID)
	if err = ie.AMFUENGAPID.Decode(r); err != nil {
		return
	}
	if err = ie.RANUENGAPID.Decode(r); err != nil {
		return
	}
	return
}
