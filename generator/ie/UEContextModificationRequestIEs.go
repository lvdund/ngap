package ie

type UEContextModificationRequestIEs struct {
	AMFUENGAPID                                 AMFUENGAPID                                 `,reject,mandatory`
	RANUENGAPID                                 RANUENGAPID                                 `,reject,mandatory`
	RANPagingPriority                           RANPagingPriority                           `,ignore,optional`
	SecurityKey                                 SecurityKey                                 `,reject,optional`
	IndexToRFSP                                 IndexToRFSP                                 `,ignore,optional`
	UEAggregateMaximumBitRate                   UEAggregateMaximumBitRate                   `,ignore,optional`
	UESecurityCapabilities                      UESecurityCapabilities                      `,reject,optional`
	CoreNetworkAssistanceInformationForInactive CoreNetworkAssistanceInformationForInactive `,ignore,optional`
	EmergencyFallbackIndicator                  EmergencyFallbackIndicator                  `,reject,optional`
	NewAMFUENGAPID                              AMFUENGAPID                                 `,reject,optional`
	RRCInactiveTransitionReportRequest          RRCInactiveTransitionReportRequest          `,ignore,optional`
	NewGUAMI                                    GUAMI                                       `,reject,optional`
	CNAssistedRANTuning                         CNAssistedRANTuning                         `,ignore,optional`
}
