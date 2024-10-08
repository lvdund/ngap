package ie

type QosCharacteristics struct {
	NonDynamic5QI    NonDynamic5QIDescriptor  `True,`
	Dynamic5QI       Dynamic5QIDescriptor     `True,`
	ChoiceExtensions QosCharacteristicsExtIEs `False,`
}
