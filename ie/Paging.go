package ie

type Paging struct {
MessageType	*MessageType
UePagingIdentity	*UePagingIdentity
PagingDrx	*PagingDrx
TaiListForPaging	*TaiListForPaging
PagingPriority	*PagingPriority
UeRadioCapabilityForPaging	*UeRadioCapabilityForPaging
PagingOrigin	*PagingOrigin
AssistanceDataForPaging	*AssistanceDataForPaging
NbIotPagingEdrxInformation	*NbIotPagingEdrxInformation
NbIotPagingDrx	*NbIotPagingDrx
EnhancedCoverageRestriction	*EnhancedCoverageRestriction
WusAssistanceInformation	*WusAssistanceInformation
EUtraPagingEdrxInformation	*EUtraPagingEdrxInformation
CeModeBRestricted	*CeModeBRestricted
NrPagingEdrxInformation	*NrPagingEdrxInformation
PagingCause	*[]byte
PeipsAssistanceInformation	*PeipsAssistanceInformation
}

type TaiListForPaging struct {
TaiListForPagingItem	*TaiListForPagingItem
}

type TaiListForPagingItem struct {
Tai	*Tai
}
