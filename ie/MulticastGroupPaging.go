package ie

type MulticastGroupPaging struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsSessionId	MbsSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsServiceArea	MbsServiceArea	//`bitstring:"sizeLB:0,sizeUB:150"`
MulticastGroupPagingAreaList	[]MulticastGroupPagingAreaItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type MulticastGroupPagingAreaItem struct {
MulticastGroupPagingArea	MulticastGroupPagingArea	//`bitstring:"sizeLB:0,sizeUB:150"`
UePagingList	[]UePagingItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type UePagingItem struct {
UeIdentityIndexValue	UeIdentityIndexValue	//`bitstring:"sizeLB:0,sizeUB:150"`
PagingDrx	PagingDrx	//`bitstring:"sizeLB:0,sizeUB:150"`
}
