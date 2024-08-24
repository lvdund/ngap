package ie

type SourceNodeId struct {
	ChoiceSourceNodeId ChoiceSourceNodeId //`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceSourceNodeId struct {
	SourceEUtranNodeId SourceEUtranNodeId //`bitstring:"sizeLB:0,sizeUB:150"`
}

type SourceEUtranNodeId struct {
	SourceEnGnbId GlobalGnbId //`bitstring:"sizeLB:0,sizeUB:150"`
}
