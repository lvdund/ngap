package ie

type UeRadioCapabilityIdMappingResponse struct {
	MessageType            MessageType            //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityId    UeRadioCapabilityId    //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapability      UeRadioCapability      //`bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics CriticalityDiagnostics //`bitstring:"sizeLB:0,sizeUB:150"`
}
