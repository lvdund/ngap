package ie

type InterSystemSonInformation struct {
ChoiceInterSystemSonInformation	ChoiceInterSystemSonInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceInterSystemSonInformation struct {
InterSystemSonInformationReport	InterSystemSonInformationReport	//`bitstring:"sizeLB:0,sizeUB:150"`
InterSystemSonInformationRequest	InterSystemSonInformationRequest	//`bitstring:"sizeLB:0,sizeUB:150"`
InterSystemSonInformationReply	InterSystemSonInformationReply	//`bitstring:"sizeLB:0,sizeUB:150"`
}
