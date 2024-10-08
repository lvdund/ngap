package ie

type TAIListForInactiveItem struct {
	TAI          TAI                          `True,`
	IEExtensions TAIListForInactiveItemExtIEs `False,OPTIONAL`
}
