package ie

type UserPlaneSecurityInformation struct {
SecurityResult	SecurityResult	//`bitstring:"sizeLB:0,sizeUB:150"`
SecurityIndication	SecurityIndication	//`bitstring:"sizeLB:0,sizeUB:150"`
}
