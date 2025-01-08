package ies

import "github.com/lvdund/ngap/aper"

type GBRQosInformation struct {
	MaximumFlowBitRateDL    int64
	MaximumFlowBitRateUL    int64
	GuaranteedFlowBitRateDL int64
	GuaranteedFlowBitRateUL int64
	NotificationControl     *NotificationControl `optional`
	MaximumPacketLossRateDL *int64               `optional`
	MaximumPacketLossRateUL *int64               `optional`
	// IEExtensions *GBRQosInformationExtIEs `optional`
}

func (ie *GBRQosInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.Zero); err != nil {
		return
	}
	optionals := []byte{0x0}
	if ie.NotificationControl != nil {
		aper.SetBit(optionals, 1)
	}
	if ie.MaximumPacketLossRateDL != nil {
		aper.SetBit(optionals, 2)
	}
	if ie.MaximumPacketLossRateUL != nil {
		aper.SetBit(optionals, 3)
	}
	w.WriteBits(optionals, 4)
	tmp_MaximumFlowBitRateDL := NewINTEGER(ie.MaximumFlowBitRateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, false)
	if err = tmp_MaximumFlowBitRateDL.Encode(w); err != nil {
		return
	}
	tmp_MaximumFlowBitRateUL := NewINTEGER(ie.MaximumFlowBitRateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, false)
	if err = tmp_MaximumFlowBitRateUL.Encode(w); err != nil {
		return
	}
	tmp_GuaranteedFlowBitRateDL := NewINTEGER(ie.GuaranteedFlowBitRateDL, aper.Constraint{Lb: 0, Ub: 4000000000000}, false)
	if err = tmp_GuaranteedFlowBitRateDL.Encode(w); err != nil {
		return
	}
	tmp_GuaranteedFlowBitRateUL := NewINTEGER(ie.GuaranteedFlowBitRateUL, aper.Constraint{Lb: 0, Ub: 4000000000000}, false)
	if err = tmp_GuaranteedFlowBitRateUL.Encode(w); err != nil {
		return
	}
	if ie.NotificationControl != nil {
		if err = ie.NotificationControl.Encode(w); err != nil {
			return
		}
	}
	if ie.MaximumPacketLossRateDL != nil {
		tmp_MaximumPacketLossRateDL := NewINTEGER(*ie.MaximumPacketLossRateDL, aper.Constraint{Lb: 0, Ub: 1000}, false)
		if err = tmp_MaximumPacketLossRateDL.Encode(w); err != nil {
			return
		}
	}
	if ie.MaximumPacketLossRateUL != nil {
		tmp_MaximumPacketLossRateUL := NewINTEGER(*ie.MaximumPacketLossRateUL, aper.Constraint{Lb: 0, Ub: 1000}, false)
		if err = tmp_MaximumPacketLossRateUL.Encode(w); err != nil {
			return
		}
	}
	return
}
func (ie *GBRQosInformation) Decode(r *aper.AperReader) (err error) {
	if _, err = r.ReadBool(); err != nil {
		return
	}
	var optionals []byte
	if optionals, err = r.ReadBits(4); err != nil {
		return
	}
	tmp_MaximumFlowBitRateDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: false,
	}
	if err = tmp_MaximumFlowBitRateDL.Decode(r); err != nil {
		return
	}
	ie.MaximumFlowBitRateDL = int64(tmp_MaximumFlowBitRateDL.Value)
	tmp_MaximumFlowBitRateUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: false,
	}
	if err = tmp_MaximumFlowBitRateUL.Decode(r); err != nil {
		return
	}
	ie.MaximumFlowBitRateUL = int64(tmp_MaximumFlowBitRateUL.Value)
	tmp_GuaranteedFlowBitRateDL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: false,
	}
	if err = tmp_GuaranteedFlowBitRateDL.Decode(r); err != nil {
		return
	}
	ie.GuaranteedFlowBitRateDL = int64(tmp_GuaranteedFlowBitRateDL.Value)
	tmp_GuaranteedFlowBitRateUL := INTEGER{
		c:   aper.Constraint{Lb: 0, Ub: 4000000000000},
		ext: false,
	}
	if err = tmp_GuaranteedFlowBitRateUL.Decode(r); err != nil {
		return
	}
	ie.GuaranteedFlowBitRateUL = int64(tmp_GuaranteedFlowBitRateUL.Value)
	if aper.IsBitSet(optionals, 1) {
		if err = ie.NotificationControl.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 2) {
		tmp_MaximumPacketLossRateDL := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1000},
			ext: false,
		}
		if err = tmp_MaximumPacketLossRateDL.Decode(r); err != nil {
			return
		}
		ie.MaximumPacketLossRateDL = (*int64)(&tmp_MaximumPacketLossRateDL.Value)
	}
	if aper.IsBitSet(optionals, 3) {
		tmp_MaximumPacketLossRateUL := INTEGER{
			c:   aper.Constraint{Lb: 0, Ub: 1000},
			ext: false,
		}
		if err = tmp_MaximumPacketLossRateUL.Decode(r); err != nil {
			return
		}
		ie.MaximumPacketLossRateUL = (*int64)(&tmp_MaximumPacketLossRateUL.Value)
	}
	return
}
