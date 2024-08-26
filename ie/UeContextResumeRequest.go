package ie

type UeContextResumeRequest struct {
	MessageType                                       MessageType                                       `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                                       AmfUeNgapId                                       `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                                       RanUeNgapId                                       `bitstring:"sizeLB:0,sizeUB:150"`
	RrcResumeCause                                    RrcEstablishmentCause                             `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceResumeList                      []PduSessionResourceResumeItem                    `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceFailedToResumeList              []PduSessionResourceFailedToResumeItem            `bitstring:"sizeLB:0,sizeUB:150"`
	SuspendRequestIndication                          SuspendRequestIndication                          `bitstring:"sizeLB:0,sizeUB:150"`
	InformationOnRecommendedCellsAndRanNodesForPaging InformationOnRecommendedCellsAndRanNodesForPaging `bitstring:"sizeLB:0,sizeUB:150"`
	PagingAssistanceDataForCeCapableUe                PagingAssistanceDataForCeCapableUe                `bitstring:"sizeLB:0,sizeUB:150"`
}
