package ie

type CountValueForPdcpSnLength12 struct {
	PdcpSnLength12       uint16 //`bitstring:"sizeLB:0,sizeUB:4095"`
	HfnForPdcpSnLength12 uint32 //`bitstring:"sizeLB:0,sizeUB:1048575"`
}
