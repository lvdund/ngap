package ie

type RANConfigurationUpdateIEs struct {
	RANNodeName                     RANNodeName                     `,ignore,optional`
	SupportedTAList                 SupportedTAList                 `,reject,optional`
	DefaultPagingDRX                PagingDRX                       `,ignore,optional`
	GlobalRANNodeID                 GlobalRANNodeID                 `,ignore,optional`
	NGRANTNLAssociationToRemoveList NGRANTNLAssociationToRemoveList `,reject,optional`
}
