package ie

type UEContextModificationResponseIEs struct {
	AMFUENGAPID             AMFUENGAPID             `,ignore,mandatory`
	RANUENGAPID             RANUENGAPID             `,ignore,mandatory`
	RRCState                RRCState                `,ignore,optional`
	UserLocationInformation UserLocationInformation `,ignore,optional`
	CriticalityDiagnostics  CriticalityDiagnostics  `,ignore,optional`
}
