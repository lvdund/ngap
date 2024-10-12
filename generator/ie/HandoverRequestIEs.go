package ie

type HandoverRequestIEs struct {
	AMFUENGAPID                                 AMFUENGAPID                                 `,reject,mandatory`
	HandoverType                                HandoverType                                `,reject,mandatory`
	Cause                                       Cause                                       `,ignore,mandatory`
	UEAggregateMaximumBitRate                   UEAggregateMaximumBitRate                   `,reject,mandatory`
	CoreNetworkAssistanceInformationForInactive CoreNetworkAssistanceInformationForInactive `,ignore,optional`
	UESecurityCapabilities                      UESecurityCapabilities                      `,reject,mandatory`
	SecurityContext                             SecurityContext                             `,reject,mandatory`
	NewSecurityContextInd                       NewSecurityContextInd                       `,reject,optional`
	NASC                                        NASPDU                                      `,reject,optional`
	PDUSessionResourceSetupListHOReq            PDUSessionResourceSetupListHOReq            `,reject,mandatory`
	AllowedNSSAI                                AllowedNSSAI                                `,reject,mandatory`
	TraceActivation                             TraceActivation                             `,ignore,optional`
	MaskedIMEISV                                MaskedIMEISV                                `,ignore,optional`
	SourceToTargetTransparentContainer          SourceToTargetTransparentContainer          `,reject,mandatory`
	MobilityRestrictionList                     MobilityRestrictionList                     `,ignore,optional`
	LocationReportingRequestType                LocationReportingRequestType                `,ignore,optional`
	RRCInactiveTransitionReportRequest          RRCInactiveTransitionReportRequest          `,ignore,optional`
	GUAMI                                       GUAMI                                       `,reject,mandatory`
	RedirectionVoiceFallback                    RedirectionVoiceFallback                    `,ignore,optional`
	CNAssistedRANTuning                         CNAssistedRANTuning                         `,ignore,optional`
}
