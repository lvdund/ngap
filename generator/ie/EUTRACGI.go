package ie

type EUTRACGI struct {
	PLMNIdentity      PLMNIdentity      `False,`
	EUTRACellIdentity EUTRACellIdentity `False,`
	IEExtensions      EUTRACGIExtIEs    `False,OPTIONAL`
}
