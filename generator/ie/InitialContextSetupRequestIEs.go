package ie

type InitialContextSetupRequestIEs struct {
	AMFUENGAPID                                 AMFUENGAPID                                 `,reject,mandatory`
	RANUENGAPID                                 RANUENGAPID                                 `,reject,mandatory`
	OldAMF                                      AMFName                                     `,reject,optional`
	UEAggregateMaximumBitRate                   UEAggregateMaximumBitRate                   `,reject,conditional`
	CoreNetworkAssistanceInformationForInactive CoreNetworkAssistanceInformationForInactive `,ignore,optional`
	GUAMI                                       GUAMI                                       `,reject,mandatory`
	PDUSessionResourceSetupListCxtReq           PDUSessionResourceSetupListCxtReq           `,reject,optional`
	AllowedNSSAI                                AllowedNSSAI                                `,reject,mandatory`
	UESecurityCapabilities                      UESecurityCapabilities                      `,reject,mandatory`
	SecurityKey                                 SecurityKey                                 `,reject,mandatory`
	TraceActivation                             TraceActivation                             `,ignore,optional`
	MobilityRestrictionList                     MobilityRestrictionList                     `,ignore,optional`
	UERadioCapability                           UERadioCapability                           `,ignore,optional`
	IndexToRFSP                                 IndexToRFSP                                 `,ignore,optional`
	MaskedIMEISV                                MaskedIMEISV                                `,ignore,optional`
	NASPDU                                      NASPDU                                      `,ignore,optional`
	EmergencyFallbackIndicator                  EmergencyFallbackIndicator                  `,reject,optional`
	RRCInactiveTransitionReportRequest          RRCInactiveTransitionReportRequest          `,ignore,optional`
	UERadioCapabilityForPaging                  UERadioCapabilityForPaging                  `,ignore,optional`
	RedirectionVoiceFallback                    RedirectionVoiceFallback                    `,ignore,optional`
	LocationReportingRequestType                LocationReportingRequestType                `,ignore,optional`
	CNAssistedRANTuning                         CNAssistedRANTuning                         `,ignore,optional`
}
