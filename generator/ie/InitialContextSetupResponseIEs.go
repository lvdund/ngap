package ie

type InitialContextSetupResponseIEs struct {
	AMFUENGAPID                               AMFUENGAPID                               `,ignore,mandatory`
	RANUENGAPID                               RANUENGAPID                               `,ignore,mandatory`
	PDUSessionResourceSetupListCxtRes         PDUSessionResourceSetupListCxtRes         `,ignore,optional`
	PDUSessionResourceFailedToSetupListCxtRes PDUSessionResourceFailedToSetupListCxtRes `,ignore,optional`
	CriticalityDiagnostics                    CriticalityDiagnostics                    `,ignore,optional`
}
