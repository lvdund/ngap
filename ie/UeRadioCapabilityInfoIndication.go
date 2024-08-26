package ie

type UeRadioCapabilityInfoIndication struct {
	MessageType                  MessageType                  `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                  AmfUeNgapId                  `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                  RanUeNgapId                  `bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapability            UeRadioCapability            `bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityForPaging   UeRadioCapabilityForPaging   `bitstring:"sizeLB:0,sizeUB:150"`
	UeRadioCapabilityEUtraFormat UeRadioCapabilityEUtraFormat `bitstring:"sizeLB:0,sizeUB:150"`
}
