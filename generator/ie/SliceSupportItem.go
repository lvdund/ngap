package ie

type SliceSupportItem struct {
	SNSSAI       SNSSAI                 `True,`
	IEExtensions SliceSupportItemExtIEs `False,OPTIONAL`
}
