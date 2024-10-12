package ie

type CellTrafficTraceIEs struct {
	AMFUENGAPID                    AMFUENGAPID           `,reject,mandatory`
	RANUENGAPID                    RANUENGAPID           `,reject,mandatory`
	NGRANTraceID                   NGRANTraceID          `,ignore,mandatory`
	NGRANCGI                       NGRANCGI              `,ignore,mandatory`
	TraceCollectionEntityIPAddress TransportLayerAddress `,ignore,mandatory`
}
