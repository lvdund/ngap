package ie

type MbsDistributionSetupRequestTransfer struct {
	MbsSessionId                   MbsSessionId                `bitstring:"sizeLB:0,sizeUB:150"`
	MbsAreaSessionId               MbsAreaSessionId            `bitstring:"sizeLB:0,sizeUB:150"`
	SharedNgUUnicastTnlInformation UpTransportLayerInformation `bitstring:"sizeLB:0,sizeUB:150"`
}
