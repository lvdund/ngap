package ies

import "github.com/lvdund/ngap/aper"

type VolumeTimedReportItem struct {
	StartTimeStamp *aper.OctetString `False,`
	EndTimeStamp   *aper.OctetString `False,`
	UsageCountUL   *aper.Integer     `False,`
	UsageCountDL   *aper.Integer     `False,`
	// IEExtensions VolumeTimedReportItemExtIEs `False,OPTIONAL`
}

func (ie *VolumeTimedReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	if ie.StartTimeStamp != nil {
		if err = w.WriteOctetString(*ie.StartTimeStamp, &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return
		}
	}
	if ie.EndTimeStamp != nil {
		if err = w.WriteOctetString(*ie.EndTimeStamp, &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return
		}
	}
	if ie.UsageCountUL != nil {
		if err = w.WriteInteger(int64(*ie.UsageCountUL), &aper.Constraint{Lb: 0, Ub: 9223372036854775807}, false); err != nil {
			return
		}
	}
	if ie.UsageCountDL != nil {
		if err = w.WriteInteger(int64(*ie.UsageCountDL), &aper.Constraint{Lb: 0, Ub: 9223372036854775807}, false); err != nil {
			return
		}
	}
	return
}
func (ie *VolumeTimedReportItem) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	var v int64
	var o []byte
	if o, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return
	} else {
		ie.StartTimeStamp = (*aper.OctetString)(&o)
	}
	if o, err = r.ReadOctetString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return
	} else {
		ie.EndTimeStamp = (*aper.OctetString)(&o)
	}
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9223372036854775807}, false); err != nil {
		return
	} else {
		ie.UsageCountUL = (*aper.Integer)(&v)
	}
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9223372036854775807}, false); err != nil {
		return
	} else {
		ie.UsageCountDL = (*aper.Integer)(&v)
	}
	return
}
