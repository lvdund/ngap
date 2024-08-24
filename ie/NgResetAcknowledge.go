package ie

type NgResetAcknowledge struct {
	MessageType                         MessageType                         //`bitstring:"sizeLB:0,sizeUB:150"`
	UeAssociatedLogicalNgConnectionList UeAssociatedLogicalNgConnectionList //`bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics              CriticalityDiagnostics              //`bitstring:"sizeLB:0,sizeUB:150"`
}
