package ie

type UeRadioCapabilityCheckRequest struct {
	MessageType         MessageType         //`bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId         AmfUeNgapId         //`bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId         RanUeNgapId         //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapability   UeRadioCapability   //`bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityId UeRadioCapabilityId //`bitstring:"sizeLB:0,sizeUB:150"`
}
