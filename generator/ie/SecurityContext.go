package ie

type SecurityContext struct {
	NextHopChainingCount NextHopChainingCount  `False,`
	NextHopNH            SecurityKey           `False,`
	IEExtensions         SecurityContextExtIEs `False,OPTIONAL`
}
