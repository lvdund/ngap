package ie

type CellIDBroadcastEUTRAItem struct {
	EUTRACGI     EUTRACGI                       `True,`
	IEExtensions CellIDBroadcastEUTRAItemExtIEs `False,OPTIONAL`
}
