package ie

import "ngap/aper"

type EventTriggerLoggedMdtConfiguration struct {
	ChoiceEventTriggerType ChoiceEventTriggerType //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceEventTriggerType struct {
	OutOfCoverage OutOfCoverage //`bitstring:"sizeLB:0,sizeUB:150"`
	L1Event       L1Event       //`bitstring:"sizeLB:0,sizeUB:150"`
}

type OutOfCoverage struct {
	OutOfCoverageConfiguration []byte //`bitstring:"sizeLB:0,sizeUB:150"`
}

type L1Event struct {
	ChoiceL1EventThreshold ChoiceL1EventThreshold //`bitstring:"sizeLB:0,sizeUB:150"`
	Hysteresis             aper.Integer           //`Integer:"valueLB:0,valueUB:30"`
	TimeToTrigger          []byte                 //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceL1EventThreshold struct {
	Rsrp Rsrp //`bitstring:"sizeLB:0,sizeUB:150"`
	Rsrq Rsrq //`bitstring:"sizeLB:0,sizeUB:150"`
}
