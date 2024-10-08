package ie

type AreaOfInterestTAIItem struct {
	TAI          TAI                         `True,`
	IEExtensions AreaOfInterestTAIItemExtIEs `False,OPTIONAL`
}
