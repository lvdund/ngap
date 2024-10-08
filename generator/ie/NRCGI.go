package ie

type NRCGI struct {
	PLMNIdentity   PLMNIdentity   `False,`
	NRCellIdentity NRCellIdentity `False,`
	IEExtensions   NRCGIExtIEs    `False,OPTIONAL`
}
