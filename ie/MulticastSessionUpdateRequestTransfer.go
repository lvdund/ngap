package ie

type MulticastSessionUpdateRequestTransfer struct {
MbsSessionId	MbsSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsServiceArea	MbsServiceArea	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsQosFlowsToBeSetupOrModifiedList	MbsQosFlowsToBeSetupList	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsQosFlowToReleaseList	QosFlowListWithCause	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSessionTnlInformation5Gc	MbsSessionTnlInformation5Gc	//`bitstring:"sizeLB:0,sizeUB:150"`
}
