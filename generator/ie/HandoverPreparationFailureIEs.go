package ie

type HandoverPreparationFailureIEs struct {
	AMFUENGAPID            AMFUENGAPID            `,ignore,mandatory`
	RANUENGAPID            RANUENGAPID            `,ignore,mandatory`
	Cause                  Cause                  `,ignore,mandatory`
	CriticalityDiagnostics CriticalityDiagnostics `,ignore,optional`
}
