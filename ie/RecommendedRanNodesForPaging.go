package ie

type RecommendedRanNodesForPaging struct {
RecommendedRanNodeList	*[]RecommendedRanNodeItem
}

type RecommendedRanNodeItem struct {
ChoiceAmfPagingTarget	*ChoiceAmfPagingTarget
}

type ChoiceAmfPagingTarget struct {
RanNode	*RanNode
Tai	*Tai
}

type RanNode struct {
GlobalRanNodeId	*GlobalRanNodeId
}
