package ie

type PWSCancelResponseIEs struct {
	MessageIdentifier          MessageIdentifier          `,reject,mandatory`
	SerialNumber               SerialNumber               `,reject,mandatory`
	BroadcastCancelledAreaList BroadcastCancelledAreaList `,ignore,optional`
	CriticalityDiagnostics     CriticalityDiagnostics     `,ignore,optional`
}
