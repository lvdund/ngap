package ies

import "github.com/lvdund/ngap/aper"

type LastVisitedEUTRANCellInformation struct {
	Value aper.OctetString `False,`
}

func (ie *LastVisitedEUTRANCellInformation) Encode(w *aper.AperWriter) (err error) {
	err = w.WriteOctetString(ie.Value, nil, false)
	return
}
func (ie *LastVisitedEUTRANCellInformation) Decode(r *aper.AperReader) (err error) {
	var v aper.OctetString
	if v, err = r.ReadOctetString(nil, false); err != nil {
		return
	}
	ie.Value = v
	return
}
