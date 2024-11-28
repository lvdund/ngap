package ies

import "github.com/lvdund/ngap/aper"

type UERadioCapabilityForPagingOfNR struct {
	Value aper.OctetString `False,0,0`
}

func (ie *UERadioCapabilityForPagingOfNR) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, nil, false)
	return
}
func (ie *UERadioCapabilityForPagingOfNR) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(nil, false); err != nil {
		return
	}
	ie.Value = v
	return
}
