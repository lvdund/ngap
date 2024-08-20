package ie

type PwsCancelRequest struct {
MessageType	*MessageType
MessageIdentifier	*MessageIdentifier
SerialNumber	*SerialNumber
WarningAreaList	*WarningAreaList
CancelAllWarningMessagesIndicator	*CancelAllWarningMessagesIndicator
}
