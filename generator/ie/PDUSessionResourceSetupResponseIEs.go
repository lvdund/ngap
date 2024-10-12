package ie

type PDUSessionResourceSetupResponseIEs struct {
	AMFUENGAPID                              AMFUENGAPID                              `,ignore,mandatory`
	RANUENGAPID                              RANUENGAPID                              `,ignore,mandatory`
	PDUSessionResourceSetupListSURes         PDUSessionResourceSetupListSURes         `,ignore,optional`
	PDUSessionResourceFailedToSetupListSURes PDUSessionResourceFailedToSetupListSURes `,ignore,optional`
	CriticalityDiagnostics                   CriticalityDiagnostics                   `,ignore,optional`
}
