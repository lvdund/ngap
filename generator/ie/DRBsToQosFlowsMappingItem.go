package ie

type DRBsToQosFlowsMappingItem struct {
	DRBID                 DRBID                           `False,`
	AssociatedQosFlowList AssociatedQosFlowList           `False,`
	IEExtensions          DRBsToQosFlowsMappingItemExtIEs `False,OPTIONAL`
}
