package ie

type InitialUEMessageIEs struct {
	RANUENGAPID                         RANUENGAPID                         `,reject,mandatory`
	NASPDU                              NASPDU                              `,reject,mandatory`
	UserLocationInformation             UserLocationInformation             `,reject,mandatory`
	RRCEstablishmentCause               RRCEstablishmentCause               `,ignore,mandatory`
	FiveGSTMSI                          FiveGSTMSI                          `,reject,optional`
	AMFSetID                            AMFSetID                            `,ignore,optional`
	UEContextRequest                    UEContextRequest                    `,ignore,optional`
	AllowedNSSAI                        AllowedNSSAI                        `,reject,optional`
	SourceToTargetAMFInformationReroute SourceToTargetAMFInformationReroute `,ignore,optional`
	SelectedPLMNIdentity                PLMNIdentity                        `,ignore,optional`
}
