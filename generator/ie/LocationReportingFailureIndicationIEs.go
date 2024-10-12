package ie

type LocationReportingFailureIndicationIEs struct {
	AMFUENGAPID AMFUENGAPID `,reject,mandatory`
	RANUENGAPID RANUENGAPID `,reject,mandatory`
	Cause       Cause       `,ignore,mandatory`
}
