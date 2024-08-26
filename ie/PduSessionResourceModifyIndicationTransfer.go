package ie

type PduSessionResourceModifyIndicationTransfer struct {
	DlQosFlowPerTnlInformation                    QosFlowPerTnlInformation     `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalDlQosFlowPerTnlInformation          QosFlowPerTnlInformationList `bitstring:"sizeLB:0,sizeUB:150"`
	SecondaryRatUsageInformation                  SecondaryRatUsageInformation `bitstring:"sizeLB:0,sizeUB:150"`
	SecurityResult                                SecurityResult               `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantDlQosFlowPerTnlInformation           QosFlowPerTnlInformation     `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalRedundantDlQosFlowPerTnlInformation QosFlowPerTnlInformationList `bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeIdOfSecondaryNgRanNode           GlobalRanNodeId              `bitstring:"sizeLB:0,sizeUB:150"`
}
