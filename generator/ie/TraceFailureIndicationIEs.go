package ie

type TraceFailureIndicationIEs struct {
	AMFUENGAPID  AMFUENGAPID  `,reject,mandatory`
	RANUENGAPID  RANUENGAPID  `,reject,mandatory`
	NGRANTraceID NGRANTraceID `,ignore,mandatory`
	Cause        Cause        `,ignore,mandatory`
}
