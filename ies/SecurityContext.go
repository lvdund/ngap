package ies

import (
	"github.com/lvdund/ngap/aper"
	"github.com/reogac/utils"
)

type SecurityContext struct {
	NextHopChainingCount int64
	NextHopNH            []byte
	// IEExtensions *SecurityContextExtIEs `optional`
}

func (ie *SecurityContext) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	w.WriteBits(optionals, 1)
	tmp_NextHopChainingCount := NewINTEGER(ie.NextHopChainingCount, aper.Constraint{Lb: 0, Ub: 7}, false)
	if err = tmp_NextHopChainingCount.Encode(w); err != nil {
		err = utils.WrapError("Read NextHopChainingCount", err)
		return
	}
	tmp_NextHopNH := NewBITSTRING(ie.NextHopNH, aper.Constraint{Lb: 256, Ub: 256}, false)
	if err = tmp_NextHopNH.Encode(w); err != nil {
		err = utils.WrapError("Read NextHopNH", err)
		return
	}
	return
}
func (ie *SecurityContext) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	if _, err = r.ReadBits(1); err != nil {
		return
	}
	tmp_NextHopChainingCount := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 7},
		ext: false,
	}
	if err = tmp_NextHopChainingCount.Decode(r); err != nil {
		err = utils.WrapError("Read NextHopChainingCount", err)
		return
	}
	ie.NextHopChainingCount = int64(tmp_NextHopChainingCount.Value)
	tmp_NextHopNH := BITSTRING{
		c:   aper.Constraint{Lb: 256, Ub: 256},
		ext: false,
	}
	if err = tmp_NextHopNH.Decode(r); err != nil {
		err = utils.WrapError("Read NextHopNH", err)
		return
	}
	ie.NextHopNH = tmp_NextHopNH.Value.Bytes
	return
}
