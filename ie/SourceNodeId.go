package ie

type SourceNodeId struct {
	ChoiceSourceNodeId *ChoiceSourceNodeId
}

type ChoiceSourceNodeId struct {
	SourceEUtranNodeId *SourceEUtranNodeId
}

type SourceEUtranNodeId struct {
	SourceEnGnbId *GlobalGnbId
}
