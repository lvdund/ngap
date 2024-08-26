package ie

type PathSwitchRequestTransfer struct {
	DlNgUUpTnlInformation                         UpTransportLayerInformation    `bitstring:"sizeLB:0,sizeUB:150"`
	DlNgUTnlInformationReused                     []byte                         `bitstring:"sizeLB:0,sizeUB:150"`
	UserPlaneSecurityInformation                  UserPlaneSecurityInformation   `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowAcceptedList                           []QosFlowAcceptedItem          `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalDlQosFlowPerTnlInformation          QosFlowPerTnlInformationList   `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantDlNgUUpTnlInformation                UpTransportLayerInformation    `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantDlNgUTnlInformationReused            []byte                         `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalRedundantDlQosFlowPerTnlInformation QosFlowPerTnlInformationList   `bitstring:"sizeLB:0,sizeUB:150"`
	UsedRsnInformation                            RedundantPduSessionInformation `bitstring:"sizeLB:0,sizeUB:150"`
	GlobalRanNodeIdOfSecondaryNgRanNode           GlobalRanNodeId                `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSupportIndicator                           MbsSupportIndicator            `bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowAcceptedItem struct {
	QosFlowIdentifier            QosFlowIdentifier                `bitstring:"sizeLB:0,sizeUB:150"`
	CurrentQosParametersSetIndex AlternativeQosParametersSetIndex `bitstring:"sizeLB:0,sizeUB:150"`
}
