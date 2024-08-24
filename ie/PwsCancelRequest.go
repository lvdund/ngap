package ie

type PwsCancelRequest struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
MessageIdentifier	MessageIdentifier	//`bitstring:"sizeLB:0,sizeUB:150"`
SerialNumber	SerialNumber	//`bitstring:"sizeLB:0,sizeUB:150"`
WarningAreaList	WarningAreaList	//`bitstring:"sizeLB:0,sizeUB:150"`
CancelAllWarningMessagesIndicator	CancelAllWarningMessagesIndicator	//`bitstring:"sizeLB:0,sizeUB:150"`
}
