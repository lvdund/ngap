package ie

type AMFTNLAssociationSetupItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation      `False,`
	IEExtensions             AMFTNLAssociationSetupItemExtIEs `False,OPTIONAL`
}
