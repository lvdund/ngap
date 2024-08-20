package ie

type EUtranRadioResourceStatus struct {
	DlGbrPrbUsage             uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	UlGbrPrbUsage             uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	DlNonGbrPrbUsage          uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	UlNonGbrPrbUsage          uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	DlTotalPrbUsage           uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	UlTotalPrbUsage           uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	DlSchedulingPdcchCceUsage uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
	UlSchedulingPdcchCceUsage uint8 //`bitstring:"sizeLB:0,sizeUB:100"`
}
