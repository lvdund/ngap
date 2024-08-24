package ie

type RimInformationTransfer struct {
	TargetRanNodeIdRim TargetRanNodeIdRim //`bitstring:"sizeLB:0,sizeUB:150"`
	SourceRanNodeId    SourceRanNodeId    //`bitstring:"sizeLB:0,sizeUB:150"`
	RimInformation     RimInformation     //`bitstring:"sizeLB:0,sizeUB:150"`
}

type TargetRanNodeIdRim struct {
	GlobalRanNodeId GlobalRanNodeId //`bitstring:"sizeLB:0,sizeUB:150"`
	SelectedTai     Tai             //`bitstring:"sizeLB:0,sizeUB:150"`
}

type SourceRanNodeId struct {
	GlobalRanNodeId GlobalRanNodeId //`bitstring:"sizeLB:0,sizeUB:150"`
	SelectedTai     Tai             //`bitstring:"sizeLB:0,sizeUB:150"`
}
