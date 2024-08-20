package ie

type PagingAttemptInformation struct {
	PagingAttemptCount             uint8 //`bitstring:"sizeLB:1,sizeUB:16"`
	IntendedNumberOfPagingAttempts uint8 //`bitstring:"sizeLB:1,sizeUB:16"`
	NextPagingAreaScope            *[]byte
}
