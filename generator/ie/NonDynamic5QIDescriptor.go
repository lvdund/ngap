package ie

type NonDynamic5QIDescriptor struct {
	FiveQI                 FiveQI                        `False,`
	PriorityLevelQos       PriorityLevelQos              `False,OPTIONAL`
	AveragingWindow        AveragingWindow               `False,OPTIONAL`
	MaximumDataBurstVolume MaximumDataBurstVolume        `False,OPTIONAL`
	IEExtensions           NonDynamic5QIDescriptorExtIEs `False,OPTIONAL`
}
