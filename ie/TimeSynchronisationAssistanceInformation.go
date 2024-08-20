package ie

type TimeSynchronisationAssistanceInformation struct {
TimeDistributionIndication	*[]byte
UuTimeSynchronisationErrorBudget	uint32	//`bitstring:"sizeLB:1,sizeUB:1000000"`
}
