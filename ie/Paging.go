package ie

type Paging struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
UePagingIdentity	UePagingIdentity	//`bitstring:"sizeLB:0,sizeUB:150"`
PagingDrx	PagingDrx	//`bitstring:"sizeLB:0,sizeUB:150"`
TaiListForPaging	TaiListForPaging	//`bitstring:"sizeLB:0,sizeUB:150"`
PagingPriority	PagingPriority	//`bitstring:"sizeLB:0,sizeUB:150"`
UeRadioCapabilityForPaging	UeRadioCapabilityForPaging	//`bitstring:"sizeLB:0,sizeUB:150"`
PagingOrigin	PagingOrigin	//`bitstring:"sizeLB:0,sizeUB:150"`
AssistanceDataForPaging	AssistanceDataForPaging	//`bitstring:"sizeLB:0,sizeUB:150"`
NbIotPagingEdrxInformation	NbIotPagingEdrxInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
NbIotPagingDrx	NbIotPagingDrx	//`bitstring:"sizeLB:0,sizeUB:150"`
EnhancedCoverageRestriction	EnhancedCoverageRestriction	//`bitstring:"sizeLB:0,sizeUB:150"`
WusAssistanceInformation	WusAssistanceInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
EUtraPagingEdrxInformation	EUtraPagingEdrxInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
CeModeBRestricted	CeModeBRestricted	//`bitstring:"sizeLB:0,sizeUB:150"`
NrPagingEdrxInformation	NrPagingEdrxInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
PagingCause	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
PeipsAssistanceInformation	PeipsAssistanceInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiListForPaging struct {
TaiListForPagingItem	TaiListForPagingItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type TaiListForPagingItem struct {
Tai	Tai	//`bitstring:"sizeLB:0,sizeUB:150"`
}
