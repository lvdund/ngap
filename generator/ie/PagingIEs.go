package ie

type PagingIEs struct {
	UEPagingIdentity           UEPagingIdentity           `,ignore,mandatory`
	PagingDRX                  PagingDRX                  `,ignore,optional`
	TAIListForPaging           TAIListForPaging           `,ignore,mandatory`
	PagingPriority             PagingPriority             `,ignore,optional`
	UERadioCapabilityForPaging UERadioCapabilityForPaging `,ignore,optional`
	PagingOrigin               PagingOrigin               `,ignore,optional`
	AssistanceDataForPaging    AssistanceDataForPaging    `,ignore,optional`
}
