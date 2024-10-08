package ie

type UeContextSuspendRequest struct {
	MessageType                                       MessageType                                       `bitstring:"sizeLB:0,sizeUB:150"`
	AmfUeNgapId                                       AmfUeNgapId                                       `bitstring:"sizeLB:0,sizeUB:150"`
	RanUeNgapId                                       RanUeNgapId                                       `bitstring:"sizeLB:0,sizeUB:150"`
	InformationOnRecommendedCellsAndRanNodesForPaging InformationOnRecommendedCellsAndRanNodesForPaging `bitstring:"sizeLB:0,sizeUB:150"`
	PagingAssistanceDataForCeCapableUe                PagingAssistanceDataForCeCapableUe                `bitstring:"sizeLB:0,sizeUB:150"`
	PduSessionResourceSuspendList                     []PduSessionResourceSuspendItem                   `bitstring:"sizeLB:0,sizeUB:150"`
}

type PduSessionResourceSuspendItem struct {
	PduSessionId                    PduSessionId                    `bitstring:"sizeLB:0,sizeUB:150"`
	UeContextSuspendRequestTransfer UeContextSuspendRequestTransfer `bitstring:"sizeLB:0,sizeUB:150"`
}
