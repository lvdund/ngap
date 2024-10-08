package ie

type SNSSAI struct {
	SST          SST          `False,`
	SD           SD           `False,OPTIONAL`
	IEExtensions SNSSAIExtIEs `False,OPTIONAL`
}
