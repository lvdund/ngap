package ie

type PathSwitchRequestAcknowledgeTransfer struct {
	UlNgUUpTnlInformation                  UpTransportLayerInformation         `bitstring:"sizeLB:0,sizeUB:150"`
	SecurityIndication                     SecurityIndication                  `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalNgUUpTnlInformation          UpTransportLayerInformationPairList `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantUlNgUUpTnlInformation         UpTransportLayerInformation         `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalRedundantNgUUpTnlInformation UpTransportLayerInformationPairList `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowParametersList                  []QosFlowParametersItem             `bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowParametersItem struct {
	QosFlowIdentifier               QosFlowIdentifier               `bitstring:"sizeLB:0,sizeUB:150"`
	AlternativeQosParametersSetList AlternativeQosParametersSetList `bitstring:"sizeLB:0,sizeUB:150"`
	CnPacketDelayBudgetDownlink     ExtendedPacketDelayBudget       `bitstring:"sizeLB:0,sizeUB:150"`
	CnPacketDelayBudgetUplink       ExtendedPacketDelayBudget       `bitstring:"sizeLB:0,sizeUB:150"`
	BurstArrivalTimeDownlink        BurstArrivalTime                `bitstring:"sizeLB:0,sizeUB:150"`
}
