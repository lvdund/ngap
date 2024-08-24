package ie

type TscAssistanceInformation struct {
	Periodicity      Periodicity      //`bitstring:"sizeLB:0,sizeUB:150"`
	BurstArrivalTime BurstArrivalTime //`bitstring:"sizeLB:0,sizeUB:150"`
	SurvivalTime     SurvivalTime     //`bitstring:"sizeLB:0,sizeUB:150"`
}
