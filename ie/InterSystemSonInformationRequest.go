package ie

type InterSystemSonInformationRequest struct {
ChoiceInterSystemSonInformationRequest	ChoiceInterSystemSonInformationRequest	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceInterSystemSonInformationRequest struct {
NgRanCellActivation	NgRanCellActivation	//`bitstring:"sizeLB:0,sizeUB:150"`
ResourceStatus	ResourceStatus	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type NgRanCellActivation struct {
InterSystemCellActivationRequest	InterSystemCellActivationRequest	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ResourceStatus struct {
InterSystemResourceStatusRequest	InterSystemResourceStatusRequest	//`bitstring:"sizeLB:0,sizeUB:150"`
}
