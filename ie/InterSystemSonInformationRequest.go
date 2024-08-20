package ie

type InterSystemSonInformationRequest struct {
	ChoiceInterSystemSonInformationRequest *ChoiceInterSystemSonInformationRequest
}

type ChoiceInterSystemSonInformationRequest struct {
	NgRanCellActivation *NgRanCellActivation
	ResourceStatus      *ResourceStatus
}

type NgRanCellActivation struct {
	InterSystemCellActivationRequest *InterSystemCellActivationRequest
}

type ResourceStatus struct {
	InterSystemResourceStatusRequest *InterSystemResourceStatusRequest
}
