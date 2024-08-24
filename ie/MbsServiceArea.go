package ie

type MbsServiceArea struct {
ChoiceSessionType	ChoiceSessionType	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type MbsServiceAreaInformationItem struct {
MbsAreaSessionId	MbsAreaSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsServiceAreaInformation	MbsServiceAreaInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}
