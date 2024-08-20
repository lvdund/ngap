package ie

type UeContextResumeRequest struct {
	MessageType                                       *MessageType
	AmfUeNgapId                                       *AmfUeNgapId
	RanUeNgapId                                       *RanUeNgapId
	RrcResumeCause                                    *RrcEstablishmentCause
	PduSessionResourceResumeList                      *[]PduSessionResourceResumeItem
	PduSessionResourceFailedToResumeList              *[]PduSessionResourceFailedToResumeItem
	SuspendRequestIndication                          *SuspendRequestIndication
	InformationOnRecommendedCellsAndRanNodesForPaging *InformationOnRecommendedCellsAndRanNodesForPaging
	PagingAssistanceDataForCeCapableUe                *PagingAssistanceDataForCeCapableUe
}
