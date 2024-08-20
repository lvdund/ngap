package ie

type SourceToTargetAmfInformationReroute struct {
ConfiguredNssai	[]byte	//`bitstring:"sizeLB:128,sizeUB:128"`
RejectedNssaiInPlmn	[]byte	//`bitstring:"sizeLB:32,sizeUB:32"`
RejectedNssaiInTa	[]byte	//`bitstring:"sizeLB:32,sizeUB:32"`
}
