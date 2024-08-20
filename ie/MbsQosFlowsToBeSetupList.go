package ie

type MbsQosFlowsToBeSetupList struct {
MbsQosFlowsSetupRequestItem	*MbsQosFlowsSetupRequestItem
}

type MbsQosFlowsSetupRequestItem struct {
MbsQosFlowIdentifier	*QosFlowIdentifier
MbsQosFlowLevelQosParameters	*QosFlowLevelQosParameters
}
