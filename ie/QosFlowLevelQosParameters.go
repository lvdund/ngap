package ie

type QosFlowLevelQosParameters struct {
ChoiceQosCharacteristics	*ChoiceQosCharacteristics
AllocationAndRetentionPriority	*AllocationAndRetentionPriority
GbrQosFlowInformation	*GbrQosFlowInformation
ReflectiveQosAttribute	*[]byte
AdditionalQosFlowInformation	*[]byte
QosMonitoringRequest	*[]byte
QosMonitoringReportingFrequency	*int
}

type ChoiceQosCharacteristics struct {
NonDynamic5Qi	*NonDynamic5Qi
Dynamic5Qi	*Dynamic5Qi
}

type NonDynamic5Qi struct {
NonDynamic5QiDescriptor	*NonDynamic5QiDescriptor
}

type Dynamic5Qi struct {
Dynamic5QiDescriptor	*Dynamic5QiDescriptor
}
