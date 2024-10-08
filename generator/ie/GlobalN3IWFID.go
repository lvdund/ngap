package ie

type GlobalN3IWFID struct {
	PLMNIdentity PLMNIdentity        `False,`
	N3IWFID      N3IWFID             `False,`
	IEExtensions GlobalN3IWFIDExtIEs `False,OPTIONAL`
}
