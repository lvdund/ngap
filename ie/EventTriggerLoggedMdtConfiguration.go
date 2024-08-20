package ie

type EventTriggerLoggedMdtConfiguration struct {
	ChoiceEventTriggerType *ChoiceEventTriggerType
}

type ChoiceEventTriggerType struct {
	OutOfCoverage *OutOfCoverage
	L1Event       *L1Event
}

type OutOfCoverage struct {
	OutOfCoverageConfiguration *[]byte
}

type L1Event struct {
	ChoiceL1EventThreshold *ChoiceL1EventThreshold
	Hysteresis             uint8 //`bitstring:"sizeLB:0,sizeUB:30"`
	TimeToTrigger          *[]byte
}

type ChoiceL1EventThreshold struct {
	Rsrp *Rsrp
	Rsrq *Rsrq
}
