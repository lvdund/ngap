package ie

type ExpectedUeActivityBehaviour struct {
	ExpectedActivityPeriod                 uint8 //`bitstring:"sizeLB:1,sizeUB:30"`
	ExpectedIdlePeriod                     uint8 //`bitstring:"sizeLB:1,sizeUB:30"`
	SourceOfUeActivityBehaviourInformation *[]byte
}
