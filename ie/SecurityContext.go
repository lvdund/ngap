package ie

type SecurityContext struct {
	NextHopChainingCount uint8 //`bitstring:"sizeLB:0,sizeUB:7"`
	NextHopNh            *SecurityKey
}
