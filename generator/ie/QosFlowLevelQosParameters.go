package ie

type QosFlowLevelQosParameters struct {
	QosCharacteristics             QosCharacteristics              `False,`
	AllocationAndRetentionPriority AllocationAndRetentionPriority  `True,`
	GBRQosInformation              GBRQosInformation               `True,OPTIONAL`
	ReflectiveQosAttribute         ReflectiveQosAttribute          `False,OPTIONAL`
	AdditionalQosFlowInformation   AdditionalQosFlowInformation    `False,OPTIONAL`
	IEExtensions                   QosFlowLevelQosParametersExtIEs `False,OPTIONAL`
}
