package ie

type MbsDistributionReleaseRequestTransfer struct {
MbsSessionId	MbsSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
MbsAreaSessionId	MbsAreaSessionId	//`bitstring:"sizeLB:0,sizeUB:150"`
SharedNgUUnicastTnlInformation	UpTransportLayerInformation	//`bitstring:"sizeLB:0,sizeUB:150"`
Cause	Cause	//`bitstring:"sizeLB:0,sizeUB:150"`
}
