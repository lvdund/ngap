package ie

type Cause struct {
	RadioNetwork     CauseRadioNetwork `False,`
	Transport        CauseTransport    `False,`
	Nas              CauseNas          `False,`
	Protocol         CauseProtocol     `False,`
	Misc             CauseMisc         `False,`
	ChoiceExtensions CauseExtIEs       `False,`
}
