package ie

type QmcConfigurationInformation struct {
UeApplicationLayerMeasurementInformationList	[]UeApplicationLayerMeasurementInformationItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type UeApplicationLayerMeasurementInformationItem struct {
UeApplicationLayerMeasurementConfigurationInformation	UeApplicationLayerMeasurementConfigurationInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}
