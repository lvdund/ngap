package ie

type AMFPagingTarget struct {
	GlobalRANNodeID  GlobalRANNodeID       `False,`
	TAI              TAI                   `True,`
	ChoiceExtensions AMFPagingTargetExtIEs `False,`
}
