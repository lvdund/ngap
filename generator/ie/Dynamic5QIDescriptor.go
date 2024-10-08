package ie

type Dynamic5QIDescriptor struct {
	PriorityLevelQos       PriorityLevelQos           `False,`
	PacketDelayBudget      PacketDelayBudget          `False,`
	PacketErrorRate        PacketErrorRate            `True,`
	FiveQI                 FiveQI                     `False,OPTIONAL`
	DelayCritical          DelayCritical              `False,OPTIONAL`
	AveragingWindow        AveragingWindow            `False,OPTIONAL`
	MaximumDataBurstVolume MaximumDataBurstVolume     `False,OPTIONAL`
	IEExtensions           Dynamic5QIDescriptorExtIEs `False,OPTIONAL`
}
