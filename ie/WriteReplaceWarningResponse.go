package ie

type WriteReplaceWarningResponse struct {
	MessageType                MessageType                //`bitstring:"sizeLB:0,sizeUB:150"`
	MessageIdentifier          MessageIdentifier          //`bitstring:"sizeLB:0,sizeUB:150"`
	SerialNumber               SerialNumber               //`bitstring:"sizeLB:0,sizeUB:150"`
	BroadcastCompletedAreaList BroadcastCompletedAreaList //`bitstring:"sizeLB:0,sizeUB:150"`
	CriticalityDiagnostics     CriticalityDiagnostics     //`bitstring:"sizeLB:0,sizeUB:150"`
}
