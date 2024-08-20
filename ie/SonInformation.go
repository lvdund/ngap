package ie

type SonInformation struct {
ChoiceSonInformation	*ChoiceSonInformation
}

type ChoiceSonInformation struct {
SonInformationRequest	*SonInformationRequest
SonInformationReply	*SonInformationReply
SonInformationReport	*SonInformationReport
}

type SonInformationRequest struct {
SonInformationRequest	*[]byte
}
