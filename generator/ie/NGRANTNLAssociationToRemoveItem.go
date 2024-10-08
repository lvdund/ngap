package ie

type NGRANTNLAssociationToRemoveItem struct {
	TNLAssociationTransportLayerAddress    CPTransportLayerInformation           `False,`
	TNLAssociationTransportLayerAddressAMF CPTransportLayerInformation           `False,OPTIONAL`
	IEExtensions                           NGRANTNLAssociationToRemoveItemExtIEs `False,OPTIONAL`
}
