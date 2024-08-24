package ie

import "ngap/aper"

type QosFlowLevelQosParameters struct {
ChoiceQosCharacteristics	ChoiceQosCharacteristics	//`bitstring:"sizeLB:0,sizeUB:150"`
AllocationAndRetentionPriority	AllocationAndRetentionPriority	//`bitstring:"sizeLB:0,sizeUB:150"`
GbrQosFlowInformation	GbrQosFlowInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
ReflectiveQosAttribute	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
AdditionalQosFlowInformation	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
QosMonitoringRequest	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
QosMonitoringReportingFrequency	aper.Integer	//`Integer:"valueLB:0,valueUB:150"`
}

type ChoiceQosCharacteristics struct {
NonDynamic5Qi	NonDynamic5Qi	//`bitstring:"sizeLB:0,sizeUB:150"`
Dynamic5Qi	Dynamic5Qi	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type NonDynamic5Qi struct {
NonDynamic5QiDescriptor	NonDynamic5QiDescriptor	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type Dynamic5Qi struct {
Dynamic5QiDescriptor	Dynamic5QiDescriptor	//`bitstring:"sizeLB:0,sizeUB:150"`
}
