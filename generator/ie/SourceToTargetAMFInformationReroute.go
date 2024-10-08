package ie

type SourceToTargetAMFInformationReroute struct {
	ConfiguredNSSAI     ConfiguredNSSAI                           `False,OPTIONAL`
	RejectedNSSAIinPLMN RejectedNSSAIinPLMN                       `False,OPTIONAL`
	RejectedNSSAIinTA   RejectedNSSAIinTA                         `False,OPTIONAL`
	IEExtensions        SourceToTargetAMFInformationRerouteExtIEs `False,OPTIONAL`
}
