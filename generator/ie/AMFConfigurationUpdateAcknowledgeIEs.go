package ie

type AMFConfigurationUpdateAcknowledgeIEs struct {
	AMFTNLAssociationSetupList         AMFTNLAssociationSetupList `,ignore,optional`
	AMFTNLAssociationFailedToSetupList TNLAssociationList         `,ignore,optional`
	CriticalityDiagnostics             CriticalityDiagnostics     `,ignore,optional`
}
