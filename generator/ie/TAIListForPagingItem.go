package ie

type TAIListForPagingItem struct {
	TAI          TAI                        `True,`
	IEExtensions TAIListForPagingItemExtIEs `False,OPTIONAL`
}
