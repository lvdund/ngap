package ie

type OverloadStartIEs struct {
	AMFOverloadResponse               OverloadResponse               `,reject,optional`
	AMFTrafficLoadReductionIndication TrafficLoadReductionIndication `,ignore,optional`
	OverloadStartNSSAIList            OverloadStartNSSAIList         `,ignore,optional`
}
