package ie

type AssistanceDataForPaging struct {
	AssistanceDataForRecommendedCells AssistanceDataForRecommendedCells `True,OPTIONAL`
	PagingAttemptInformation          PagingAttemptInformation          `True,OPTIONAL`
	IEExtensions                      AssistanceDataForPagingExtIEs     `False,OPTIONAL`
}
