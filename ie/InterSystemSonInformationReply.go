package ie

type InterSystemSonInformationReply struct {
	ChoiceInterSystemSonInformationReply ChoiceInterSystemSonInformationReply //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceInterSystemSonInformationReply struct {
	NgRanCellActivation NgRanCellActivation //`bitstring:"sizeLB:0,sizeUB:150"`
	ResourceStatus      ResourceStatus      //`bitstring:"sizeLB:0,sizeUB:150"`
}
