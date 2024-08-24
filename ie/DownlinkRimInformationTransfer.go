package ie

type DownlinkRimInformationTransfer struct {
MessageType	MessageType	//`bitstring:"sizeLB:0,sizeUB:150"`
RimInformationTransfer	RimInformationTransfer	//`bitstring:"sizeLB:0,sizeUB:150"`
}
