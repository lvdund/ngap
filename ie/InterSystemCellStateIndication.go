package ie

type InterSystemCellStateIndication struct {
NotificationCellList	[]NotificationCellItem	//`bitstring:"sizeLB:0,sizeUB:150"`
}

type NotificationCellItem struct {
NgRanCgi	NgRanCgi	//`bitstring:"sizeLB:0,sizeUB:150"`
NotifyFlag	[]byte	//`bitstring:"sizeLB:0,sizeUB:150"`
}
