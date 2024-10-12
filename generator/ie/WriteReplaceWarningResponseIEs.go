package ie

type WriteReplaceWarningResponseIEs struct {
	MessageIdentifier          MessageIdentifier          `,reject,mandatory`
	SerialNumber               SerialNumber               `,reject,mandatory`
	BroadcastCompletedAreaList BroadcastCompletedAreaList `,ignore,optional`
	CriticalityDiagnostics     CriticalityDiagnostics     `,ignore,optional`
}
