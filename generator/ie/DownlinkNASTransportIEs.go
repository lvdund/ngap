package ie

type DownlinkNASTransportIEs struct {
	AMFUENGAPID               AMFUENGAPID               `,reject,mandatory`
	RANUENGAPID               RANUENGAPID               `,reject,mandatory`
	OldAMF                    AMFName                   `,reject,optional`
	RANPagingPriority         RANPagingPriority         `,ignore,optional`
	NASPDU                    NASPDU                    `,reject,mandatory`
	MobilityRestrictionList   MobilityRestrictionList   `,ignore,optional`
	IndexToRFSP               IndexToRFSP               `,ignore,optional`
	UEAggregateMaximumBitRate UEAggregateMaximumBitRate `,ignore,optional`
	AllowedNSSAI              AllowedNSSAI              `,reject,optional`
}
