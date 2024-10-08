package ie

type AMFTNLAssociationToUpdateItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation         `False,`
	TNLAssociationUsage      TNLAssociationUsage                 `False,OPTIONAL`
	TNLAddressWeightFactor   TNLAddressWeightFactor              `False,OPTIONAL`
	IEExtensions             AMFTNLAssociationToUpdateItemExtIEs `False,OPTIONAL`
}
