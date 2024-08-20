package ie

type InterSystemCellStateIndication struct {
NotificationCellList	*[]NotificationCellItem
}

type NotificationCellItem struct {
NgRanCgi	*NgRanCgi
NotifyFlag	*[]byte
}
