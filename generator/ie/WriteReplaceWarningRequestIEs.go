package ie

type WriteReplaceWarningRequestIEs struct {
	MessageIdentifier           MessageIdentifier           `,reject,mandatory`
	SerialNumber                SerialNumber                `,reject,mandatory`
	WarningAreaList             WarningAreaList             `,ignore,optional`
	RepetitionPeriod            RepetitionPeriod            `,reject,mandatory`
	NumberOfBroadcastsRequested NumberOfBroadcastsRequested `,reject,mandatory`
	WarningType                 WarningType                 `,ignore,optional`
	WarningSecurityInfo         WarningSecurityInfo         `,ignore,optional`
	DataCodingScheme            DataCodingScheme            `,ignore,optional`
	WarningMessageContents      WarningMessageContents      `,ignore,optional`
	ConcurrentWarningMessageInd ConcurrentWarningMessageInd `,reject,optional`
	WarningAreaCoordinates      WarningAreaCoordinates      `,ignore,optional`
}
