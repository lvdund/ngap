package ie

type MbsQosFlowsToBeSetupList struct {
	MbsQosFlowsSetupRequestItem MbsQosFlowsSetupRequestItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsQosFlowsSetupRequestItem struct {
	MbsQosFlowIdentifier         QosFlowIdentifier         //`bitstring:"sizeLB:0,sizeUB:150"`
	MbsQosFlowLevelQosParameters QosFlowLevelQosParameters //`bitstring:"sizeLB:0,sizeUB:150"`
}
