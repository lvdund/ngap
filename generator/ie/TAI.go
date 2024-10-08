package ie

type TAI struct {
	PLMNIdentity PLMNIdentity `False,`
	TAC          TAC          `False,`
	IEExtensions TAIExtIEs    `False,OPTIONAL`
}
