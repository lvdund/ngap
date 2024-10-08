package ie

type MbsDistributionSetupResponseTransfer struct {
	MbsSessionId                     MbsSessionId                     `bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId                 MbsAreaSessionId                 `bitstring:"sizeLB:0,sizeUB:150"`
	SharedNgUMulticastTnlInformation SharedNgUMulticastTnlInformation `bitstring:"sizeLB:0,sizeUB:150"`
	MbsQosFlowsToBeSetupList         MbsQosFlowsToBeSetupList         `bitstring:"sizeLB:0,sizeUB:150"`
	MbsSessionStatus                 MbsSessionStatus                 `bitstring:"sizeLB:0,sizeUB:150"`
	MbsServiceArea                   MbsServiceArea                   `bitstring:"sizeLB:0,sizeUB:150"`
}
