package ie

type UserLocationInformationN3IWF struct {
	IPAddress    TransportLayerAddress              `False,`
	PortNumber   PortNumber                         `False,`
	IEExtensions UserLocationInformationN3IWFExtIEs `False,OPTIONAL`
}
