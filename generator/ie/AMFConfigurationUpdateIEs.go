package ie

type AMFConfigurationUpdateIEs struct {
	AMFName                       AMFName                       `,reject,optional`
	ServedGUAMIList               ServedGUAMIList               `,reject,optional`
	RelativeAMFCapacity           RelativeAMFCapacity           `,ignore,optional`
	PLMNSupportList               PLMNSupportList               `,reject,optional`
	AMFTNLAssociationToAddList    AMFTNLAssociationToAddList    `,ignore,optional`
	AMFTNLAssociationToRemoveList AMFTNLAssociationToRemoveList `,ignore,optional`
	AMFTNLAssociationToUpdateList AMFTNLAssociationToUpdateList `,ignore,optional`
}
