package ie

type SliceOverloadItem struct {
	SNSSAI       SNSSAI                  `True,`
	IEExtensions SliceOverloadItemExtIEs `False,OPTIONAL`
}
