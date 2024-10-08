package ie

type InfoOnRecommendedCellsAndRANNodesForPaging struct {
	RecommendedCellsForPaging  RecommendedCellsForPaging                        `True,`
	RecommendRANNodesForPaging RecommendedRANNodesForPaging                     `True,`
	IEExtensions               InfoOnRecommendedCellsAndRANNodesForPagingExtIEs `False,OPTIONAL`
}
