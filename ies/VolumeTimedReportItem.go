package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type VolumeTimedReportItem struct {
	StartTimeStamp []byte `lb:4,ub:4,madatory`
	EndTimeStamp   []byte `lb:4,ub:4,madatory`
	UsageCountUL   int64  `lb:0,ub:1844674407370955161,madatory`
	UsageCountDL   int64  `lb:0,ub:1844674407370955161,madatory`
	// IEExtensions *VolumeTimedReportItemExtIEs `optional`
}

func (ie *VolumeTimedReportItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_StartTimeStamp := NewOCTETSTRING(ie.StartTimeStamp, aper.Constraint{Lb: 4, Ub: 4}, false)
	if err = tmp_StartTimeStamp.Encode(w); err != nil {
		err = utils.WrapError("Encode StartTimeStamp", err)
		return
	}
	tmp_EndTimeStamp := NewOCTETSTRING(ie.EndTimeStamp, aper.Constraint{Lb: 4, Ub: 4}, false)
	if err = tmp_EndTimeStamp.Encode(w); err != nil {
		err = utils.WrapError("Encode EndTimeStamp", err)
		return
	}
	tmp_UsageCountUL := NewINTEGER(ie.UsageCountUL, aper.Constraint{Lb: 0, Ub: 1844674407370955161}, false)
	if err = tmp_UsageCountUL.Encode(w); err != nil {
		err = utils.WrapError("Encode UsageCountUL", err)
		return
	}
	tmp_UsageCountDL := NewINTEGER(ie.UsageCountDL, aper.Constraint{Lb: 0, Ub: 1844674407370955161}, false)
	if err = tmp_UsageCountDL.Encode(w); err != nil {
		err = utils.WrapError("Encode UsageCountDL", err)
		return
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
	tmp_StartTimeStamp := OCTETSTRING{
		c:   aper.Constraint{Lb: 4, Ub: 4},
		ext: false,
	}
	if err = tmp_StartTimeStamp.Decode(r); err != nil {
		err = utils.WrapError("Read StartTimeStamp", err)
		return
	}
	ie.StartTimeStamp = tmp_StartTimeStamp.Value
	tmp_EndTimeStamp := OCTETSTRING{
		c:   aper.Constraint{Lb: 4, Ub: 4},
		ext: false,
	}
	if err = tmp_EndTimeStamp.Decode(r); err != nil {
		err = utils.WrapError("Read EndTimeStamp", err)
		return
	}
	ie.EndTimeStamp = tmp_EndTimeStamp.Value
	tmp_UsageCountUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 1844674407370955161},
		ext: false,
	}
	if err = tmp_UsageCountUL.Decode(r); err != nil {
		err = utils.WrapError("Read UsageCountUL", err)
		return
	}
	ie.UsageCountUL = int64(tmp_UsageCountUL.Value)
	tmp_UsageCountDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 1844674407370955161},
		ext: false,
	}
	if err = tmp_UsageCountDL.Decode(r); err != nil {
		err = utils.WrapError("Read UsageCountDL", err)
		return
	}
	ie.UsageCountDL = int64(tmp_UsageCountDL.Value)
	return
}
