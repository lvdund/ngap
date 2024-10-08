package ie

type CompletedCellsInEAINRItem struct {
	NRCGI        NRCGI                           `True,`
	IEExtensions CompletedCellsInEAINRItemExtIEs `False,OPTIONAL`
}
