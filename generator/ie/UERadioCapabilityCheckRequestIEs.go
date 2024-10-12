package ie

type UERadioCapabilityCheckRequestIEs struct {
	AMFUENGAPID       AMFUENGAPID       `,reject,mandatory`
	RANUENGAPID       RANUENGAPID       `,reject,mandatory`
	UERadioCapability UERadioCapability `,ignore,optional`
}
