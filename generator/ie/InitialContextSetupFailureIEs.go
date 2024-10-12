package ie

type InitialContextSetupFailureIEs struct {
	AMFUENGAPID                                AMFUENGAPID                                `,ignore,mandatory`
	RANUENGAPID                                RANUENGAPID                                `,ignore,mandatory`
	PDUSessionResourceFailedToSetupListCxtFail PDUSessionResourceFailedToSetupListCxtFail `,ignore,optional`
	Cause                                      Cause                                      `,ignore,mandatory`
	CriticalityDiagnostics                     CriticalityDiagnostics                     `,ignore,optional`
}
