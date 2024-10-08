package ie

type TNLAssociationItem struct {
	TNLAssociationAddress CPTransportLayerInformation `False,`
	Cause                 Cause                       `False,`
	IEExtensions          TNLAssociationItemExtIEs    `False,OPTIONAL`
}
