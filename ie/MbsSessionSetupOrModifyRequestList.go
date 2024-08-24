package ie

type MbsSessionSetupOrModifyRequestList struct {
MbsSessionSetupOrModifyRequestList	[]MbsSessionSetupOrModifyRequestItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsSessionSetupOrModifyRequestItem struct {
MbsSessionId	MbsSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsAreaSessionId	MbsAreaSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
AssociatedMbsQosFlowSetupOrModifyRequestList	[]AssociatedMbsQosFlowSetupOrModifyRequestItem	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsQosFlowToReleaseList	QosFlowListWithCause	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type AssociatedMbsQosFlowSetupOrModifyRequestItem struct {
MbsQosFlowIdentifier	QosFlowIdentifier	//`bitstring:"sizeLB:0,sizeUB:150"`
AssociatedUnicastQosFlowIdentifier	QosFlowIdentifier	//`bitstring:"sizeLB:0,sizeUB:150"`
}
