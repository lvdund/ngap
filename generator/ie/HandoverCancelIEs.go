package ie

type HandoverCancelIEs struct {
	AMFUENGAPID AMFUENGAPID `,reject,mandatory`
	RANUENGAPID RANUENGAPID `,reject,mandatory`
	Cause       Cause       `,ignore,mandatory`
}
