package ie

type PagingAttemptInformation struct {
	PagingAttemptCount             PagingAttemptCount             `False,`
	IntendedNumberOfPagingAttempts IntendedNumberOfPagingAttempts `False,`
	NextPagingAreaScope            NextPagingAreaScope            `False,OPTIONAL`
	IEExtensions                   PagingAttemptInformationExtIEs `False,OPTIONAL`
}
