package ie

type UERadioCapabilityCheckResponseIEs struct {
	AMFUENGAPID              AMFUENGAPID              `,ignore,mandatory`
	RANUENGAPID              RANUENGAPID              `,ignore,mandatory`
	IMSVoiceSupportIndicator IMSVoiceSupportIndicator `,reject,mandatory`
	CriticalityDiagnostics   CriticalityDiagnostics   `,ignore,optional`
}
