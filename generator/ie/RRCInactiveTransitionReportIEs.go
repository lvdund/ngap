package ie

type RRCInactiveTransitionReportIEs struct {
	AMFUENGAPID             AMFUENGAPID             `,reject,mandatory`
	RANUENGAPID             RANUENGAPID             `,reject,mandatory`
	RRCState                RRCState                `,ignore,mandatory`
	UserLocationInformation UserLocationInformation `,ignore,mandatory`
}
