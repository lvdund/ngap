package ie

type NrRadioResourceStatus struct {
DlGbrPrbUsageForMimo	uint8	//`bitstring:"sizeLB:0,sizeUB:100"`
UlGbrPrbUsageForMimo	uint8	//`bitstring:"sizeLB:0,sizeUB:100"`
DlNonGbrPrbUsageForMimo	uint8	//`bitstring:"sizeLB:0,sizeUB:100"`
UlNonGbrPrbUsageForMimo	uint8	//`bitstring:"sizeLB:0,sizeUB:100"`
DlTotalPrbUsageForMimo	uint8	//`bitstring:"sizeLB:0,sizeUB:100"`
UlTotalPrbUsageForMimo	uint8	//`bitstring:"sizeLB:0,sizeUB:100"`
}
