package ie

type UeContextReleaseComplete struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
AmfUeNgapId	AmfUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
RanUeNgapId	RanUeNgapId	//`bitstring:"sizeLB:0,sizeUB:150"`
UserLocationInformation	UserLocationInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
InformationOnRecommendedCellsAndRanNodesForPaging	InformationOnRecommendedCellsAndRanNodesForPaging	//`bitstring:"sizeLB:0,sizeUB:150"`
PduSessionResourceList	[]PduSessionResourceItem	//`bitstring:"sizeLB:0,sizeUB:150"`
CriticalityDiagnostics	CriticalityDiagnostics	//`bitstring:"sizeLB:0,sizeUB:150"`
PagingAssistanceDataForCeCapableUe	PagingAssistanceDataForCeCapableUe	//`bitstring:"sizeLB:0,sizeUB:150"`
}
