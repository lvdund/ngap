package ie

type UeContextResumeRequestTransfer struct {
	QosFlowFailedToResumeList QosFlowListWithCause `bitstring:"sizeLB:0,sizeUB:150"`
}
