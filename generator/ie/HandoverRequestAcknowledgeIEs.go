package ie

type HandoverRequestAcknowledgeIEs struct {
	AMFUENGAPID                              AMFUENGAPID                              `,ignore,mandatory`
	RANUENGAPID                              RANUENGAPID                              `,ignore,mandatory`
	PDUSessionResourceAdmittedList           PDUSessionResourceAdmittedList           `,ignore,mandatory`
	PDUSessionResourceFailedToSetupListHOAck PDUSessionResourceFailedToSetupListHOAck `,ignore,optional`
	TargetToSourceTransparentContainer       TargetToSourceTransparentContainer       `,reject,mandatory`
	CriticalityDiagnostics                   CriticalityDiagnostics                   `,ignore,optional`
}
