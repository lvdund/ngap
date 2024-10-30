package ies

import "github.com/lvdund/ngap/aper"

type GBRQosInformation struct {
	MaximumFlowBitRateDL    *BitRate             `False,`
	MaximumFlowBitRateUL    *BitRate             `False,`
	GuaranteedFlowBitRateDL *BitRate             `False,`
	GuaranteedFlowBitRateUL *BitRate             `False,`
	NotificationControl     *NotificationControl `False,OPTIONAL`
	MaximumPacketLossRateDL *PacketLossRate      `False,OPTIONAL`
	MaximumPacketLossRateUL *PacketLossRate      `False,OPTIONAL`
	// IEExtensions GBRQosInformationExtIEs `False,OPTIONAL`
}

func (ie *GBRQosInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(aper.One); err != nil {
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
	if ie.MaximumFlowBitRateDL != nil {
		if err = ie.MaximumFlowBitRateDL.Encode(w); err != nil {
			return
		}
	}
	if ie.MaximumFlowBitRateUL != nil {
		if err = ie.MaximumFlowBitRateUL.Encode(w); err != nil {
			return
		}
	}
	if ie.GuaranteedFlowBitRateDL != nil {
		if err = ie.GuaranteedFlowBitRateDL.Encode(w); err != nil {
			return
		}
	}
	if ie.GuaranteedFlowBitRateUL != nil {
		if err = ie.GuaranteedFlowBitRateUL.Encode(w); err != nil {
			return
		}
	}
	if ie.NotificationControl != nil {
		if err = ie.NotificationControl.Encode(w); err != nil {
			return
		}
	}
	if ie.MaximumPacketLossRateDL != nil {
		if err = ie.MaximumPacketLossRateDL.Encode(w); err != nil {
			return
		}
	}
	if ie.MaximumPacketLossRateUL != nil {
		if err = ie.MaximumPacketLossRateUL.Encode(w); err != nil {
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
	ie.MaximumFlowBitRateDL = new(BitRate)
	ie.MaximumFlowBitRateUL = new(BitRate)
	ie.GuaranteedFlowBitRateDL = new(BitRate)
	ie.GuaranteedFlowBitRateUL = new(BitRate)
	ie.NotificationControl = new(NotificationControl)
	ie.MaximumPacketLossRateDL = new(PacketLossRate)
	ie.MaximumPacketLossRateUL = new(PacketLossRate)
	if err = ie.MaximumFlowBitRateDL.Decode(r); err != nil {
		return
	}
	if err = ie.MaximumFlowBitRateUL.Decode(r); err != nil {
		return
	}
	if err = ie.GuaranteedFlowBitRateDL.Decode(r); err != nil {
		return
	}
	if err = ie.GuaranteedFlowBitRateUL.Decode(r); err != nil {
		return
	}
	if aper.IsBitSet(optionals, 2) {
		if err = ie.NotificationControl.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 3) {
		if err = ie.MaximumPacketLossRateDL.Decode(r); err != nil {
			return
		}
	}
	if aper.IsBitSet(optionals, 4) {
		if err = ie.MaximumPacketLossRateUL.Decode(r); err != nil {
			return
		}
	}
	return
}
