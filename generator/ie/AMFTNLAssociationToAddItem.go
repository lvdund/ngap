package ie

type AMFTNLAssociationToAddItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation      `False,`
	TNLAssociationUsage      TNLAssociationUsage              `False,OPTIONAL`
	TNLAddressWeightFactor   TNLAddressWeightFactor           `False,`
	IEExtensions             AMFTNLAssociationToAddItemExtIEs `False,OPTIONAL`
}
