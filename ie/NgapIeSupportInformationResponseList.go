package ie

type NgapIeSupportInformationResponseList struct {
NgapIeSupportInformationResponseItem	NgapIeSupportInformationResponseItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgapIeSupportInformationResponseItem struct {
NgapProtocolIeId	NgapProtocolIeId	//`bitstring:"sizeLB:0,sizeUB:150"`
NgapProtocolIeSupportInformation	NgapProtocolIeSupportInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
NgapProtocolIePresenceInformation	NgapProtocolIePresenceInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}
