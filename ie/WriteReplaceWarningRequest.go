package ie

type WriteReplaceWarningRequest struct {
MessageType	*MessageType
MessageIdentifier	*MessageIdentifier
SerialNumber	*SerialNumber
WarningAreaList	*WarningAreaList
RepetitionPeriod	*RepetitionPeriod
NumberOfBroadcastsRequested	*NumberOfBroadcastsRequested
WarningType	*WarningType
WarningSecurityInformation	[]byte	//`bitstring:"sizeLB:50,sizeUB:50"`
DataCodingScheme	*DataCodingScheme
WarningMessageContents	*WarningMessageContents
ConcurrentWarningMessageIndicator	*ConcurrentWarningMessageIndicator
WarningAreaCoordinates	*WarningAreaCoordinates
}
