package ie

type AllowedNSSAIItem struct {
	SNSSAI       SNSSAI                 `True,`
	IEExtensions AllowedNSSAIItemExtIEs `False,OPTIONAL`
}
