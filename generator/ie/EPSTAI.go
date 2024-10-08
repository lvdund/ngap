package ie

type EPSTAI struct {
	PLMNIdentity PLMNIdentity `False,`
	EPSTAC       EPSTAC       `False,`
	IEExtensions EPSTAIExtIEs `False,OPTIONAL`
}
