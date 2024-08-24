package ie

type AssistanceDataForPaging struct {
	AssistanceDataForRecommendedCells  AssistanceDataForRecommendedCells  //`bitstring:"sizeLB:0,sizeUB:150"`
	PagingAttemptInformation           PagingAttemptInformation           //`bitstring:"sizeLB:0,sizeUB:150"`
	NpnPagingAssistanceInformation     NpnPagingAssistanceInformation     //`bitstring:"sizeLB:0,sizeUB:150"`
	PagingAssistanceDataForCeCapableUe PagingAssistanceDataForCeCapableUe //`bitstring:"sizeLB:0,sizeUB:150"`
}
