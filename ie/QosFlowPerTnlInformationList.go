package ie

type QosFlowPerTnlInformationList struct {
	QosFlowPerTnlInformationItem QosFlowPerTnlInformationItem //`bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowPerTnlInformationItem struct {
	QosFlowPerTnlInformation QosFlowPerTnlInformation //`bitstring:"sizeLB:0,sizeUB:150"`
}
