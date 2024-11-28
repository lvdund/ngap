package ies

import "github.com/lvdund/ngap/aper"

type SourceToTargetTransparentContainer struct {
	Value aper.OctetString `False,`
}

func (ie *SourceToTargetTransparentContainer) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, nil, false)
	return
}
func (ie *SourceToTargetTransparentContainer) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(nil, false); err != nil {
		return
	}
	ie.Value = v
	return
}
