package ie

type MulticastGroupPaging struct {
MessageType	*MessageType
MbsSessionId	*MbsSessionId
MbsServiceArea	*MbsServiceArea
MulticastGroupPagingAreaList	*[]MulticastGroupPagingAreaItem
}

type MulticastGroupPagingAreaItem struct {
MulticastGroupPagingArea	*MulticastGroupPagingArea
UePagingList	*[]UePagingItem
}

type UePagingItem struct {
UeIdentityIndexValue	*UeIdentityIndexValue
PagingDrx	*PagingDrx
}
