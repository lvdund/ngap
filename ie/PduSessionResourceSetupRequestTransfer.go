package ie

type PduSessionResourceSetupRequestTransfer struct {
	PduSessionAggregateMaximumBitRate        PduSessionAggregateMaximumBitRate `bitstring:"sizeLB:0,sizeUB:150"`
	UlNgUUpTnlInformation                    UpTransportLayerInformation       `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalUlNgUUpTnlInformation          UpTransportLayerInformationList   `bitstring:"sizeLB:0,sizeUB:150"`
	DataForwardingNotPossible                DataForwardingNotPossible         `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionType                           PduSessionType                    `bitstring:"sizeLB:0,sizeUB:150"`
	SecurityIndication                       SecurityIndication                `bitstring:"sizeLB:0,sizeUB:150"`
	NetworkInstance                          NetworkInstance                   `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowSetupRequestList                  []QosFlowSetupRequestItem         `bitstring:"sizeLB:0,sizeUB:150"`
	CommonNetworkInstance                    CommonNetworkInstance             `bitstring:"sizeLB:0,sizeUB:150"`
	DirectForwardingPathAvailability         DirectForwardingPathAvailability  `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantUlNgUUpTnlInformation           UpTransportLayerInformation       `bitstring:"sizeLB:0,sizeUB:150"`
	AdditionalRedundantUlNgUUpTnlInformation UpTransportLayerInformationList   `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantCommonNetworkInstance           CommonNetworkInstance             `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantPduSessionInformation           RedundantPduSessionInformation    `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionSetupRequestList               MbsSessionSetupRequestList        `bitstring:"sizeLB:0,sizeUB:150"`
}

type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         QosFlowIdentifier         `bitstring:"sizeLB:0,sizeUB:150"`
	QosFlowLevelQosParameters QosFlowLevelQosParameters `bitstring:"sizeLB:0,sizeUB:150"`
	ERabId                    ERabId                    `bitstring:"sizeLB:0,sizeUB:150"`
	TscTrafficCharacteristics TscTrafficCharacteristics `bitstring:"sizeLB:0,sizeUB:150"`
	RedundantQosFlowIndicator RedundantQosFlowIndicator `bitstring:"sizeLB:0,sizeUB:150"`
}
