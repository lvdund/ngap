package ie

type GBRQosInformation struct {
	MaximumFlowBitRateDL    BitRate                 `False,`
	MaximumFlowBitRateUL    BitRate                 `False,`
	GuaranteedFlowBitRateDL BitRate                 `False,`
	GuaranteedFlowBitRateUL BitRate                 `False,`
	NotificationControl     NotificationControl     `False,OPTIONAL`
	MaximumPacketLossRateDL PacketLossRate          `False,OPTIONAL`
	MaximumPacketLossRateUL PacketLossRate          `False,OPTIONAL`
	IEExtensions            GBRQosInformationExtIEs `False,OPTIONAL`
}
