package ie

type PathSwitchRequestAcknowledgeIEs struct {
	AMFUENGAPID                                 AMFUENGAPID                                 `,ignore,mandatory`
	RANUENGAPID                                 RANUENGAPID                                 `,ignore,mandatory`
	UESecurityCapabilities                      UESecurityCapabilities                      `,reject,optional`
	SecurityContext                             SecurityContext                             `,reject,mandatory`
	NewSecurityContextInd                       NewSecurityContextInd                       `,reject,optional`
	PDUSessionResourceSwitchedList              PDUSessionResourceSwitchedList              `,ignore,mandatory`
	PDUSessionResourceReleasedListPSAck         PDUSessionResourceReleasedListPSAck         `,ignore,optional`
	AllowedNSSAI                                AllowedNSSAI                                `,reject,mandatory`
	CoreNetworkAssistanceInformationForInactive CoreNetworkAssistanceInformationForInactive `,ignore,optional`
	RRCInactiveTransitionReportRequest          RRCInactiveTransitionReportRequest          `,ignore,optional`
	CriticalityDiagnostics                      CriticalityDiagnostics                      `,ignore,optional`
	RedirectionVoiceFallback                    RedirectionVoiceFallback                    `,ignore,optional`
	CNAssistedRANTuning                         CNAssistedRANTuning                         `,ignore,optional`
}
