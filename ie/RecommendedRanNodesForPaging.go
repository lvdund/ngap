package ie

type RecommendedRanNodesForPaging struct {
RecommendedRanNodeList	[]RecommendedRanNodeItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type RecommendedRanNodeItem struct {
ChoiceAmfPagingTarget	ChoiceAmfPagingTarget	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type ChoiceAmfPagingTarget struct {
RanNode	RanNode	//`bitstring:"sizeLB:0,sizeUB:150"`
Tai	Tai	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type RanNode struct {
GlobalRanNodeId	GlobalRanNodeId	//`bitstring:"sizeLB:0,sizeUB:150"`
}
