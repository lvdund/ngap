package ie

type SonInformation struct {
	ChoiceSonInformation ChoiceSonInformation `bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceSonInformation struct {
	SonInformationRequest SonInformationRequest `bitstring:"sizeLB:0,sizeUB:150"`
	SonInformationReply   SonInformationReply   `bitstring:"sizeLB:0,sizeUB:150"`
	SonInformationReport  SonInformationReport  `bitstring:"sizeLB:0,sizeUB:150"`
}

type SonInformationRequest struct {
	SonInformationRequest []byte `bitstring:"sizeLB:0,sizeUB:150"`
}
