package ie

type SONInformation struct {
	SONInformationRequest SONInformationRequest `False,`
	SONInformationReply   SONInformationReply   `True,`
	ChoiceExtensions      SONInformationExtIEs  `False,`
}
